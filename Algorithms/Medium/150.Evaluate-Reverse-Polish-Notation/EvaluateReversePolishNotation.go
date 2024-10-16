package medium

import "strconv"

// Time complexity: O(n)
// Space complexity: O(n)
func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	pop := func() int {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return top
	}

	opers := map[string]func(a, b int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	for _, elem := range tokens {
		if v, ok := opers[elem]; ok {
			oper2 := pop()
			oper1 := pop()
			stack = append(stack, v(oper1, oper2))
			continue
		}

		num, _ := strconv.Atoi(elem)
		stack = append(stack, num)
	}
	return stack[len(stack)-1]
}
