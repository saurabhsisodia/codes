package main

import "fmt"

type Stack struct {
	arr []int
	n   int
}

func (s *Stack) Push(ele int) {
	s.arr = append(s.arr, ele)
	s.n++
}
func (s *Stack) Pop() int {
	out := s.arr[s.n-1]
	s.n--
	s.arr = s.arr[:s.n]
	return out
}
func (s *Stack) Top() int {
	return s.arr[s.n-1]
}
func (s *Stack) Len() int {
	return s.n
}
func NewStack(last int) *Stack {
	s := Stack{arr: []int{last}, n: 1}
	return &s
}
func main() {
	arr := []int{6, 8, 0, 1, 3}
	ans := make([]int, len(arr))
	ans[len(arr)-1] = -1

	s := NewStack(arr[len(arr)-1])
	for i := len(arr) - 2; i >= 0; i-- {

		for s.Len() > 0 && s.Top() <= arr[i] {
			s.Pop()
		}
		if s.Len() == 0 {
			ans[i] = -1
		} else {
			ans[i] = s.Top()
		}
		s.Push(arr[i])

	}
	fmt.Println(ans)
}
