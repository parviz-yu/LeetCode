package medium

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

type BrowserHistory struct {
	history []string
	ptr     int
	len     int
}

func NewBrowserHistory(homepage string) BrowserHistory {
	return BrowserHistory{
		history: []string{homepage},
		ptr:     0,
		len:     1,
	}
}

// Time Complexity: O(1)
func (this *BrowserHistory) Visit(url string) {
	if len(this.history) < this.ptr+2 {
		this.history = append(this.history, url)
	} else {
		this.history[this.ptr+1] = url
	}
	this.ptr++
	this.len = this.ptr + 1
}

// Time Complexity: O(1)
func (this *BrowserHistory) Back(steps int) string {
	this.ptr = helpers.Max(0, this.ptr-steps)
	return this.history[this.ptr]
}

// Time Complexity: O(1)
func (this *BrowserHistory) Forward(steps int) string {
	this.ptr = helpers.Min(this.len-1, this.ptr+steps)
	return this.history[this.ptr]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
