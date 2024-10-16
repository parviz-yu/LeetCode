package easy

// Time complexity: O(m + n)
// Space complexity: O(m + n)
func merge_1(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, len(nums1))
	copy(temp, nums1)

	i, j, k := 0, 0, 0
	for i < m && j < n {
		if temp[i] < nums2[j] {
			nums1[k] = temp[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}

	for i < m {
		nums1[k] = temp[i]
		i++
		k++
	}

	for j < n {
		nums1[k] = nums2[j]
		j++
		k++
	}
}

// Time complexity: O(m + n)
// Space complexity: O(1)
func merge_2(nums1 []int, m int, nums2 []int, n int) {
	lastIdx := m + n - 1
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[lastIdx] = nums1[m-1]
			m--
		} else {
			nums1[lastIdx] = nums2[n-1]
			n--
		}
		lastIdx--
	}

	for n > 0 {
		nums1[lastIdx] = nums2[n-1]
		n--
		lastIdx--
	}
}
