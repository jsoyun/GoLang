package dataStruct

import "fmt"


type Node struct {
	Next *Node
	//더블링크드를 위한 이전prev
	Prev *Node
	Val  int
}
//결합성 높이고 의존성 줄여~
//관련있는 애들 묶어주자. 하나의 struct으로 묶기


//1. 새 struct만듦
type LinkedList struct { 
	//루트를 포인트형으로 가지고 있음
	Root *Node
	Tail *Node
}



//2. 메서드 3개 추가함  add remove printNode
//LinkedList가 가지고 있는기능, 메서드
func (l *LinkedList) AddNode(val int) {
	if l.Root ==nil {
		//없는 상태면
		//루트는, 새로만든 노드의 메모리 주소를 포인트 형태로 갖고 있고
		l.Root = &Node{Val:val}
		//테일은 자료가 없으니까 루트랑 똑같음
		l.Tail= l.Root
		return
	}
	//꼬리 다음을 새로운 노드 만들어서 붙이면 됨
	l.Tail.Next = &Node{Val:val}
	//이전은 기존 테일
	prev := l.Tail
	//새로운 테일로 갱신
	l.Tail= l.Tail.Next
	//새테일의 이전이 prev
	l.Tail.Prev = prev 
}
//맨뒤에 값을 반환하는 함수
func(l *LinkedList) Back() int {
	if l.Tail != nil {
		return l.Tail.Val

	}
	return 0
}

func (l *LinkedList) PopBack(){
	if l.Tail ==nil {
		return
	}
	l.RemoveNode(l.Tail)
}

func  (l *LinkedList)  RemoveNode(node *Node ) {
	if node == l.Root {
		l.Root = l.Root.Next 
		l.Root.Prev = nil
		node.Next = nil
		return
	}

	prev := node.Prev
	if node== l.Tail {
		prev.Next = nil
		l.Tail.Prev = nil
		l.Tail = prev
	} else {
		node.Prev = nil
		//next 다음거에 연결
		prev.Next= prev.Next.Next
		//더블드 리스트라 이전으로도 연결해줘야함 
		prev.Next.Prev = prev
	}
	node.Next= nil
}


func (l *LinkedList) Empty() bool {
	return l.Root == nil
   }


   func (l *LinkedList) PopFront() {
	if l.Root == nil {
	 return
	}
	l.RemoveNode(l.Root)
   }



   func (l *LinkedList) Front() int {
	if l.Root != nil {
	 return l.Root.Val
	}
	return 0
   }



func (l *LinkedList)  PrintNodes() {
	//그다음 노드가 없을때까지 전진
	node := l.Root
	for node.Next != nil {
		fmt.Printf("%d -> ",node.Val)
		node = node.Next
	}
	fmt.Printf("%d\n",node.Val)
}

// 역으로 출력하는 함수
func (l *LinkedList) PrintReverse(){
	node := l.Tail
	for node.Prev != nil {
		fmt.Printf("%d->", node.Val)
		node= node.Prev
	} 
	fmt.Printf("%d\n", node.Val)

}
