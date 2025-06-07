package leetcode

func main() {
	var arr = []int{2, 7, 11, 15}
	twoSum(arr, 9)
}

func twoSum(nums []int, target int) []int {
	mapRes := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]

		if val, exists := mapRes[diff]; exists {
			return []int{val, i}
		}
		mapRes[nums[i]] = i
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	// Create a map to store numbers and their corresponding indices
	numToIndexMap := make(map[int]int)

	// Loop through the array
	for i, num := range nums {
		// Calculate the difference between the target and the current number
		diff := target - num

		// Check if the difference already exists in the map
		if idx, found := numToIndexMap[diff]; found {
			// If it exists, return the indices of the current number and the number that adds up to the target
			return []int{i, idx}
		}

		// If it doesn't exist, add the current number and its index to the map
		numToIndexMap[num] = i
	}

	// If no two numbers add up to the target, return nil
	return nil
}
