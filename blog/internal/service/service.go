package service

import (
	"blog/internal/biz"
	"blog/internal/cache"
	"blog/internal/mq"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService)

var JobProviderSet = wire.NewSet(NewJobService)

type JobService struct {
	ttl     int
	blog    *biz.BlogUsecase
	summary *biz.SummaryUsecase
	log     *log.Helper
	cache   cache.Cache
	mq      mq.MessageQueue
}
