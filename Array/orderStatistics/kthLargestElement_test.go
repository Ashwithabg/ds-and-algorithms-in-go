package orderStatistics

import "testing"

func Test_getLargestElement(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "return kth largest element",
			args: args{
				arr: []int{6, 7, 10, 11, 2, 1},
				k:   2,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLargestElement(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("getLargestElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
