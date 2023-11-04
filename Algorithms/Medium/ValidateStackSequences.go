package medium

// Time complexity: O(n)
// Space complexity: O(n)
func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0, len(pushed))
	i, j := 0, 0

	push := func(item int) {
		stack = append(stack, item)
	}
	pop := func() int {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		return top
	}

	for i < len(pushed) {
		if len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			pop()
			j++
		} else {
			push(pushed[i])
			i++
		}
	}

	for j < len(popped) {
		if pop() != popped[j] {
			return false
		}
		j++
	}

	return true
}
