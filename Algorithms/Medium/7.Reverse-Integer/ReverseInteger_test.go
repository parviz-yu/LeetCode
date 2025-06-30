package medium

import "testing"

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{name: "positive number", args: 123, want: 321},
		{name: "negative number", args: -123, want: -321},
		{name: "zero digit", args: 0, want: 0},
		{name: "positive digit", args: 1, want: 1},
		{name: "negative digit", args: -1, want: -1},
		{name: "positive out of range", args: 1534236469, want: 0},
		{name: "negative out of range", args: -1534236469, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
