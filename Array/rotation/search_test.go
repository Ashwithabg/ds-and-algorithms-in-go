package rotation

import (
	"testing"
)

func Test_getIndexIfExists(t *testing.T) {
	type args struct {
		numbers []int
		element int
		start   int
		end     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "returns index when element exists in the middle",
			args: args{
				numbers: []int{5, 6, 7, 8, 9, 10, 1, 2, 3},
				element: 9,
				start:   0,
				end:     8,
			},
			want: 4,
		},
		{
			name: "returns index when element exists in the front",
			args: args{
				numbers: []int{5, 6, 7, 8, 9, 10, 1, 2, 3},
				element: 5,
				start:   0,
				end:     8,
			},
			want: 0,
		},
		{
			name: "returns index when element exists in the end",
			args: args{
				numbers: []int{5, 6, 7, 8, 9, 10, 1, 2, 3},
				element: 3,
				start:   0,
				end:     8,
			},
			want: 8,
		},
		{
			name: "returns -1 when element not found",
			args: args{
				numbers: []int{5, 6, 7, 8, 9, 10, 1, 2, 3},
				element: 13,
				start:   0,
				end:     8,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIndexIfExists(tt.args.numbers, tt.args.element, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("getIndexIfExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
