package МТС

import (
	"fmt"
	"sync"
)

// Реализуйте очередь с ограниченным размером, которая использует sync.Map
// для хранения элементов. Каждый элемент имеет свой уникальный ID, и когда
// очередь достигает максимального размера, новый элемент должен заменять старый.

type Queue struct {
	ID      int   // Следующий ID
	order   []int // Порядок добавления ID
	m       sync.Mutex
	store   sync.Map // ID -> значение
	maxSize int
}

func (q *Queue) Add(value string) (id int) {
	q.m.Lock()
	defer q.m.Unlock()

	id = q.ID
	q.ID++

	q.store.Store(id, value)

	// Если превышен размер — удаляем самый старый элемент
	if len(q.order) > q.maxSize {
		oldestID := q.order[0]
		q.store.Delete(oldestID)
		q.order = q.order[1:]
	}

	return id
}

func (q *Queue) Get(id int) (value string, ok bool) {
	v, ok := q.store.Load(id)
	if ok {
		value = v.(string)
	}
	return
}

func main() {
	q := Queue{}
	aId := q.Add("A")
	bId := q.Add("B")
	cId := q.Add("C") // должен вытеснить A

	fmt.Println("IDs:", aId, bId, cId)
}
