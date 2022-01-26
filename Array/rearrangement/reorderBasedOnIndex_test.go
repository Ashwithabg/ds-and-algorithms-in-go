package rearrangement

import (
	"reflect"
	"testing"
)

func Test_reorder(t *testing.T) {
	type args struct {
		arr     []int
		indices []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "reorder based on indices",
			args: args{
				arr:     []int{10, 11, 12},
				indices: []int{1, 0, 2},
			},
			want: []int{11, 10, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorder(tt.args.arr, tt.args.indices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
