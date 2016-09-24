package background

type Dispatcher struct {
	PoolSize int
	Pool     chan chan Task
}

func NewDispatcher(n int) Dispatcher {
	return Dispatcher{
		PoolSize: n,
		Pool:     make(chan chan Task, n),
	}
}

func (d Dispatcher) Start() {
	for i := 0; i < d.PoolSize; i++ {
		worker := NewWorker(i+1, d.Pool)
		worker.Start()
	}
}

func (d Dispatcher) Dispatch(queue chan Task) {
	go func() {
		for {
			select {
			case work := <-queue:
				go func() {
					worker := <-d.Pool
					worker <- work
				}()
			}
		}
	}()
}
