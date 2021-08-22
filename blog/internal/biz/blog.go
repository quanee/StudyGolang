package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Blog struct {
	Id    int64
	Title string
	Content string
}

type BlogRepo interface {
	CreateBlog(ctx context.Context, sum *Blog) (int64, error)
	UpdateBlog(ctx context.Context, id int64, sum *Blog) error
	GetBlog(ctx context.Context, id int64) (*Blog, error)
	DeleteBlog(ctx context.Context, id int64) error
}

type BlogUsecase struct {
	repo BlogRepo
}

func NewBlogUsecase(repo BlogRepo, logger log.Logger) *BlogUsecase {
	return &BlogUsecase{repo: repo}
}

func (uc *BlogUsecase) Create(ctx context.Context, sum *Blog) (int64, error) {
	return uc.repo.CreateBlog(ctx, sum)
}

func (uc *BlogUsecase) Update(ctx context.Context, id int64, sum *Blog) error {
	return uc.repo.UpdateBlog(ctx, id, sum)
}

func (uc *BlogUsecase) Get(ctx context.Context, id int64) (*Blog, error) {
	return uc.repo.GetBlog(ctx, id)
}

func (uc *BlogUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteBlog(ctx, id)
}
