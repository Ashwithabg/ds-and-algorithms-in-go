package rotation

import "testing"

func Test_findMaxRotation(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find max rotation",
			args: args{
				arr: []int{1, 20, 2, 10},
			},
			want: 72,
		},
		{
			name: "find max rotation",
			args: args{
				arr: []int{1, 2,3},
			},
			want: 8,
		},
		{
			name: "find max rotation",
			args: args{
				arr: []int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: 330,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxRotation(tt.args.arr); got != tt.want {
				t.Errorf("findMaxRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
