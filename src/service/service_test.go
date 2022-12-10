package service

import "testing"

func Test_daysIn(t *testing.T) {
	type args struct {
		ym string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				ym: "202206",
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
