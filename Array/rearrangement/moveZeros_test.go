package rearrangement

import (
	"reflect"
	"testing"
)

func Test_moveZeros(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "moves all zeros to the end",
			args: args{
				arr: []int{1, 2, 0, 4, 3, 0, 5, 0},
			},
			want: []int{1, 2, 4, 3, 5, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveZeros(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveZeros() = %v, want %v", got, tt.want)
			}
		})
	}
}
