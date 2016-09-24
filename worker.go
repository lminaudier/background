package background

func NewWorker(id int, pool chan chan Task) Worker {
	return Worker{
		ID:       id,
		Queue:    make(chan Task),
		Pool:     pool,
		QuitChan: make(chan bool),
	}
}

type Worker struct {
	ID       int
	Queue    chan Task
	Pool     chan chan Task
	QuitChan chan bool
}

func (w *Worker) Start() {
	go func() {
		for {
			w.ready()
			w.work()
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

// Wait for work to arrive and performs it
func (w *Worker) work() {
	select {
	case task := <-w.Queue:
		task.Perform()
	case <-w.QuitChan:
		return
	}

}

// Informs the dispatcher that the worker is ready to access work
func (w *Worker) ready() {
	w.Pool <- w.Queue
}
