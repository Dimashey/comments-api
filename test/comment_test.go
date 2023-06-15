//go:build e2e
// +build e2e

package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	BASE_URL = "http://localhost:8080"
)

func createToken() string {
	token := jwt.New(jwt.SigningMethodHS256)

	tokenString, err := token.SignedString([]byte("missionimpossible"))
	if err != nil {
		fmt.Println(err)
	}

	return tokenString
}

func TestPostComment(t *testing.T) {
	t.Run("can post comment", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", "bearer "+createToken()).
			SetBody(`{"slug": "/", "author": "12345", "body": "hello world"}`).
			Post(BASE_URL + "/api/v1/comment")
		assert.NoError(t, err)

		assert.Equal(t, 200, resp.StatusCode())
	})

	t.Run("cannot post comment without JWT", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetBody(`{"slug": "/", "author": "12345", "body": "hello world"}`).
			Post(BASE_URL + "/api/v1/comment")
		assert.NoError(t, err)

		assert.Equal(t, 401, resp.StatusCode())
	})
}
