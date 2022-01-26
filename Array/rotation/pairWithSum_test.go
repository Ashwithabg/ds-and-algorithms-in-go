package rotation

import "testing"

func Test_exitsPairWithSum(t *testing.T) {
	type args struct {
		numbers []int
		sum     int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "return true when sum exists",
			args: args{
				numbers: []int{11, 15, 6, 8, 9, 10},
				sum:     16,
			},
			want: true,
		},
		{
			name: "return true when sum exists at the beginning",
			args: args{
				numbers: []int{5, 6, 1, 2, 3, 4},
				sum:     7,
			},
			want: true,
		},
		{
			name: "return false when sum does not exists",
			args: args{
				numbers: []int{11, 15, 6, 8, 9, 10},
				sum:     12,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exitsPairWithSum(tt.args.numbers, tt.args.sum); got != tt.want {
				t.Errorf("exitsPairWithSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
