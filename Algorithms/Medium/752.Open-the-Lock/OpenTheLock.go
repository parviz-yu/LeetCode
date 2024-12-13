package medium

import "strconv"

type tuple struct {
	code string
	step int
}

func openLock(deadends []string, target string) int {
	visited := make(map[string]bool, len(deadends))
	for _, d := range deadends {
		visited[d] = true
	}

	if visited["0000"] {
		return -1
	}

	queue := []tuple{{code: "0000", step: 0}}
	visited["0000"] = true

	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]

		if t.code == target {
			return t.step
		}

		for _, n := range neighbours(t.code) {
			if !visited[n] {
				queue = append(queue, tuple{code: n, step: t.step + 1})
				visited[n] = true
			}
		}
	}

	return -1
}

func neighbours(code string) []string {
	result := make([]string, 0, 8)
	for i, d := range code {
		digit, _ := strconv.Atoi(string(d))
		for _, diff := range []int{1, -1} {
			str := code[:i] + strconv.Itoa((digit+diff+10)%10) + code[i+1:]
			result = append(result, str)
		}
	}

	return result
}
