package gorray

import "strconv"

type Array[T any] struct {
	len   uint
	value []T
}

func NewArray[T any](len uint) *Array[T] {
	return &Array[T]{
		len:   len,
		value: make([]T, 0, len),
	}
}

func (a *Array[T]) Push(value T) {
	if len(a.value) > int(a.len) {
		panic("index out of bounds: the length is " + strconv.Itoa(int(a.len)) + " but the index is " + strconv.Itoa(int(a.len+1)))
	}
	a.value = append(a.value, value)
}
func (a *Array[T]) PushAll(value []T) {
	if len(a.value) != 0 {
		panic("size array is not equal 0")
	}
	for _, e := range value {
		a.Push(e)
	}
}

func (a *Array[T]) Map(cb func(index int, element T) T) []T {
	res := make([]T, 0, len(a.value)+1)
	for i, e := range a.value {
		res = append(res, cb(i, e))
	}
	return res
}

func (a *Array[T]) Iter() []T {
	return a.value
}

func (a *Array[T]) Shift() T {
	if len(a.value) == 0 {
		panic("index out of bounds: the length is " + strconv.Itoa(int(a.len)) + " but the index is " + strconv.Itoa(int(a.len+1)))
	}
	last := a.value[len(a.value)-1]
	a.value = a.value[0 : len(a.value)-1]
	return last
}

func (a *Array[T]) Pop() T {
	if len(a.value) == 0 {
		panic("index out of bounds: the length is " + strconv.Itoa(int(a.len)) + " but the index is " + strconv.Itoa(int(a.len+1)))
	}
	first := a.value[0]
	a.value = a.value[1:len(a.value)]
	return first
}

func (a *Array[T]) GetByIndex(index uint) T {
	if int(index) > len(a.value)-1 {
		panic("index out of bounds: the length is " + strconv.Itoa(int(a.len)) + " but the index is " + strconv.Itoa(int(index)))
	}
	return a.value[index]
}
