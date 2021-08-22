package data

import (
	"blog/internal/biz"
	"blog/internal/data/ent"
	"blog/internal/data/ent/article"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

// NewarticleRepo .
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ar *articleRepo) Pricing(ctx context.Context, bizArticle *biz.Article) (*biz.Article, error) {
	p, err := ar.data.db.Article.
		Query().
		Where(
			article.And(
				article.IDEQ(int(bizArticle.Id)))).
		First(ctx)

	if err != nil {
		return nil, err
	}
	return &biz.Article{
		Id:      int64(p.ID),
		Content: p.Content,
	}, nil
}

func (ar *articleRepo) CreateArticle(ctx context.Context, article *biz.Article) (int64, error) {

	p, err := ar.Pricing(ctx, article)
	if err != nil && !ent.IsNotFound(err) {
		return -1, err
	}

	if p != nil {
		return -1, errors.New("data: create to article error")
	}

	r, err := ar.data.db.Article.
		Create().
		SetContent(article.Content).
		Save(ctx)
	if err != nil {
		return -1, err
	}
	return int64(r.ID), nil
}

func (ar *articleRepo) UpdateArticle(ctx context.Context, id int64, article *biz.Article) error {
	p, err := ar.data.db.Article.Get(ctx, int(id))
	if err != nil {
		return err
	}
	_, err = p.Update().
		SetContent(article.Content).
		Save(ctx)
	return err
}

func (ar *articleRepo) GetArticle(ctx context.Context, id int64) (*biz.Article, error) {
	p, err := ar.data.db.Article.Get(ctx, int(id))
	if err != nil {
		return nil, err
	}

	return &biz.Article{
		Id:      int64(p.ID),
		Content: p.Content,
	}, nil
}

func (ar *articleRepo) DeleteArticle(ctx context.Context, id int64) error {
	return ar.data.db.Article.DeleteOneID(int(id)).Exec(ctx)
}
