package medium

type tuple struct {
	idx  int
	char byte
}

// Time: O(n)
// Space: O(n)
func canChange(start string, target string) bool {
	sQueue := make([]tuple, 0)
	tQueue := make([]tuple, 0)

	for i := range start {
		if start[i] != '_' {
			sQueue = append(sQueue, tuple{i, start[i]})
		}
		if target[i] != '_' {
			tQueue = append(tQueue, tuple{i, target[i]})
		}
	}

	if len(sQueue) != len(tQueue) {
		return false
	}

	var sT, tT tuple
	for len(sQueue) != 0 && len(tQueue) != 0 {
		sT, sQueue = sQueue[0], sQueue[1:]
		tT, tQueue = tQueue[0], tQueue[1:]

		if sT.char != tT.char {
			return false
		}
		if sT.char == 'L' && tT.idx > sT.idx {
			return false
		}
		if sT.char == 'R' && tT.idx < sT.idx {
			return false
		}
	}

	return true
}
