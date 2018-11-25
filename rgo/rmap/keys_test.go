package rmap

import (
	"reflect"
	"testing"
)

func TestKeys(t *testing.T) {
	type args struct {
		m map[interface{}]interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			"能获取map的所有key",
			args{map[interface{}]interface{}{"abc": "111", 123: "hello"}},
			[]interface{}{"abc", 123},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Keys(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}
