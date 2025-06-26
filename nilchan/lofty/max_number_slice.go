package lofty

import "fmt"

// найти макс. число в слайсе
func main() {
	sl := []int{1, 2, 3, 4, 5, 6}

	maxNumberCalc(sl)
	fmt.Println(maxNumberCalc(sl))
}

func maxNumberCalc(sl []int) int {
	maxNumber := 0
	for _, val := range sl {
		if val > maxNumber {
			maxNumber = val
		}
	}

	return maxNumber
}
