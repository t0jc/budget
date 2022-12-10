package service

import (
	"main/src/repo"
	"time"
)

type IBucketService interface {
	Query(start, end time.Time) int64
}

type BucketService struct {
	repo repo.IBudgetRepo
}

const YM = "200602"

// Query implements IBucketService
func (b *BucketService) Query(start, end time.Time) (budget int64) {
	if start.After(end) || len(b.repo.GetAll()) == 0 {
		return 0
	}
	var lst []repo.BudgetData
	for _, v := range b.repo.GetAll() {
		
	}
	return 0
}

func MonthDays(ym string) int {
	t, err := time.Parse(YM, ym)
	if err != nil {
		panic(err)
	}
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.Local)..Day()
}

func NewBucketService(r repo.IBudgetRepo) IBucketService {
	return &BucketService{
		repo: r,
	}
}
