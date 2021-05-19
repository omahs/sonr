package bind

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/sonr-io/core/internal/crypto"
	"github.com/sonr-io/core/internal/topic"
	sc "github.com/sonr-io/core/pkg/client"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// @ Return URLLink
func GetURLLink(url string) []byte {
	// Create Link
	link := md.NewURLLink(url)

	// Marshal
	bytes, err := proto.Marshal(link)
	if err != nil {
		return nil
	}
	return bytes
}

// @ Gets User from Storj
func GetUser(data []byte) []byte {
	// Unmarshal Request
	request := &md.StorjRequest{}
	proto.Unmarshal(data, request)

	// Get User from Uplink
	user, err := sc.GetUser(context.Background(), request.StorjApiKey, request.GetUserID())
	if err != nil {
		sentry.CaptureException(err)
		return nil
	}

	// Marshal
	bytes, err := proto.Marshal(user)
	if err != nil {
		return nil
	}
	return bytes
}

// @ Puts User into Storj
func PutUser(data []byte) bool {
	// Unmarshal Request
	request := &md.StorjRequest{}
	proto.Unmarshal(data, request)

	// Put User
	err := sc.PutUser(context.Background(), request.StorjApiKey, request.GetUser())
	if err != nil {
		sentry.CaptureException(err)
		return false
	}
	return true

}

// @ Join Existing Group
func (mn *Node) CreateRemote() []byte {
	if mn.isReady() {
		// Generate Word List
		_, wordList, serr := crypto.RandomWords("english", 3)
		if serr != nil {
			mn.handleError(serr)
			return nil
		}
		// Create Remote Request and Join Lobby
		remote := md.NewRemoteInfo(wordList)

		// Join Lobby
		tm, serr := mn.client.JoinLobby(remote, true)
		if serr != nil {
			mn.handleError(serr)
			return nil
		}

		// Set Topic
		mn.topics[remote.Topic] = tm

		// Marshal
		data, err := proto.Marshal(remote)
		if err != nil {
			mn.handleError(md.NewError(err, md.ErrorMessage_MARSHAL))
			return nil
		}
		return data
	}
	return nil
}

// @ Join Existing Group
func (mn *Node) JoinRemote(data []byte) {
	if mn.isReady() {
		// Unpackage Data
		remote := &md.RemoteInfo{}
		err := proto.Unmarshal(data, remote)
		if err != nil {
			mn.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
			return
		}

		// Join Lobby
		tm, serr := mn.client.JoinLobby(remote, false)
		if err != nil {
			mn.handleError(serr)
			return
		}

		// Set Topic
		mn.topics[remote.Topic] = tm
	}
}

// @ Leave Existing Group
func (mn *Node) LeaveRemote(data []byte) {
	if mn.isReady() {
		// Unpackage Data
		remote := md.RemoteInfo{}
		err := proto.Unmarshal(data, &remote)
		if err != nil {
			mn.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
			return
		}

		// Join Lobby
		serr := mn.client.LeaveLobby(mn.topics[remote.Topic])
		if err != nil {
			mn.handleError(serr)
			return
		}
	}
}

// @ Update proximity/direction and Notify Lobby
func (mn *Node) Update(data []byte) {
	if mn.isReady() {
		// Initialize from Request
		update := &md.UpdateRequest{}
		if err := proto.Unmarshal(data, update); err != nil {
			mn.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
			return
		}

		// Update Peer
		mn.user.Update(update)

		// Notify Local Lobby
		err := mn.client.Update(mn.local)
		if err != nil {
			mn.handleError(err)
			return
		}
	}
}

// @ Invite Processes Data and Sends Invite to Peer
func (mn *Node) Invite(data []byte) {
	if mn.isReady() {
		// Update Status
		mn.setStatus(md.Status_PENDING)

		// Initialize from Request
		req := &md.InviteRequest{}
		if err := proto.Unmarshal(data, req); err != nil {
			mn.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
			return
		}

		// Retreive Invite Topic
		var topic *topic.TopicManager
		if req.IsRemote && req.Remote != nil {
			topic = mn.topics[req.Remote.Topic]
		} else {
			topic = mn.local
		}

		// @ 2. Check Transfer Type
		if req.Payload == md.Payload_CONTACT || req.Payload == md.Payload_FLAT_CONTACT {
			err := mn.client.InviteContact(req, topic, mn.user.Contact)
			if err != nil {
				mn.handleError(err)
				return
			}
		} else if req.Payload == md.Payload_URL {
			err := mn.client.InviteLink(req, topic)
			if err != nil {
				mn.handleError(err)
				return
			}
		} else {
			// Invite With file
			err := mn.client.InviteFile(req, topic)
			if err != nil {
				mn.handleError(err)
				return
			}
		}
	}
}

// @ Respond to an Invite with Decision
func (mn *Node) Respond(data []byte) {
	if mn.isReady() {
		// Initialize from Request
		req := &md.RespondRequest{}
		if err := proto.Unmarshal(data, req); err != nil {
			mn.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
			return
		}

		// Retreive Invite Topic
		var topic *topic.TopicManager
		if req.IsRemote && req.Remote != nil {
			topic = mn.topics[req.Remote.Topic]
		} else {
			topic = mn.local
		}

		mn.client.Respond(req, topic)
		// Update Status
		if req.Decision {
			mn.setStatus(md.Status_INPROGRESS)
		} else {
			mn.setStatus(md.Status_AVAILABLE)
		}
	}
}
