package auth


import (
	stdjwt "github.com/dgrijalva/jwt-go"
	gokitjwt "github.com/go-kit/kit/auth/jwt"
	"context"
)

type User interface{
	GetUserName() (string)
	GetAudience() (string)
	GetFullName() (string)
}

type user struct {
	Username string
	Audience string //App(Web, Mobile, Third Party's name) sign in
	FullName string
}

func (u user) GetUserName() string {
	return u.Username
}

func (u user) GetAudience() string {
	return u.Audience
}

func (u user) GetFullName() string {
	return u.FullName
}

func LoadUserFromContext(ctx context.Context) User{
	claim := ctx.Value(gokitjwt.JWTClaimsContextKey)
	u := claim.(*stdjwt.StandardClaims)
	return &user{Username: u.Subject, FullName:u.Issuer, Audience:u.Audience}
}
