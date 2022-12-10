// It returns the budget of a given time period.
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

const YM = "200601"

// Query implements IBucketService
func (b *BucketService) Query(start, end time.Time) (budget int64) {
	if start.After(end) || len(b.repo.GetAll()) == 0 {
		return 0
	}
	var lst []repo.BudgetData
	for _, v := range b.repo.GetAll() {
		_ym, _ := time.Parse(YM, string(v.YearMonth))
		if inMonth(_ym, start, end) {
			lst = append(lst, v)
		}
	}

	// for _, v := range lst {

	// }

	return 0
}

func MonthDays(ym string) int {
	t, err := time.Parse(YM, ym)
	if err != nil {
		panic(err)
	}
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.Local).Day()
}

func GetAllDays(start, end time.Time) (days []string) {
	for t := start; t.After(start) && t.Before(end); t.AddDate(0, 0, 1) {
		days = append(days, t.Format("20060102"))
	}
	return
}

func inMonth(target, start, end time.Time) bool {
	if start.After(target) && end.Before(target) {
		return true
	}
	if target.Year() == start.Year() && target.Month() == start.Month() {
		return true
	}
	if target.Year() == end.Year() && target.Month() == end.Month() {
		return true
	}
	return false
}

func NewBucketService(r repo.IBudgetRepo) IBucketService {
	return &BucketService{
		repo: r,
	}
}
