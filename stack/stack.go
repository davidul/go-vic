package stack

import "github.com/davidul/go-vic/linkedlist"

type Stack[T comparable] struct {
	l linkedlist.LinkedList[T]
}

func NewStack[T comparable]() *Stack[T] {
	lst := linkedlist.LinkedList[T]{}
	return &Stack[T]{l: lst}
}

func (s *Stack[T]) Push(t T) {
	s.l.Add(t)
}

func (s *Stack[T]) Pop() T {
	return s.l.RemoveLast()
}

func (s *Stack[T]) Peek() T {
	return s.l.PeekLast()
}
