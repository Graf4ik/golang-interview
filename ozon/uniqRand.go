package main

import (
	"fmt"
	"math/rand"
	"time"
)

func uniqRand(n int) []int {
	if n < 0 {
		return []int{}
	}

	res := make([]int, 0, n)
	uniq := make(map[int]bool)

	rand.Seed(time.Now().UnixNano())

	for len(res) < n {
		num := rand.Intn(100)
		if !uniq[num] {
			uniq[num] = true
			res = append(res, num)
		}
	}
	return res
}

func main() {
	randNum := uniqRand(5)
	fmt.Println(randNum)
}
