// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"blog/internal/biz"
	"blog/internal/conf"
	"blog/internal/data"
	"blog/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Data, *conf.Cache, *conf.MessageQueue, log.Logger) (func(context.Context) error, func(), error) {
	panic(wire.Build(
		cache.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.JobProviderSet,
		mq.ProviderSet,
		newApp))
}
