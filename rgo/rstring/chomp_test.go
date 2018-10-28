package rstring

import "testing"

func TestChomp(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"去掉左右空格", args{"   abc   "}, "abc"},
		{"去掉左右空格", args{"   你好   "}, "你好"},
		{"去掉左右不可见字符", args{"\n\r你好   "}, "你好"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chomp(tt.args.str); got != tt.want {
				t.Errorf("Chomp() = %v, want %v", got, tt.want)
			}
		})
	}
}
