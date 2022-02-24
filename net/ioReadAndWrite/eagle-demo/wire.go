//go:build wireinject
// +build wireinject

package main

import (
	"eagle-demo/internal/server"
	"eagle-demo/internal/service"
	eagle "github.com/go-eagle/eagle/pkg/app"
	"github.com/google/wire"
)

func InitApp(cfg *eagle.Config, config *eagle.ServerConfig) (*eagle.App, error) {
	wire.Build(server.ProviderSet, service.ProviderSet, newApp)
	return &eagle.App{}, nil
}
