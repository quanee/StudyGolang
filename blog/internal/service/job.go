package service

import (
	"blog/internal/data/ent"
	"context"
	"encoding/json"
	"strconv"

	"blog/internal/biz"
	"blog/internal/cache/redis"
	"blog/internal/mq/kafka"
	"github.com/go-kratos/kratos/v2/log"
)

func NewJobService(summary *biz.SummaryUsecase, cache *redis.RadixRC3, mq *kafka.KafkaClient, logger log.Logger) *JobService {
	return &JobService{
		ttl:   5000,
		summary:  summary,
		cache: cache,
		mq:    mq,
		log:   log.NewHelper(logger),
	}
}

func (j *JobService) UpdateCache(ctx context.Context) error {
	return j.mq.Consume(j.writeCache)
}

func (j *JobService) writeCache(ctx context.Context, msg string, err error) error {
	if err != nil {
		return err
	}

	var req ent.Summary
	err = json.Unmarshal([]byte(msg), &req)
	if err != nil {
		return err
	}

	j.log.Infof("set %v", req.ID)
	j.log.Infof("set str %v", msg)

	return j.cache.Set(strconv.Itoa(req.ID), msg, j.ttl)
}
