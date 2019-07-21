package client

import (
	"context"

	jwt "github.com/dmitrymomot/go-micro/auth/jwtsrv/proto/jwt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
)

type jwtKey struct{}

// JwtFromContext retrieves the client from the Context
func JwtFromContext(ctx context.Context) (jwt.JwtService, bool) {
	c, ok := ctx.Value(jwtKey{}).(jwt.JwtService)
	return c, ok
}

// JwtWrapper returns a wrapper for the JwtClient
func JwtWrapper(service micro.Service) server.HandlerWrapper {
	client := jwt.NewJwtService("jwtsrv", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, jwtKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
