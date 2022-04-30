package main

import "fmt"

type Stack struct {
	arr    []int
	n      int
	minArr []int
}

func (s *Stack) Push(ele int) {
	s.arr = append(s.arr, ele)
	min := ele
	if s.n > 0 {
		x := s.TopOfMin()
		if x < min {
			min = x
		}
	}
	s.minArr = append(s.minArr, min)
	s.n++
}
func (s *Stack) Pop() int {
	ele := s.arr[s.n-1]
	s.n--
	s.arr = s.arr[:s.n]
	s.minArr = s.minArr[:s.n]
	return ele
}
func (s *Stack) GetMin() int {
	return s.TopOfMin()
}
func (s *Stack) TopOfMin() int {
	return s.minArr[s.n-1]
}
func NewStack() *Stack {
	return &Stack{arr: []int{}, n: 0, minArr: []int{}}
}
func main() {
	s := NewStack()
	s.Push(5)
	s.Push(3)
	s.Push(8)
	s.Push(7)
	s.Push(1)
	s.Push(9)
	s.Push(2)
	s.Push(2)
	fmt.Println(s.arr, s.minArr, s.GetMin())
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	fmt.Println(s.arr, s.minArr, s.GetMin())
	s.Pop()
	s.Pop()
	fmt.Println(s.arr, s.minArr, s.GetMin())
}
