package leetcode

func main() {
	singleNumber([]int{1, 2, 3, 2, 3})
}

func singleNumber(nums []int) int {
	var res int

	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}

	return res
}
