// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/jwt/jwt.proto

package go_micro_srv_jwt

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Jwt service

type JwtService interface {
	NewToken(ctx context.Context, in *NewTokenRequest, opts ...client.CallOption) (*NewTokenResponse, error)
}

type jwtService struct {
	c    client.Client
	name string
}

func NewJwtService(name string, c client.Client) JwtService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.jwt"
	}
	return &jwtService{
		c:    c,
		name: name,
	}
}

func (c *jwtService) NewToken(ctx context.Context, in *NewTokenRequest, opts ...client.CallOption) (*NewTokenResponse, error) {
	req := c.c.NewRequest(c.name, "Jwt.NewToken", in)
	out := new(NewTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Jwt service

type JwtHandler interface {
	NewToken(context.Context, *NewTokenRequest, *NewTokenResponse) error
}

func RegisterJwtHandler(s server.Server, hdlr JwtHandler, opts ...server.HandlerOption) error {
	type jwt interface {
		NewToken(ctx context.Context, in *NewTokenRequest, out *NewTokenResponse) error
	}
	type Jwt struct {
		jwt
	}
	h := &jwtHandler{hdlr}
	return s.Handle(s.NewHandler(&Jwt{h}, opts...))
}

type jwtHandler struct {
	JwtHandler
}

func (h *jwtHandler) NewToken(ctx context.Context, in *NewTokenRequest, out *NewTokenResponse) error {
	return h.JwtHandler.NewToken(ctx, in, out)
}