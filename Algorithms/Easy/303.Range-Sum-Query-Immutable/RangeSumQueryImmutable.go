package easy

type NumArray struct {
	arr []int
}

// Time complexity: O(n)
func Constructor(nums []int) NumArray {
	res := NumArray{
		arr: make([]int, len(nums)+1),
	}

	for i := range nums {
		res.arr[i+1] = res.arr[i] + nums[i]
	}

	return res
}

// Time complexity: O(1)
func (this *NumArray) SumRange(left int, right int) int {
	return this.arr[right+1] - this.arr[left]
}
