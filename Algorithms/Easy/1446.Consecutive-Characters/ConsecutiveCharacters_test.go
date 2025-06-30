package easy

import "testing"

func Test_maxPower(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "1", s: "leetcode", want: 2},
		{name: "2", s: "aaaaaaaaaaaaa", want: 13},
		{name: "3", s: "aabbbcccc", want: 4},
		{name: "4", s: "abab", want: 1},
		{name: "5", s: "aaab", want: 3},
		{name: "6", s: "abbbbbc", want: 5},
		{name: "7", s: "abbcccddddeeeeedcba", want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPower(tt.s); got != tt.want {
				t.Errorf("maxPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
