package concurrent

import "context"

<<<<<<< HEAD
type Executor interface {
	Go(handler func(ctx context.Context))
}
=======
// Executor replace go keyword to start a new goroutine
// the goroutine should cancel itself if the context passed in has been cancelled
// the goroutine started by the executor, is owned by the executor
// we can cancel all executors owned by the executor just by stop the executor itself
// however Executor interface does not Stop method, the one starting and owning executor
// should use the concrete type of executor, instead of this interface.
type Executor interface {
	// Go starts a new goroutine controlled by the context
	Go(handler func(ctx context.Context))
}
>>>>>>> 9362ae084505e4d2b7e6c8fa897cf6dfdb8d64f7
