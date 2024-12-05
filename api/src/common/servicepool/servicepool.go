package servicepool

type ServicePool[T any] interface {
	Get() (service T, free func())
}

//

//

type servicePool[T any] struct{ pool chan T }

func (services *servicePool[T]) Get() (T, func()) {
	s := <-services.pool
	return s, func() { services.pool <- s }
}

func NewPool[T any](fn func() T, size int) ServicePool[T] {
	pool := make(chan T, size)
	for i := 0; i < size; i++ {
		pool <- fn()
	}
	return &servicePool[T]{pool: pool}
}

func PoolFromArray[T any](elements []T) ServicePool[T] {
	pool := make(chan T, len(elements))
	for _, element := range elements {
		pool <- element
	}
	return &servicePool[T]{pool: pool}
}
