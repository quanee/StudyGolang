package data

import (
	"blog/internal/biz"
	"blog/internal/data/ent/summary"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type summaryRepo struct {
	data *Data
	log  *log.Helper
}

// NewsummaryRepo .
func NewSummaryRepo(data *Data, logger log.Logger) biz.SummaryRepo {
	return &summaryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ar *summaryRepo) CreateSummary(ctx context.Context, summary *biz.Summary) (int64, error) {
	fmt.Println(summary)
	r, err := ar.data.db.Summary.
		Create().
		SetTitle(summary.Title).
		Save(ctx)
	if err != nil {
		return -1, err
	}
	return int64(r.ID), nil
}

func (ar *summaryRepo) UpdateSummary(ctx context.Context, id int64, summary *biz.Summary) error {
	p, err := ar.data.db.Summary.Get(ctx, int(id))
	if err != nil {
		return err
	}
	_, err = p.Update().
		SetTitle(summary.Title).
		Save(ctx)
	return err
}

func (ar *summaryRepo) GetSummary(ctx context.Context, id int64) (*biz.Summary, error) {
	p, err := ar.data.db.Summary.Get(ctx, int(id))
	if err != nil {
		return nil, err
	}

	return &biz.Summary{
		Id:    int64(p.ID),
		Title: p.Title,
	}, nil
}

func (ar *summaryRepo) DeleteSummary(ctx context.Context, id int64) error {
	return ar.data.db.Summary.DeleteOneID(int(id)).Exec(ctx)
}

func (ar *summaryRepo) ListSummary(ctx context.Context, start_id, limit int64) ([]*biz.Summary, error) {
	fmt.Println("list summary ##################################")
	summ, err := ar.data.db.Summary.Query().Where(
		summary.IDGTE(int(start_id))).Limit(int(limit)).All(ctx)
	summaries := make([]*biz.Summary, 0, len(summ))

	for _, sum := range summ {
		summaries = append(summaries, &biz.Summary{Id: int64(sum.ID), Title: sum.Title})
	}
	return summaries, err
}
