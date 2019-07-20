// Package k8s implements a go-micro service for kubernetes
package k8s

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/kubernetes"

	// static selector offloads load balancing to k8s services
	// enable with MICRO_SELECTOR=static or --selector=static
	// requires user to create k8s services
	"github.com/micro/go-plugins/client/selector/static"
)

// NewWebService returns a web service for kubernetes
func NewWebService(opts ...web.Option) web.Service {
	// setup
	k := kubernetes.NewRegistry()
	st := static.NewSelector()

	// create new service
	service := grpc.NewService(
		micro.Registry(k),
		micro.Selector(st),
	)

	// prepend option
	options := []web.Option{
		web.MicroService(service),
	}

	options = append(options, opts...)

	// return new service
	return web.NewService(options...)
}
