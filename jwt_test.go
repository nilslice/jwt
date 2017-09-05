package jwt_test

import (
	"log"
	"testing"

	"github.com/nilslice/jwt"
)

func TestGrantPasses(t *testing.T) {
	claims := map[string]interface{}{
		"testID": 1,
		"name":   "GrantPasses",
	}

	grant, err := jwt.New(claims)
	if err != nil {
		t.Log(claims, grant, err)
		t.Fail()
	}

	if !jwt.Passes(grant) {
		t.Fail()
	}
}

func TestGrantFails(t *testing.T) {
	claims := map[string]interface{}{
		"testID": 2,
		"name":   "GrantFails",
	}

	grant, err := jwt.New(claims)
	if err != nil {
		t.Log(claims, grant, err)
		t.Fail()
	}

	jwt.Secret([]byte("NotPreviousSecret"))

	if jwt.Passes(grant) {
		t.Fail()
	}
}

func TestGetClaims(t *testing.T) {
	email := "hello@nilslice.xyz"
	authLevel := float64(7)

	claims := map[string]interface{}{
		"email":      email,
		"auth_level": authLevel,
	}

	token, err := jwt.New(claims)
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	claims = jwt.GetClaims(token)
	if claims["email"].(string) != email ||
		claims["auth_level"].(float64) != authLevel {
		t.Fail()
	}

}
