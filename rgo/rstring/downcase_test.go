package rstring

import "testing"

func TestDowncase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"转成小写", args{"Down"}, "down"},
		{"转成小写", args{"你好"}, "你好"},
		{"转成小写", args{"XXXX"}, "xxxx"},
		{"转成小写", args{""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Downcase(tt.args.str); got != tt.want {
				t.Errorf("Downcase() = %v, want %v", got, tt.want)
			}
		})
	}
}
