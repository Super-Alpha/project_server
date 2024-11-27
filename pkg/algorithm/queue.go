package algorithm

type Stack struct {
	Queue []int
}

func (s *Stack) Push(x int) {
	s.Queue = append(s.Queue, x)
}

func (s *Stack) Pop() {
	s.Queue = s.Queue[1:len(s.Queue)]
}

func (s *Stack) Top() int {
	return s.Queue[len(s.Queue)-1]
}

func (s *Stack) Empty() bool {
	return len(s.Queue) == 0
}
