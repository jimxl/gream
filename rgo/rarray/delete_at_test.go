package rarray

import (
	"reflect"
	"testing"
)

func TestDeleteAt(t *testing.T) {
	type args struct {
		array []interface{}
		index int
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			"删除元素1",
			args{
				array: []interface{}{"a", "b", "c"},
				index: 1,
			},
			[]interface{}{"a", "c"},
		},
		{
			"删除元素2",
			args{
				array: []interface{}{"a", "b", "c"},
				index: 2,
			},
			[]interface{}{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteAt(tt.args.array, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
