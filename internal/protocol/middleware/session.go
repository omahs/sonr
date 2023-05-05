package middleware

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/internal/protocol/config"
	servicetypes "github.com/sonrhq/core/x/service/types"
)

type Session struct {
	ServiceOrigin string `json:"service_origin"`
	Challenge     string `json:"challenge"`
	Alias         string `json:"alias"`
}


var conf *config.ProtocolConfig

func Init(c *config.ProtocolConfig) {
	conf = c
}


// GetServiceRecord returns the service record for the current session.
func (s *Session) GetServiceRecord() (*servicetypes.ServiceRecord, error) {
	if s.ServiceOrigin == "" {
		return nil, fmt.Errorf("service origin not set")
	}
	return local.Context().GetService(context.Background(), s.ServiceOrigin)
}

func FetchSession(c *fiber.Ctx) (*Session, error) {
	sess, err := conf.SessionStore.Get(c)
	if err != nil {
		return nil, err
	}
	q := ParseQuery(c)
	sessKey := KeySessionID(q.Origin(), q.Alias())

	jsonValue := sess.Get(sessKey)
	if jsonValue == nil {
		return nil, fmt.Errorf("session not found")
	}

    jsonBz, ok := jsonValue.([]byte)
    if !ok {
        return nil, fmt.Errorf("invalid session data type")
    }


	var session Session
	err = json.Unmarshal(jsonBz, &session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func StoreSession(c *fiber.Ctx, challenge string) error {
	sess, err := conf.SessionStore.Get(c)
	if err != nil {
		return err
	}
    q := ParseQuery(c)
	session := &Session{
		ServiceOrigin: q.Origin(),
		Challenge:     challenge,
		Alias:         q.Alias(),
	}
    sessKey := KeySessionID(session.ServiceOrigin, session.Alias)
    jsonBz, err := json.Marshal(session)
    if err != nil {
        return err
    }

    sess.Set(sessKey, jsonBz)
    if err != nil {
        return err
    }

    return nil
}

func KeySessionID(origin, alias string) string {
	return fmt.Sprintf("%s:%s", origin, alias)
}
