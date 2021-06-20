package dao

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/quanee/go-training/week4/internal/model"
	"github.com/spf13/viper"
)

var ErrRecordNotFound = errors.New("Not Found")

type Dao interface {
	GetArticle(ctx context.Context, id int) (*model.Article, error)
}

type dao struct {
	db *sql.DB
}

func (d *dao) GetArticle(ctx context.Context, id int) (*model.Article, error) {
	article := &model.Article{}
	row := d.db.QueryRowContext(ctx, "select id,title,content from ss_articles where id=?", id)
	err := row.Scan(&article.Id, &article.Title, &article.Content)
	if err == sql.ErrNoRows {
		return nil, errors.Wrap(ErrRecordNotFound, "No corresponding article")
	}
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get article")
	}
	return article, nil
}

func NewDao(db *sql.DB) Dao {
	return &dao{db: db}
}

func NewDB() (db *sql.DB, cleanup func(), err error) {
	db, err = sql.Open("mysql", viper.GetString("mysql.dsn"))
	cleanup = func() {
		if err == nil {
			db.Close()
		}
	}
	return
}

var Provider = wire.NewSet(NewDB, NewDao)
