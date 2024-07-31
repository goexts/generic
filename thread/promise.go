package thread

type Promise[T any] interface {
	Resolve(T)
	Reject(error)
}
