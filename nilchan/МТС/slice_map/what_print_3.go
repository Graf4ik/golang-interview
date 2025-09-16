package МТС

import (
	"fmt"
)

// Что выведет код?
func main() {
	c := []string{"A", "B", "D", "E"} // 4 4
	b := c[1:2]                       // B len=1 cap=4
	b = append(b, "TT")               // B TT cap=4
	fmt.Println(c)                    // A B TT E
	fmt.Println(b)                    // B TT
}
