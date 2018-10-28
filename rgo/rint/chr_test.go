package rint

import "testing"

func TestChr(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"97 to a", args{97}, "a"},
		{"122 to z", args{122}, "z"},
		{"65 to A", args{65}, "A"},
		{"90 to Z", args{90}, "Z"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chr(tt.args.i); got != tt.want {
				t.Errorf("Chr() = %v, want %v", got, tt.want)
			}
		})
	}
}
