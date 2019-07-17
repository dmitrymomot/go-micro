// Package k8s implements go-config with env and kubernetes configmap
package k8s

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-plugins/config/source/configmap"
)

// NewConfig returns config with env and k8s configmap setup
func NewConfig(opts ...config.Option) config.Config {
	cfg := config.NewConfig(opts...)
	cfg.Load(
		env.NewSource(),
		configmap.NewSource(),
	)
	return cfg
}
