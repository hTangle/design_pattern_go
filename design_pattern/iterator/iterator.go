package iterator

type Iterator[T any] interface {
	HasNext() bool
	Next()
	Current() T
}

type ArrayIterator[T any] struct {
	Array []T
	Index int
}

func (a *ArrayIterator[T]) HasNext() bool {
	return len(a.Array)-1 >= a.Index
}

func (a *ArrayIterator[T]) Next() {
	a.Index++
}

func (a *ArrayIterator[T]) Current() T {
	return a.Array[a.Index]
}

type Array[T any] []T

func (a Array[T]) Iterator() Iterator[T] {
	return &ArrayIterator[T]{
		Array: a,
		Index: 0,
	}
}
