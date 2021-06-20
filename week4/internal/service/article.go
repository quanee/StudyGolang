package service

import (
	"context"
	"errors"

	"github.com/google/wire"
	pb "github.com/quanee/go-training/week4/api/article/v1"
	"github.com/quanee/go-training/week4/internal/dao"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	pb.UnimplementedArticleServer
	dao dao.Dao
}

func NewService(d dao.Dao) *Service {
	return &Service{dao: d}
}

func (s *Service) GetArticle(ctx context.Context, req *pb.ArticleRequest) (*pb.ArticleReply, error) {
	article, err := s.dao.GetArticle(ctx, int(req.Id))
	if err != nil {
		if errors.Is(err, dao.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "Object Not Found")
		}
		return nil, status.Errorf(codes.Internal, "Error:%v", err)
	}
	return &pb.ArticleReply{Id: int64(article.Id), Title: article.Title, Content: article.Content}, nil
}

var Provider = wire.NewSet(NewService, dao.Provider)
