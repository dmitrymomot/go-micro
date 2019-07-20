package main

import (
	"jwt-srv/handler"

	"github.com/dmitrymomot/go-micro/micro/k8s"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	jwt "jwt-srv/proto/jwt"
)

func main() {
	// New Service
	service := k8s.NewService(
		micro.Name("go.micro.srv.jwt"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	jwt.RegisterJwtHandler(service.Server(), new(handler.Jwt))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
