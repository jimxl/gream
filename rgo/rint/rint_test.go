package rint

import (
	"testing"
)

func TestFdiv(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"除法得到浮点数", args{20, 5}, 4.0},
		{"除法得到浮点数", args{20, 3}, 6.666666666666667},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fdiv(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Fdiv() = %v, want %v", got, tt.want)
			}
		})
	}
}
