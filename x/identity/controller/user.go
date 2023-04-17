package controller

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v4"
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

	// Map of the dids of the keyshares to the dids of the accounts
	Accounts []string `json:"accounts"`

	// DidDocument of the primary identity
	PrimaryIdentity *types.DidDocument `json:"primaryIdentity"`

	controller Controller
}

func NewUser(c Controller, username string) *User {
	accDids := make([]string, 0)
	accs, err := c.ListAccounts()
	if err != nil {
		return nil
	}

	for _, acc := range accs {
		accDids = append(accDids, acc.Did())
	}

	return &User{
		Did:             c.Did(),
		Accounts:        accDids,
		PrimaryIdentity: c.PrimaryIdentity(),
		Username:        username,
		controller:      c,
	}
}

func LoadUser(data []byte) (*User, error) {
	var u User
	err := json.Unmarshal(data, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (u *User) Marshal() ([]byte, error) {
	return json.Marshal(u)
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
		"accounts": u.Accounts,
	}
}

func (u *User) JWT(secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, u.JWTClaims())
	return token.SignedString(secret)
}

func UserFromJWT(token *jwt.Token) (*User, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	did, ok := claims["did"].(string)
	if !ok {
		return nil, ErrInvalidToken
	}

	username, ok := claims["username"].(string)
	if !ok {
		return nil, ErrInvalidToken
	}

	accounts, ok := claims["accounts"].([]string)
	if !ok {
		return nil, ErrInvalidToken
	}

	return &User{
		Did:             did,
		Username:        username,
		Accounts:        accounts,
	}, nil
}
