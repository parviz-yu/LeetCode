package easy

// Time complexity: O(n)
// Space complexity: O(1)
func sortedSquares_1(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	l, r := 0, n-1
	for i := n - 1; i > -1; i-- {
		if nums[l]*nums[l] > nums[r]*nums[r] {
			res[i] = nums[l] * nums[l]
			l++
		} else {
			res[i] = nums[r] * nums[r]
			r--
		}
	}

	return res
}

// Time complexity: O(n)
// Space complexity: O(1)
func sortedSquares_2(nums []int) []int {
	n := len(nums)
	res := make([]int, 0, n)

	// Obtaining the last element less than or equal to 0 using binary search
	l, r := 0, n-1
	for l < r {
		m := (l + r + 1) / 2
		if nums[m] <= 0 {
			l = m
		} else {
			r = m - 1
		}
	}
	r = l + 1

	for i := range nums {
		nums[i] *= nums[i]
	}

	for l > -1 && r < n {
		if nums[l] < nums[r] {
			res = append(res, nums[l])
			l--
		} else {
			res = append(res, nums[r])
			r++
		}
	}

	for l > -1 {
		res = append(res, nums[l])
		l--
	}

	for r < n {
		res = append(res, nums[r])
		r++
	}

	return res
}
