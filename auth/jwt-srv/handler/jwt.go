package handler

import (
	"context"
	"errors"
	"time"

	"github.com/micro/go-micro/util/log"

	jwt "github.com/dmitrymomot/go-micro/auth/jwt-srv/proto/jwt"
)

// Jwt struct
type Jwt struct{}

// NewToken is a single request handler called via client.NewToken or the generated client code
func (e *Jwt) NewToken(ctx context.Context, req *jwt.NewTokenRequest, rsp *jwt.NewTokenResponse) error {
	log.Log("Received Jwt.NewToken request")
	if req.UserId == "" {
		return errors.New("wrong user id")
	}
	rsp.AccessToken = "token_for_user_" + req.UserId
	rsp.ExpiresAt = time.Now().Add(10080 * time.Second).Unix()
	return nil
}
