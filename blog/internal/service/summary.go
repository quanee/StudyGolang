package service

import (
	"blog/internal/data/ent"
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"

	pb "blog/api/summary/v1"
	"blog/internal/biz"
	"blog/internal/cache"
	"blog/internal/cache/redis"
	"blog/internal/mq"
	"blog/internal/mq/kafka"
)

type SummaryService struct {
	pb.UnimplementedSummaryServiceServer

	log     *log.Helper
	summary *biz.SummaryUsecase
	cache   cache.Cache
	mq      mq.MessageQueue
}

func NewSummaryServiceService(summary *biz.SummaryUsecase, cache *redis.RadixRC3, mq *kafka.KafkaClient, logger log.Logger) *SummaryService {
	return &SummaryService{
		summary: summary,
		cache:   cache,
		mq:      mq,
		log:     log.NewHelper(logger),
	}
}

func (s *SummaryService) CreateSummary(ctx context.Context, req *pb.CreateSummaryRequest) (*pb.CreateSummaryReply, error) {
	s.log.Infof("create summary req %v", req)
	id, err := s.summary.Create(ctx, &biz.Summary{
		Title: req.Title,
	})

	return &pb.CreateSummaryReply{
		Summary: &pb.Summary{
			Id:    id,
			Title: req.Title,
		},
	}, err
}

func (s *SummaryService) UpdateSummary(ctx context.Context, req *pb.UpdateSummaryRequest) (*pb.UpdateSummaryReply, error) {
	s.log.Infof("update summary req %v", req)
	err := s.summary.Update(ctx, req.Id, &biz.Summary{
		Id:    req.Id,
		Title: req.Title,
	})

	res := pb.Summary{
		Id: req.Id,
		Title: req.Title,
	}

	msg, _ := json.Marshal(res)
	_ = s.mq.Produce(string(msg))
	return &pb.UpdateSummaryReply{
		Id: req.Id,
	}, err
}

func (s *SummaryService) DeleteSummary(ctx context.Context, req *pb.DeleteSummaryRequest) (*pb.DeleteSummaryReply, error) {
	s.log.Infof("delete summary req %v", req)
	err := s.summary.Delete(ctx, req.Id)
	return &pb.DeleteSummaryReply{
		Id: req.Id,
	}, err
}

func (s *SummaryService) GetSummary(ctx context.Context, req *pb.GetSummaryRequest) (*pb.GetSummaryReply, error) {
	s.log.Infof("get summary req %v", req)

	csum, err := s.cache.Get(strconv.Itoa(int(req.Id)))
	s.log.Infof("cache get summary %v", csum)

	if csum != "" && err == nil {
		s.log.Infof("cache get summary")
		res := ent.Summary{}

		err := json.Unmarshal([]byte(csum), &res)
		if err != nil {
			return nil, err
		}
		return &pb.GetSummaryReply{
			Summary: &pb.Summary{
				Id: int64(res.ID),
				Title: res.Title,
			},
		}, nil
	}

	sum, err := s.summary.Get(ctx, req.Id)
	s.log.Infof("get %v got %v", req.Id, sum)
	if sum != nil && err == nil {
		res := ent.Summary{
			ID: int(sum.Id),
			Title: sum.Title,
		}
		msg, _ := json.Marshal(res)
		_ = s.mq.Produce(string(msg))
		s.log.Infof("send kafka %v", string(msg))
		return &pb.GetSummaryReply{
			Summary: &pb.Summary{
				Id:    sum.Id,
				Title: sum.Title,
			},
		}, nil
	}
	return &pb.GetSummaryReply{
		Summary: &pb.Summary{},
	}, err
}

func (s *SummaryService) ListSummary(ctx context.Context, req *pb.ListSummaryRequest) (*pb.ListSummaryReply, error) {
	s.log.Infof("list summary req %v", req)
	sums, err := s.summary.ListSummary(ctx, req.StartId, req.Limit)
	summaries := make([]*pb.Summary, 0, len(sums))
	for _, sum := range sums {
		summaries = append(summaries, &pb.Summary{Id: sum.Id, Title: sum.Title})
	}

	return &pb.ListSummaryReply{
		Summaries: summaries,
	}, err
}
