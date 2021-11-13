// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"recognition/internal/biz"
	"recognition/internal/conf"
	"recognition/internal/data"
	"recognition/internal/server"
	"recognition/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// init app by wire injection
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
