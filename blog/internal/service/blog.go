package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	pb "blog/api/blog/v1"
	"blog/internal/biz"
	"blog/internal/cache"
	"blog/internal/cache/redis"
	"blog/internal/mq"
	"blog/internal/mq/kafka"
)

type BlogService struct {
	pb.UnimplementedBlogServer

	log   *log.Helper
	summary *biz.SummaryUsecase
	article *biz.ArticleUsecase
	cache cache.Cache
	mq    mq.MessageQueue
}

func NewBlogService(summary *biz.SummaryUsecase, article *biz.ArticleUsecase, cache *redis.RadixRC3, mq *kafka.KafkaClient, logger log.Logger) *BlogService {
	return &BlogService{
		summary: summary,
		article: article,
		cache: cache,
		mq:    mq,
		log:   log.NewHelper(logger),
	}
}

func (s *BlogService) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogReply, error) {
	s.log.Infof("create blog req %v", req)
	id, err := s.summary.Create(ctx, &biz.Summary{
		Title: req.Title,
	})
	_, err = s.article.Create(ctx, &biz.Article{
		Id:      id,
		Content: req.Content,
	})
	return &pb.CreateBlogReply{
		Id:    id,
		Title: req.Title,
		Content: req.Content,
	}, err
}

func (s *BlogService) UpdateBlog(ctx context.Context, req *pb.UpdateBlogRequest) (*pb.UpdateBlogReply, error) {
	s.log.Infof("update blog req %v", req)
	err := s.article.Update(ctx, req.Id, &biz.Article{
		Id:    req.Id,
		Content: req.Content,
	})
	err = s.summary.Update(ctx, req.Id, &biz.Summary{
		Id:    req.Id,
		Title: req.Title,
	})
	return &pb.UpdateBlogReply{
		Id: req.Id,
		Title: req.Title,
		Content: req.Content,
	}, err
}

func (s *BlogService) DeleteBlog(ctx context.Context, req *pb.DeleteBlogRequest) (*pb.DeleteBlogReply, error) {
	s.log.Infof("delete blog req %v", req)
	err := s.summary.Delete(ctx, req.Id)
	err = s.article.Delete(ctx, req.Id)
	return &pb.DeleteBlogReply{
		Id: req.Id,
	}, err
}

func (s *BlogService) GetBlog(ctx context.Context, req *pb.GetBlogRequest) (*pb.GetBlogReply, error) {
	s.log.Infof("get blog req %v", req)
	sum, err := s.summary.Get(ctx, req.Id)
	art, err := s.article.Get(ctx, req.Id)
	if sum != nil && err == nil {
		return &pb.GetBlogReply{
			Id:    sum.Id,
			Title: sum.Title,
			Content: art.Content,
		}, nil
	}
	return &pb.GetBlogReply{}, err
}
