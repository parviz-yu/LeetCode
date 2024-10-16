package easy

// Time complexity: O(n)
// Space complexity: O(n)
func majorityElement_1(nums []int) int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	majElem, countElem := 0, 0
	for elem, count := range counter {
		if count > countElem {
			majElem = elem
			countElem = count
		}
	}

	return majElem
}

// Time complexity: O(n)
// Space complexity: O(1)
//
// Boyer–Moore majority vote algorithm
func majorityElement_2(nums []int) int {
	majElem, counter := 0, 0
	for _, num := range nums {
		if counter == 0 {
			majElem = num
			counter = 1
		} else if majElem == num {
			counter++
		} else {
			counter--
		}
	}

	return majElem
}
