package kaspersky

//type WorkerPool struct {
//	tasks []func()
//	workerQty int
//}
//
//func NewWorkerPool(numOfWorkers int) *WorkerPool {
//	return &WorkerPool{
//		tasks: make([]func(), 0),
//		workerQty: numOfWorkers,
//	}
//}
//
//// Submit - добавить таску в воркер пул
//func (wp *WorkerPool) Submit(task func()) {
//
//}
//
//// SubmitWait - добавить таску в воркер пул и дождаться окончания ее выполнения
//func (wp *WorkerPool) SubmitWait(task func()) {
//
//}
//
//// Stop - остановить воркер пул, дождаться выполнения только тех тасок, который выполняются сейчас
//func (wp *WorkerPool) Stop() {
//
//}
//
//// StopWait - остановить воркер пул, дождаться выполнения всех тасок, даже тех, что не начали выполняться
//func (wp *WorkerPool) StopWait() {
//
//}
