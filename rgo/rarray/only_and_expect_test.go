package rarray

import (
	"reflect"
	"testing"
)

func TestOnlyAndExpect(t *testing.T) {
	type args struct {
		only   []string
		expect []string
		all    []string
	}
	tests := []struct {
		name        string
		args        args
		wantResults map[string]bool
	}{
		{
			"only 和 expect都为空",
			args{
				only:   []string{},
				expect: []string{},
				all:    []string{"abc", "123"},
			},
			map[string]bool{
				"abc": true,
				"123": true,
			},
		},
		{
			"only 存在, expect为空",
			args{
				only:   []string{"abc"},
				expect: []string{},
				all:    []string{"abc", "123"},
			},
			map[string]bool{
				"abc": true,
				"123": false,
			},
		},
		{
			"only 为空, expect存在",
			args{
				only:   []string{},
				expect: []string{"abc"},
				all:    []string{"abc", "123"},
			},
			map[string]bool{
				"abc": false,
				"123": true,
			},
		},
		{
			"only 和 expect都存在",
			args{
				only:   []string{"abc", "ttt"},
				expect: []string{"abc"},
				all:    []string{"abc", "123", "ttt"},
			},
			map[string]bool{
				"abc": false,
				"123": false,
				"ttt": true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResults := OnlyAndExpect(tt.args.only, tt.args.expect, tt.args.all); !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("OnlyAndExpect() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}
