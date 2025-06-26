package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
	"unsafe"
)

type Bad struct {
	b int64
	d int32
	a int8
	c int8
}

func main() {
	fmt.Println(unsafe.Sizeof(Bad{})) // 16

	user := struct {
		balance       float64
		isTestProfile bool
		age           uint64
		isDesktop     bool
		location      struct{}
	}{}
	fmt.Println(unsafe.Sizeof(user)) // 32
	ptr := unsafe.Pointer(&user)
	fmt.Println(unsafe.Sizeof(ptr))          // 8
	fmt.Println(unsafe.Sizeof([]struct{}{})) // 24
	fmt.Println(unsafe.Sizeof([5]int64{}))   // 40
	fmt.Println(unsafe.Sizeof("h"))          // 16
	var longStr string
	for i := 0; i < 100; i++ {
		longStr += strconv.Itoa(i)
	}
	fmt.Println(unsafe.Sizeof(longStr)) // 16
	bigMap := make(map[string][]string)
	m := sync.Mutex{}
	for i := 0; i < 10; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			bigMap[time.Now().String()] = []string{time.Now().String(), time.Now().String()}
		}()
	}
	fmt.Println(unsafe.Sizeof(bigMap))                  // 8
	fmt.Println(unsafe.Sizeof(make(chan struct{}, 10))) // 8
	var mem runtime.MemStats
	runtime.GC()
	var elements [][]int64
	for i := 0; i < 500; i++ {
		elements = append(elements, getLastElem())
		runtime.GC()
		runtime.ReadMemStats(&mem)
	}
	fmt.Printf("after = %v Mib\n", mem.Alloc/1024/1024) // 4MB

	arr := [5]int{1, 2, 3}
	ptr = unsafe.Pointer(&arr[0])
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(ptr) + unsafe.Sizeof(arr[0])*3)))
}

func getLastElem() []int64 {
	elements := make([]int64, 0, 1000)
	for i := 0; i < 1000; i++ {
		elements = append(elements, int64(i))
	}
	return elements[999:]
}
