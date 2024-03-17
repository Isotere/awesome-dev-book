package main

import (
	"sync/atomic"
	"unsafe"
)

type item struct {
	value int
	next  unsafe.Pointer
}

type Stack struct {
	head unsafe.Pointer
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(value int) {
	node := &item{value: value}

	for {
		head := atomic.LoadPointer(&s.head)
		node.next = head

		if atomic.CompareAndSwapPointer(&s.head, head, unsafe.Pointer(node)) {
			return
		}
	}
}

func (s *Stack) Pop() int {
	for {
		head := atomic.LoadPointer(&s.head)
		if head == nil {
			return -1
		}

		next := atomic.LoadPointer(&(*item)(head).next)
		if atomic.CompareAndSwapPointer(&s.head, head, next) {
			return (*item)(head).value
		}
	}
}

func main() {

}
