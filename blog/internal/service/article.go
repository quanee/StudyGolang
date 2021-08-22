package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	pb "blog/api/article/v1"
	"blog/internal/biz"
	"blog/internal/cache"
	"blog/internal/cache/redis"
	"blog/internal/mq"
	"blog/internal/mq/kafka"
)

type ArticleService struct {
	pb.UnimplementedArticleServiceServer

	log     *log.Helper
	article *biz.ArticleUsecase
	cache   cache.Cache
	mq      mq.MessageQueue
}

func NewArticleService(article *biz.ArticleUsecase, cache *redis.RadixRC3, mq *kafka.KafkaClient, logger log.Logger) *ArticleService {
	return &ArticleService{
		article: article,
		cache:   cache,
		mq:      mq,
		log:     log.NewHelper(logger),
	}
}

func (a *ArticleService) CreateArticles(ctx context.Context, req *pb.CreateArticlesRequest) (*pb.CreateArticlesReply, error) {
	a.log.Infof("create article req %v", req)
	id, err := a.article.Create(ctx, &biz.Article{
		Content: req.Content,
	})
	return &pb.CreateArticlesReply{
		Id: id,
	}, err
}

func (a *ArticleService) UpdateArticles(ctx context.Context, req *pb.UpdateArticlesRequest) (*pb.UpdateArticlesReply, error) {
	a.log.Infof("update article req %v", req)
	err := a.article.Update(ctx, req.Article.Id, &biz.Article{
		Id:      req.Article.Id,
		Content: req.Article.Content,
	})
	return &pb.UpdateArticlesReply{
		Id: req.Article.Id,
	}, err
}

func (a *ArticleService) DeleteArticles(ctx context.Context, req *pb.DeleteArticlesRequest) (*pb.DeleteArticlesReply, error) {
	a.log.Infof("delete article req %v", req)
	err := a.article.Delete(ctx, req.Id)
	return &pb.DeleteArticlesReply{
		Id: req.Id,
	}, err
}

func (a *ArticleService) GetArticles(ctx context.Context, req *pb.GetArticlesRequest) (*pb.GetArticlesReply, error) {
	a.log.Infof("get article req %v", req)
	art, err := a.article.Get(ctx, req.Id)
	return &pb.GetArticlesReply{
		Article: &pb.Article{
			Id:      art.Id,
			Content: art.Content,
		},
	}, err
}
