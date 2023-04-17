package controller

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/x/identity/types"

	v1 "github.com/sonrhq/core/types/highway/v1"
)

var (
	// ErrInvalidToken is returned when the token is invalid
	ErrInvalidToken = fmt.Errorf("invalid token")
)

type User struct {
	// DID of the user
	Did string `json:"_id"`

	// DID document of the primary identity
	Username string `json:"username"`



	controller Controller
}

func NewUser(c Controller, username string) *User {
	return &User{
		Did:             c.Did(),
		Username:        username,
		controller:      c,
	}
}

func LoadUser(token *jwt.Token) (*User, error) {
	claims := token.Claims.(jwt.MapClaims)
	did, ok := claims["did"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid did")
	}
	username, ok := claims["username"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid username")
	}
	return &User{
		Did:             did,
		Username:        username,
	}, nil
}


func (u *User) ListAccounts() ([]*v1.Account, error) {
	accs := make([]*v1.Account, 0)
	lclAccs, err := u.controller.ListAccounts()
	if err != nil {
		return nil, err
	}
	for _, lclAcc := range lclAccs {
		accs = append(accs, lclAcc.ToProto())
	}
	return accs, nil
}

func (u *User) JWTClaims() (jwt.MapClaims) {
	return jwt.MapClaims{
		"did": u.Did,
		"username": u.Username,
	}
}
func (u *User) PrimaryIdentity() (*types.DidDocument, error) {
	return local.Context().GetDID(context.Background(), u.Did)
}
func (u *User) JWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, u.JWTClaims())
	return token.SignedString(local.Context().SigningKey())
}
