package medium

// Time: O(n)
// Space: O(1)
func numOfSubarrays(arr []int, k int, threshold int) int {
	threshold *= k // to avoid dividing by k every round

	var meetsThreshold, begin, sum int

	for end := range arr {
		sum += arr[end]
		if end-begin+1 == k {
			if sum >= threshold {
				meetsThreshold++
			}
			sum -= arr[begin]
			begin++
		}
	}

	return meetsThreshold
}
