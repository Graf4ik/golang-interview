package burgerKing

import "fmt"

func main() {
	{
		appendAndModify := func(s1 []int) {
			newSlice := append(s1, 1)
			newSlice[0] = -1
		}
		slice1 := []int{1, 2, 3}
		appendAndModify(slice1) // 1 2 3
		slice2 := make([]int, 0, 10)
		slice2 = append(slice2, 1, 2, 3) // 1 2 3
		appendAndModify(slice2)          // -1 2 3
		fmt.Println(slice1, slice2)      // [1 2 3] [-1 2 3]
	}

	fmt.Println("_____")
}
