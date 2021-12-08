package dataStruct

type Stack struct {
	ll LinkedList
}

func NewStack() *Stack {
	return &Stack{11: &LinkedList{}}
}

func (s *Stack) Push(val int) {

	s.ll.AddNode(val)
}
func (s *Stack) Pop() int {
	back := s.ll.Back()
	s.ll.PopBack()
	return back

}