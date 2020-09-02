package utils

import "testing"

func TestIsSameDay(t *testing.T) {
	type args struct {
		curTime  int64
		lastTime int64
		restTime int64
		zone     int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"case1", args{1598997601, 1598997599, 6, 8}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSameDay(tt.args.curTime, tt.args.lastTime, tt.args.restTime, tt.args.zone); got != tt.want {
				t.Errorf("IsSameDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
