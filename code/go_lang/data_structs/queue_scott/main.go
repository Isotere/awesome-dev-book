package main

import (
	"sync/atomic"
	"unsafe"
)

type item struct {
	value int
	next  unsafe.Pointer
}

type Queue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

func NewQueue() Queue {
	dummy := &item{}
	return Queue{
		head: unsafe.Pointer(dummy),
		tail: unsafe.Pointer(dummy),
	}
}

func (q *Queue) Push(value int) {
	// создаем новую Node
	node := &item{value: value}

	for {
		// читаем указатели на tail && next
		tail := atomic.LoadPointer(&q.tail)
		next := atomic.LoadPointer(&(*item)(tail).next)

		// если tail не изменился
		if tail == atomic.LoadPointer(&q.tail) {
			if next == nil {
				// нет нужды корректировать tail, пробуем CAS
				if atomic.CompareAndSwapPointer(&(*item)(tail).next, next, unsafe.Pointer(node)) {
					// CAS успешен, правим tail
					atomic.CompareAndSwapPointer(&q.tail, tail, unsafe.Pointer(node))
					return
				}
			} else {
				// Пробуем скорректировать tail из другой горутины
				atomic.CompareAndSwapPointer(&q.tail, tail, unsafe.Pointer(node))
			}
		}
	}
}

func (q *Queue) Pop() int {
	for {
		head := atomic.LoadPointer(&q.head)
		tail := atomic.LoadPointer(&q.tail)
		next := atomic.LoadPointer(&(*item)(head).next)

		if head == atomic.LoadPointer(&q.head) {
			// Если head & tail одна и та же нода
			if head == tail {
				if next == nil {
					// в очереди только dummy элемент
					return -1
				} else {
					// иначе нужно подкорректировать tail
					atomic.CompareAndSwapPointer(&q.tail, tail, next)
				}
			} else {
				// Пробуем извлечь элемент из очереди
				value := (*item)(next).value
				if atomic.CompareAndSwapPointer(&q.head, head, next) {
					return value
				}
			}
		}
	}
}

func main() {

}
