package medium

// Time complexity: O(n)
// Space complexity: O(n)
func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures)/2)

	pop := func() int {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		return top
	}

	for i, t := range temperatures {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] < t {
			prev := pop()
			answer[prev] = i - prev
		}

		stack = append(stack, i)
	}

	return answer
}
