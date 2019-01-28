package utils

import (
	"container/list"
)

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	l := list.New()
	return &Stack{l}
}

func (s *Stack) Push(v interface{}) {
	s.list.PushFront(v)
}

func (s *Stack) Pop() interface{} {
	e := s.list.Front()
	if e != nil {
		s.list.Remove(e)
		return e.Value
	}
	return nil
}

func (s *Stack) Len() int {
	return s.list.Len()
}

func (s *Stack) Empty() bool {
	return s.list.Len() == 0
}
