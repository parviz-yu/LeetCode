package medium

// Time complexity: O(n)
// Space complexity: O(n)
func rotate_2(nums []int, k int) {
	k %= len(nums)
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)

	for i := 0; i < len(nums); i++ {
		newI := (i + k) % len(nums)
		nums[newI] = numsCopy[i]
	}
}

// Time complexity: O(n)
// Space complexity: O(1)
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	rotateArray(nums, 0, n-k-1)
	rotateArray(nums, n-k, n-1)
	rotateArray(nums, 0, n-1)
}

func rotateArray(arr []int, l int, r int) {
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
