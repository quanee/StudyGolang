package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Article struct {
	Id    int64
	Content string
}

type ArticleRepo interface {
	CreateArticle(ctx context.Context, article *Article) (int64, error)
	UpdateArticle(ctx context.Context, id int64, article *Article) error
	GetArticle(ctx context.Context, id int64) (*Article, error)
	DeleteArticle(ctx context.Context, id int64) error
}

type ArticleUsecase struct {
	repo ArticleRepo
}

func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo}
}

func (uc *ArticleUsecase) Create(ctx context.Context, article *Article) (int64, error) {
	return uc.repo.CreateArticle(ctx, article)
}

func (uc *ArticleUsecase) Update(ctx context.Context, id int64, article *Article) error {
	return uc.repo.UpdateArticle(ctx, id, article)
}

func (uc *ArticleUsecase) Get(ctx context.Context, id int64) (*Article, error) {
	return uc.repo.GetArticle(ctx, id)
}

func (uc *ArticleUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteArticle(ctx, id)
}
