package easy

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Positive; in the start #1", args: args{haystack: "sadbutsad", needle: "sad"}, want: 0},
		{name: "Negative", args: args{haystack: "leetcode", needle: "leeto"}, want: -1},
		{name: "Positive; in the middle", args: args{haystack: "leetcleetode", needle: "leeto"}, want: 5},
		{name: "Positive; in the end", args: args{haystack: "leetcdeleeto", needle: "leeto"}, want: 7},
		{name: "Positive; in the start #2", args: args{haystack: "leeleeleelee", needle: "lee"}, want: 0},
		{name: "Positive; in the middle #2", args: args{haystack: "aleeleeleelee", needle: "lee"}, want: 1},
		{name: "Positive; tricky", args: args{haystack: "mississippi", needle: "issip"}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
