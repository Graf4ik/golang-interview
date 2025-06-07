package avito

import (
	"fmt"
	"sync"
	"time"
)

/*
Разработать систему распределённой обработки задач с учетом равномерного нагрузки между обработчиками.

Представьте, что у вас есть система, которая получает поток задач для обработки.
Каждая задача требует некоторого времени для выполнения, и вы хотите, чтобы задачи обрабатывались параллельно,
 но с учетом равномерного рапредления нагрузки между обработчками
*/

/*
✅ Цели
Обрабатывать задачи параллельно
Распределять задачи равномерно между обработчиками
Гарантировать, что обработчики не перегружены
Обеспечить расширяемость (добавление новых обработчиков)

🏗 Архитектура (высокоуровневая)
                  +---------------+
   [Источник задач] ---> [Dispatcher]
                  +---------------+
                          |
                +---------+---------+
                |         |         |
         +-------------+ +-------------+ +-------------+
         |  Worker 1   | |  Worker 2   | |  Worker N   |
         +-------------+ +-------------+ +-------------+
*/

type Task struct {
	ID       int
	Duration time.Duration
}

func worker(id int, taskChan <-chan Task, wg *sync.WaitGroup) {
	for task := range taskChan {
		fmt.Printf("[Worker %d] Start task %d\\n", id, task.ID)
		time.Sleep(task.Duration)
		fmt.Printf("[Worker %d] End task %d\\n", id, task.ID)
		wg.Done()
	}
}

// Dispatcher с балансировкой нагрузки
type Dispatcher struct {
	workChans []chan Task
	current   int
	mu        sync.Mutex
}

func NewDispatcher(workerCount int) *Dispatcher {
	chans := make([]chan Task, workerCount)
	for i := range chans {
		chans[i] = make(chan Task)
	}
	return &Dispatcher{workChans: chans}
}

func (d *Dispatcher) Start(wg *sync.WaitGroup) {
	for i, ch := range d.workChans {
		go worker(i, ch, wg)
	}
}

func (d *Dispatcher) Dispatch(task Task) {
	d.mu.Lock()
	target := d.current
	d.current = (d.current + 1) % len(d.workChans)
	d.mu.Unlock()

	d.workChans[target] <- task
}

func main() {
	workerCount := 3
	taskCount := 10

	dispatcher := NewDispatcher(workerCount)
	var wg sync.WaitGroup
	dispatcher.Start(&wg)

	for i := 0; i < taskCount; i++ {
		task := Task{
			ID:       i + 1,
			Duration: time.Second,
		}
		wg.Add(1)
		dispatcher.Dispatch(task)
	}

	wg.Wait()

	for _, ch := range dispatcher.workChans {
		close(ch)
	}
}

/*
📊 Результат
Задачи отправляются в воркеры по кругу (round-robin).
Каждый воркер получает равное количество задач.
Все задачи обрабатываются параллельно.

🚀 Возможные улучшения
✅ Учитывать загруженность воркеров (например, очередь задач у воркера)
✅ Использовать context.Context для отмены
✅ Добавить retry или dead letter очередь
✅ Поддерживать добавление воркеров на лету
✅ Использовать message broker (Kafka, NATS, Redis Streams) для распределения между машинами
*/
