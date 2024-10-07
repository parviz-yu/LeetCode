package medium

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time: O(n)
// Space: O(n)
func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)

	for _, aster := range asteroids {
		var collision bool

		for len(stack) > 0 && canCollide(stack[len(stack)-1], aster) {
			top := stack[len(stack)-1]
			if helpers.Abs(top) < helpers.Abs(aster) {
				stack = stack[:len(stack)-1]
				continue
			}

			if helpers.Abs(top) == helpers.Abs(aster) {
				stack = stack[:len(stack)-1]
			}

			collision = true
			break
		}

		if !collision {
			stack = append(stack, aster)
		}
	}

	return stack
}

func canCollide(a, b int) bool {
	if a > 0 && b < 0 {
		return true
	}

	return false
}
