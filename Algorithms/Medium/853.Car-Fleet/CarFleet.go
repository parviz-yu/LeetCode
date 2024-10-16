package medium

import "sort"

// Time complexity: O(nlogn)
// Space complexity: O(n)
type car struct {
	position int
	speed    int
}

func carFleet(target int, position []int, speed []int) int {
	stack := make([]float64, 0, 1)
	pair := make([]car, len(position))

	for i := 0; i < len(position); i++ {
		pair[i] = car{position[i], speed[i]}
	}

	sort.Slice(pair, func(i, j int) bool {
		return pair[i].position > pair[j].position
	})

	for _, car := range pair {
		t := float64(target-car.position) / float64(car.speed)
		if len(stack) == 0 || stack[len(stack)-1] < t {
			stack = append(stack, t)
		}
	}

	return len(stack)
}
