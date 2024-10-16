package medium

import "sort"

// Time complexity: O(n^2)
// Space caomplexity: O(1)
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)

	for i := range nums {
		if nums[i] > 0 {
			return result
		}

		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}

				for l < r && nums[r] == nums[r-1] {
					r--
				}

				l++
				r--
			}
		}
	}

	return result
}

// [-4,-1,-1,-1,0,1,2,2]

// [-1,0,1,2,-1,-4]

// [-4,-1,-1,0,1,2]
// 	 ^  ^        ^  -> -3

// [-4,-1,-1,0,1,2]
// 	 ^     ^     ^  -> -3

// [-4,-1,-1,0,1,2]
// 	 ^       ^   ^  -> -2

// [-4,-1,-1,0,1,2]
// 	 ^     	   ^ ^  -> -1

// [-4,-1,-1,0,1,2]
// 	    ^  ^     ^  -> 0 [-1,-1,2]

// [-4,-1,-1,0,1,2]
// 	    ^    ^   ^  -> 1

// [-4,-1,-1,0,1,2]
// 	    ^    ^ ^    -> 0 [-1,0,1]

// [-4,-1,-1,0,1,2]
// 	    ^    ^ ^    -> 0 [-1,0,1]
