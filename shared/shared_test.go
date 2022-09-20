package shared

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				a: 42,
				b: 21,
			},
			want: 63,
		},
	}
	for _, tt := range tests {
		if got := Sum(tt.args.a, tt.args.b); got != tt.want {
			t.Errorf("Sum() = %v, want %v", got, tt.want)
		}
	}
}
