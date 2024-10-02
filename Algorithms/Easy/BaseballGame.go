package easy

import "strconv"

// Time: O(n)
// Space: O(n)
func calPoints(operations []string) int {
	stack := make([]int, 0)

	for _, op := range operations {
		switch op {
		case "+":
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = append(stack, a+b)
		case "D":
			a := stack[len(stack)-1]
			stack = append(stack, a*2)
		case "C":
			stack = stack[:len(stack)-1]
		default:
			num, _ := strconv.Atoi(op)
			stack = append(stack, num)
		}
	}

	return sum(stack)
}

func sum(nums []int) int {
	var result int
	for _, elem := range nums {
		result += elem
	}

	return result
}
