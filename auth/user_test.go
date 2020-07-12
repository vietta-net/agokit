package auth_test

import (
	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/auth/jwt"
	"google.golang.org/grpc/metadata"
	"context"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/vietta-net/agokit/auth"

)
var (
	secret = "7412ea46fc23a86b96e5f3eb192ad280"
)

func TestUser(t *testing.T) {

	ctx := SetContext()
	var kid = "kid-header"
	token := stdjwt.NewWithClaims(stdjwt.SigningMethodHS256, stdjwt.StandardClaims{})
	token.Header["kid"] = kid

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))

	assert.Nil(t, err)
	ctx = context.WithValue(ctx, jwt.JWTTokenContextKey, tokenString)


	user := auth.LoadUserFromContext(ctx)
	t.Run("GetUserName", func(t *testing.T) {
		assert.Equal(t, user.GetUserName(), "pntn79")
	})

	t.Run("GetFullName", func(t *testing.T) {
		assert.Equal(t, user.GetFullName(), "John Doe")
	})

	t.Run("GetAudience", func(t *testing.T) {
		assert.Equal(t, user.GetAudience(), "Web")
	})

}

func SetContext()(ctx context.Context){


	ctx = context.Background()

	/*
		HEADER:ALGORITHM & TOKEN TYPE
		{
		  "alg": "HS256",
		  "typ": "JWT"
		}
		JWT PAYLOAD:DATA
		{
		  "sub": "pntn79",
		  "iss": "John Doe",
		  "aud": "Web"
		}

	 */
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJwbnRuNzkiLC" +
		"Jpc3MiOiJKb2huIERvZSIsImF1ZCI6IldlYiJ9.v0ftayqSMDefdanGepNm8aYWWI3JTJ5zyTsdrhguQKs"

	//https://tools.ietf.org/html/rfc4647
	md := metadata.Pairs(
		"content-language", "vi",
		"timezone", "Asia/Ho_Chi_Minh",
		"authorization" , "Bearer " + tokenString)
	ctx = metadata.NewOutgoingContext(ctx, md)

	kf := func(token *stdjwt.Token) (interface{}, error) { return []byte(secret), nil }
	token, _ := stdjwt.ParseWithClaims(tokenString, jwt.StandardClaimsFactory(),kf)
	ctx 	= context.WithValue(ctx, jwt.JWTClaimsContextKey, token.Claims)

	return ctx
}