package medium

// Time complexity: O(nlogn)
// Space complexity: O(n)
func sortArray(nums []int) []int {
	empty := make([]int, len(nums))
	mergeSort(nums, 0, len(nums)-1, empty)

	return nums
}

func mergeSort(nums []int, l int, r int, temp []int) {
	if l < r {
		m := (l + r) / 2
		mergeSort(nums, l, m, temp)
		mergeSort(nums, m+1, r, temp)
		merge(nums, l, m+1, r, temp)
	}
}

func merge(nums []int, aFirstIdx int, bFirstIdx int, bLastIdx int, temp []int) {
	i, j := aFirstIdx, bFirstIdx
	k := i

	for i < bFirstIdx && j <= bLastIdx {
		if nums[i] <= nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
		k++
	}

	for i < bFirstIdx {
		temp[k] = nums[i]
		i++
		k++
	}

	for j <= bLastIdx {
		temp[k] = nums[j]
		j++
		k++
	}

	for i := aFirstIdx; i < bLastIdx+1; i++ {
		nums[i] = temp[i]
	}
}
