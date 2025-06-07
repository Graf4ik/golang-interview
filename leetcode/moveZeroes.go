package leetcode

func main() {
	var arr = []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
}

func moveZeroes(nums []int) {
	left := 0
	right := 0

	for left < len(nums) {
		if nums[left] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			right += 1
		}
		left += 1
	}
}
