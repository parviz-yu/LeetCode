package helpers

func Sum(nums []int) int {
	total := 0
	for _, val := range nums {
		total += val
	}

	return total
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
