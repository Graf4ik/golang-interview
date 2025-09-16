package alfaBank

import "fmt"

func main() {
	s := "hello"
	for i := range s {
		fmt.Printf("position %d: %c\n", i, s[i])
	}
	fmt.Printf("len=%d\n", len(s)) // 10
}
