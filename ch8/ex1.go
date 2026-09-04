package main

type doubleable interface {
	~int | ~int8 | ~int32 | ~int64 |
		~float32 | ~float64
}

func Doubler[T doubleable](n T) T {
	return n * 2
}
