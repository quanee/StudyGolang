// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"blog/internal/biz"
	"blog/internal/conf"
	"blog/internal/data"
	"blog/internal/server"
	"blog/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		cache.ProviderSet,
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		mq.ProviderSet,
		newApp))
}
