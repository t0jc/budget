// It returns the budget of a given time period.
package service

import (
	"main/src/repo"
	"time"
)

type IBucketService interface {
	Query(start, end time.Time) repo.Amount
}

type BucketService struct {
	repo repo.IBudgetRepo
}

const (
	YYYYMMDD = "20060102"
	YYYYMM   = "200601"
)

// Query implements IBucketService
func (b *BucketService) Query(start, end time.Time) (budget repo.Amount) {
	if start.After(end) || len(b.repo.GetAll()) == 0 {
		return 0
	}

	_map := make(map[repo.YearMonth]repo.Amount)
	for _, v := range b.repo.GetAll() {
		_map[v.YearMonth] = v.Amount
	}

	for _, v := range GetAllDays(start, end) {
		ymd := GetYM(v)
		days := MonthDays(ymd)
		db := _map[repo.YearMonth(ymd)] / repo.Amount(days)
		budget += db
	}
	return
}

func GetYM(ymd string) string {
	return ymd[:6]
}

func MonthDays(ym string) int {
	t, err := time.Parse(YYYYMM, ym)
	if err != nil {
		panic(err)
	}
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.Local).Day()
}

func GetAllDays(start, end time.Time) (days []string) {
	for t := start; t.Before(end.AddDate(0, 0, 1)); t = t.AddDate(0, 0, 1) {
		days = append(days, t.Format(YYYYMMDD))
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
