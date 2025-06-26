package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	addNums(nums[0:2])
	fmt.Println(nums) // что выведет [1 2 3]

	addNums(nums[0:2])
	fmt.Println(nums) // что выведет [1 2 3]
}

func addNums(nums []int) {
	nums = append(nums, 5, 6)
}

/*
Первый append(nums, 5, 6) создаёт новый массив, потому что cap == 3, а ты добавляешь ещё два элемента.
Но т.к. присваивания обратно нет, оригинальный nums не меняется.
*/
