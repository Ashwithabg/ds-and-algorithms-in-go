package rotation

import "testing"

func Test_findMinValue(t *testing.T) {
	type args struct {
		elements []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Find min value in rotated array",
			args: args{
				elements: []int{5, 6, 1, 2, 3, 4},
			},
			want: 1,
		},

		{
			name: "Find min value in rotated array",
			args: args{
				elements: []int{2, 3, 4, 5, 6, 1},
			},
			want: 1,
		},
		{
			name: "Find min value in rotated array",
			args: args{
				elements: []int{1, 2, 3, 4, 5, 6},
			},
			want: 1,
		},
		{
			name: "Find min value in rotated array",
			args: args{
				elements: []int{6, 1, 2, 3, 4, 5},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinValue(tt.args.elements); got != tt.want {
				t.Errorf("findMinValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
