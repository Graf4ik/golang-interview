package Uzum

import (
	"fmt"
	"sync"
)

type Once struct {
	done bool
	m    sync.Mutex
}

func (o *Once) Do(f func()) {
	o.m.Lock()
	defer o.m.Unlock()

	if !o.done {
		f()
		o.done = true
	}
}

func main() {
	o := Once{}

	o.Do(func() {
		fmt.Println(1)
	})

	o.Do(func() {
		fmt.Println(2)
	})

	o.Do(func() {
		fmt.Println(3)
	})
}
