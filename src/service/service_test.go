package service

import (
	"main/src/repo"
	"reflect"
	"testing"
	"time"
)

func Test_daysIn(t *testing.T) {
	type args struct {
		ym string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				ym: "202202",
			},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MonthDays(tt.args.ym); got != tt.want {
				t.Errorf("daysIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllDays(t *testing.T) {
	type args struct {
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name     string
		args     args
		wantDays []string
	}{
		{
			name: "test",
			args: args{
				start: time.Date(2022, 10, 1, 0, 0, 0, 0, time.Local),
				end:   time.Date(2022, 10, 3, 0, 0, 0, 0, time.Local),
			},
			wantDays: []string{
				"20221001", "20221002", "20221003",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDays := GetAllDays(tt.args.start, tt.args.end); !reflect.DeepEqual(gotDays, tt.wantDays) {
				t.Errorf("GetAllDays() = %v, want %v", gotDays, tt.wantDays)
			}
		})
	}
}

func TestStr(t *testing.T) {
	str := "20221230"
	t.Log(str[:6])
}

func TestBucketService_Query(t *testing.T) {
	type fields struct {
		repo repo.IBudgetRepo
	}
	type args struct {
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantBudget repo.Amount
	}{
		{
			name: "start > end",
			fields: fields{
				repo: &repo.BudgetRepo{
					Data: map[repo.YearMonth]repo.BudgetData{
						"202210": {
							YearMonth: "202210",
							Amount:    310,
						},
					},
				},
			},
			args: args{
				start: time.Date(2022, 10, 3, 0, 0, 0, 0, time.Local),
				end:   time.Date(2022, 10, 2, 0, 0, 0, 0, time.Local),
			},
			wantBudget: 0,
		},
		{
			name: "",
			fields: fields{
				repo: &repo.BudgetRepo{
					Data: map[repo.YearMonth]repo.BudgetData{
						"202210": {
							YearMonth: "202210",
							Amount:    310,
						},
					},
				},
			},
			args: args{
				start: time.Date(2022, 10, 1, 0, 0, 0, 0, time.Local),
				end:   time.Date(2022, 10, 3, 0, 0, 0, 0, time.Local),
			},
			wantBudget: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BucketService{
				repo: tt.fields.repo,
			}
			if gotBudget := b.Query(tt.args.start, tt.args.end); !reflect.DeepEqual(gotBudget, tt.wantBudget) {
				t.Errorf("BucketService.Query() = %v, want %v", gotBudget, tt.wantBudget)
			}
		})
	}
}

func TestAddDay(t *testing.T) {
	_time := time.Date(2022, 10, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, 2)
	t.Log(_time.Format("20060102"))
}
