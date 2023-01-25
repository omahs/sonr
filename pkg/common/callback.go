package common

import (
	"log"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

// NodeCallback is an interface with three methods: OnDiscover, OnLinking, and OnTopicMessage.
// @property OnDiscover - This is called when a node is discovered. The data is the data that was sent
// by the node.
// @property OnLinking - This is called when a node is linking to the gateway.
// @property OnTopicMessage - This is the callback that will be called when a message is received on a
// topic.
type NodeCallback interface {
	OnDiscover(data []byte)
	OnLinking(data []byte)
}

// `defaultCallback` is a type that implements the `Callback` interface.
type defaultCallback struct{}

// It returns a pointer to a defaultCallback struct
func DefaultCallback() NodeCallback {
	return &defaultCallback{}
}

// Defining a method on the defaultCallback struct.
func (cb *defaultCallback) OnDiscover(data []byte) {
	log.Println("ERROR: MotorCallback not implemented.")
}

// Defining a method on the defaultCallback struct.
func (cb *defaultCallback) OnLinking(data []byte) {
	log.Println("ERROR: MotorCallback not implemented.")
}

// Defining a method on the defaultCallback struct.
func (cb *defaultCallback) OnTopicMessage(topic string, data []byte) {
	log.Println("ERROR: MotorCallback not implemented.")
}

// `SNRPubKey` is a `PubKey` that has a `DID` and a `Multibase`
// @property {string} DID - The DID of the SNR
// @property {string} Multibase - The multibase encoding of the DID.
type SNRPubKey interface {
	cryptotypes.PubKey

	Bech32(pfix string) (string, error)
	DID() string
	Multibase() string
	Raw() []byte
}
