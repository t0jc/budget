package service

import (
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
			name:     "",
			args:     args{
				start: time.Time{},
				end:   time.Time{},
			},
			wantDays: []string{},
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
