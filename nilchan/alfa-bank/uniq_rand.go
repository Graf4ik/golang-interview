package alfaBank

import (
	"fmt"
	"math/rand"
)

func main() {
	uniqRand(10)
	fmt.Println(uniqRand(10))
}

func uniqRand(num int) []int {
	res := make([]int, 0, num)
	unique := make(map[int]bool)

	for len(res) < num {
		val := rand.Intn(num)
		_, ok := unique[val]

		if !ok {
			unique[val] = true
			res = append(res, val)
		}
	}

	return res
}
