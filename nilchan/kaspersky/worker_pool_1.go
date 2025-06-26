package kaspersky

import (
	"context"
	"sync"
)

type WorkerPool struct {
	workerQty int

	tasks chan func()        // очередь задач
	wg    sync.WaitGroup     // ждёт запущенные задачи
	ctx   context.Context    // отмена по Stop / StopWait
	stop  context.CancelFunc // вызывается при Stop*
	once  sync.Once          // гарантирует однократный Stop*
}

func NewWorkerPool(numOfWorkers int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())

	wp := &WorkerPool{
		workerQty: numOfWorkers,
		tasks:     make(chan func(), 1024),
		ctx:       ctx,
		stop:      cancel,
	}

	for i := 0; i < wp.workerQty; i++ {
		go wp.worker()
	}
	return wp
}

func (wp *WorkerPool) worker() {
	for {
		select {
		case <-wp.ctx.Done():
			return // пул остановлен
		case task, ok := <-wp.tasks:
			if !ok {
				return // канал закрыт StopWait-ом
			}
			task()
			wp.wg.Done()
		}
	}
}

// Submit - добавить таску в воркер пул
func (wp *WorkerPool) Submit(task func()) {
	wp.wg.Add(1)
	wp.tasks <- task
}

// SubmitWait помещает задачу и блокируется,
// пока она не будет выполнена.
func (wp *WorkerPool) SubmitWait(task func()) {
	var innerWG sync.WaitGroup
	innerWG.Add(1)

	wp.Submit(func() {
		task()
		innerWG.Done()
	})

	innerWG.Wait()
}

// Stop - остановить воркер пул, дождаться выполнения только тех тасок, который выполняются сейчас
func (wp *WorkerPool) Stop() {
	wp.once.Do(func() {
		close(wp.tasks) // сигнал воркерам прекратить чтение
	})
	wp.wg.Wait() // ждём, пока запущенные задачи завершатся
}

// StopWait - остановить воркер пул, дождаться выполнения всех тасок, даже тех, что не начали выполняться
func (wp *WorkerPool) StopWait() {
	wp.once.Do(func() {
		close(wp.tasks) // воркеры дочитают всё и завершатся
	})
	wp.wg.Wait()
}

/*
Submit — неблокирующее добавление;
SubmitWait — синхронный вызов;
Stop — мгновенно «замораживает» очередь, завершает только активные задачи;
StopWait — обрабатывает всё до конца.
*/
