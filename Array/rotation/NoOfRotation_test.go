package rotation

import "testing"

func Test_findRotationCount(t *testing.T) {
	type args struct {
		elements []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "return rotation count",
			args: args{
				elements: []int{15, 18, 2, 3, 6, 12},
			},
			want: 2	,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRotationCount(tt.args.elements); got != tt.want {
				t.Errorf("findRotationCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
