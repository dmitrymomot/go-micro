package main

import (
	"github.com/dmitrymomot/go-micro/micro/k8s"
	"github.com/micro/go-micro/util/log"

	"github.com/dmitrymomot/go-micro/auth/jwtapi/client"
	"github.com/dmitrymomot/go-micro/auth/jwtapi/handler"

	"github.com/micro/go-micro"

	jwt "github.com/dmitrymomot/go-micro/auth/jwtapi/proto/jwt"
)

func main() {
	// New Service
	service := k8s.NewService(
		micro.Name("jwtapi"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Jwt srv client
		micro.WrapHandler(client.JwtWrapper(service)),
	)

	// Register Handler
	jwt.RegisterJwtHandler(service.Server(), new(handler.Jwt))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
