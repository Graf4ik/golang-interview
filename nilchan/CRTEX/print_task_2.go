package CRTEX

import "fmt"

// Что выведет программа?

type impl struct{}

type I interface {
	C()
}

func (*impl) C() {}

func A() I {
	return nil // Явно nil интерфейс
}

func B() I {
	var ret *impl // Это nil указатель, но тип ret — *impl
	return ret    // Преобразуется в интерфейс I со значением (type=*impl, value=nil)
}

func main() {
	a := A()            // interface{} = nil (и type == nil, и value == nil)
	b := B()            // interface{} = (type = *impl, value = nil)
	fmt.Println(a == b) // false
}
