package Ситидрайв

// 1. что выведет

type T []int

func (T) X()  {}
func (*T) Z() {}

func main() {
	var t T
	t.X() // 1
	t.Z() // 2
	var p = &t
	p.X()   // 3
	p.Z()   // 4
	T{}.X() // 5
	T{}.Z() // 6
}
