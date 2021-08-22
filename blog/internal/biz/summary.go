package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type Summary struct {
	Id    int64
	Title string
}

type SummaryRepo interface {
	CreateSummary(ctx context.Context, sum *Summary) (int64, error)
	UpdateSummary(ctx context.Context, id int64, sum *Summary) error
	GetSummary(ctx context.Context, id int64) (*Summary, error)
	DeleteSummary(ctx context.Context, id int64) error
	ListSummary(ctx context.Context, start_id, limit int64) ([]*Summary, error)
}

type SummaryUsecase struct {
	repo SummaryRepo
}

func NewSummaryUsecase(repo SummaryRepo, logger log.Logger) *SummaryUsecase {
	return &SummaryUsecase{repo: repo}
}

func (uc *SummaryUsecase) Create(ctx context.Context, sum *Summary) (int64, error) {
	return uc.repo.CreateSummary(ctx, sum)
}

func (uc *SummaryUsecase) Update(ctx context.Context, id int64, sum *Summary) error {
	return uc.repo.UpdateSummary(ctx, id, sum)
}

func (uc *SummaryUsecase) Get(ctx context.Context, id int64) (*Summary, error) {
	return uc.repo.GetSummary(ctx, id)
}

func (uc *SummaryUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteSummary(ctx, id)
}

func (uc *SummaryUsecase) ListSummary(ctx context.Context, start_id, limit int64) (p []*Summary, err error) {
	fmt.Println("########################", start_id, limit)
	return uc.repo.ListSummary(ctx, start_id, limit)
}
