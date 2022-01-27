package orderStatistics

import "testing"

func Test_sortAndGetElement(t *testing.T) {
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
			name: "sort and get elements",
			args: args{
				arr: []int{12, 3, 5, 7, 19},
				k:   3,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortAndGetElement(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("sortAndGetElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMinElementUsingMaxheap(t *testing.T) {
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
			name: "returns kth smallest element",
			args: args{
				arr: []int{12, 3, 5, 7, 19},
				k:   3,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinElementUsingMaxheap(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("getMinElementUsingMaxheap() = %v, want %v", got, tt.want)
			}
		})
	}
}
