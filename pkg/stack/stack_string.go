package stack

import "fmt"

const (
	defaultStackStringLen = 100
)

type StackString struct {
	arr []string
}

func NewStackString() *StackString {
	return &StackString{
		arr: make([]string, 0, defaultStackStringLen),
	}
}

func (s *StackString) Push(v string) {
	s.arr = append(s.arr, v)
}

func (s *StackString) Peek() (string, bool) {
	if len(s.arr) == 0 {
		return "", false
	}

	result := s.arr[len(s.arr)-1]
	return result, true
}

func (s *StackString) Pop() (string, bool) {
	if len(s.arr) == 0 {
		return "", false
	}

	result := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return result, true
}

func (s *StackString) Print() {
	fmt.Println(s.arr)
}
