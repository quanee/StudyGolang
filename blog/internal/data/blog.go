package data

import (
	"blog/internal/biz"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type blogRepo struct {
	data *Data
	log  *log.Helper
}

// NewblogRepo .
func NewBlogRepo(data *Data, logger log.Logger) biz.BlogRepo {
	return &blogRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ar *blogRepo) CreateBlog(ctx context.Context, blog *biz.Blog) (int64, error) {
	fmt.Println(blog)
	r, err := ar.data.db.Blog.
		Create().
		SetTitle(blog.Title).
		Save(ctx)
	if err != nil {
		return -1, err
	}
	return int64(r.ID), nil
}

func (ar *blogRepo) UpdateBlog(ctx context.Context, id int64, blog *biz.Blog) error {
	p, err := ar.data.db.Blog.Get(ctx, int(id))
	if err != nil {
		return err
	}
	_, err = p.Update().
		SetTitle(blog.Title).
		Save(ctx)
	return err
}

func (ar *blogRepo) GetBlog(ctx context.Context, id int64) (*biz.Blog, error) {
	p, err := ar.data.db.Blog.Get(ctx, int(id))
	if err != nil {
		return nil, err
	}

	return &biz.Blog{
		Id:    int64(p.ID),
		Title: p.Title,
	}, nil
}

func (ar *blogRepo) DeleteBlog(ctx context.Context, id int64) error {
	return ar.data.db.Blog.DeleteOneID(int(id)).Exec(ctx)
}
