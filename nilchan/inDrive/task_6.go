package inDrive

import (
	"fmt"
	"sync"
	"time"
)

// Есть 3 горутины, в которых заблокирован запуск run
// Как сделать чтобы функция steady по команде одновременно разблокироваала выполенение всех трех run?

var (
	mu   sync.Mutex
	cond = sync.NewCond(&mu)
)

// var start = make(chan struct{})

func main() {
	go ready(1)
	go ready(2)
	go ready(3)

	time.Sleep(1 * time.Second) // имитация подготовки
	steady()
	time.Sleep(1 * time.Second) // ожидание завершения
}

// Альтернатива: sync.WaitGroup + channel (если не хочется sync.Cond)

func steady() {
	fmt.Println("steady: готов разблокировать run()")
	cond.L.Lock()
	cond.Broadcast() // пробуждает всех
	cond.L.Unlock()

	// 	close(start) // все горутины увидят закрытие канала и продолжат
}

func ready(id int) {
	run(id)
}

func run(id int) {
	cond.L.Lock()
	cond.Wait() // блокируемся до сигнала
	cond.L.Unlock()

	// 	<-start
}

/*
sync.Cond в Go — это условная переменная, которая позволяет горутинам:
ожидать наступления определённого события (Wait()),
оповещать одну (Signal()) или все (Broadcast()) ожидающие горутины о том, что событие произошло.
Она похожа на комбинацию mutex + signal/wait, как в других языках.
📌 Когда использовать sync.Cond
Когда у вас много горутин ждут одного события (например, старта работы, обновления данных, завершения ресурса) — и нужно эффективно разбудить их все или по одной.
*/
