package rstring

import "testing"

func TestCapitalize(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Cat to Cat", args{"Cat"}, "Cat"},
		{"cat to Cat", args{"cat"}, "Cat"},
		{"cat eat to Cat Eat", args{"cat eat"}, "Cat Eat"},
		{"AdbD to Adbd", args{"AdbD"}, "Adbd"},
		{"adbD to Adbd", args{"adbD"}, "Adbd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Capitalize(tt.args.str); got != tt.want {
				t.Errorf("Capitalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
