// Package k8s implements a go-micro service for kubernetes
package k8s

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/kubernetes"

	// static selector offloads load balancing to k8s services
	// note: requires user to create k8s services
	"github.com/micro/go-plugins/client/selector/static"
)

// NewWebService returns a new go-micro service pre-initialised for k8s
func NewWebService(opts ...web.Option) web.Service {
	// create registry and selector
	r := kubernetes.NewRegistry()
	s := static.NewSelector()

	// set the registry and selector
	options := []web.Option{
		web.MicroService(micro.NewService(
			micro.Registry(r),
			micro.Selector(s),
		)),
	}

	// append user options
	options = append(options, opts...)

	// return a micro.Service
	return web.NewService(options...)
}
