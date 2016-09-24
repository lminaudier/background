package background

type Task interface {
	Perform() error
}

type TaskFunc func() error

func (fn TaskFunc) Perform() error {
	return fn()
}
