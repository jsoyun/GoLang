package main

import "fmt"


type Node struct {
	next *Node
	val  int
}
//결합성 높이고 의존성 줄여~
//관련있는 애들 묶어주자. 하나의 struct으로 묶기


//1. 새 struct만듦
type LinkedList struct { 
	//루트를 포인트형으로 가지고 있음
	root *Node
	tail *Node
}

//2. 메서드 3개 추가함  add remove printNode
//LinkedList가 가지고 있는기능, 메서드
func (l *LinkedList) AddNode(val int) {
	if l.root ==nil {
		//없는 상태면
		//루트는, 새로만든 노드의 메모리 주소를 포인트 형태로 갖고 있고
		l.root = &Node{val:val}
		//테일은 자료가 없으니까 루트랑 똑같음
		l.tail= l.root
		return
	}
	//꼬리 다음을 새로운 노드 만들어서 붙이면 됨
	l.tail.next = &Node{val:val}
	l.tail= l.tail.next
}


func  (l *LinkedList)  RemoveNode(node *Node ) {
	if node == l.root {
		l.root = l.root.next 
		node.next = nil
		return
	}

	prev := l.root
	for prev.next != node {
		
		prev = prev.next

	}
	if node== l.tail {
		prev.next = nil
		l.tail = prev
	} else {
		prev.next= prev.next.next
	}
	node.next = nil
}



func (l *LinkedList)  PrintNodes() {
	//그다음 노드가 없을때까지 전진
	node := l.root
	for node.next != nil {
		fmt.Printf("%d -> ",node.val)
		node = node.next
	}
	fmt.Printf("%d\n",node.val)
}


func main() {
	//이안에 root와 tail 포함되어있어서 따로 만들 필요 없음
	list := &LinkedList{}
	list.AddNode(0)

	// var root *Node
	// var tail *Node
	// //노드의 주소를 root로 가지고 있음
	// root = &Node{val: 0}
	// //맨처음꺼 하나만 있을 때는 tail은 root와 같음
	// tail= root

	for i:=1; i<10; i++ {
		// tail= AddNode(tail,i)
	    list.AddNode(i)
	}
	
	list.PrintNodes()



   list.RemoveNode(list.root.next)

	list.PrintNodes()

    list.RemoveNode(list.root)
	list.PrintNodes()

	list.RemoveNode(list.tail)
	list.PrintNodes()
	fmt.Printf("tail:%d\n", list.tail.val)

}

// //노드 추가하는 거를 함수로 만들게요 //새로 추가된 노드로 반환값있어야함
// func AddNode(tail *Node, val int) *Node {

// 	// var tail *Node
// 	// tail = root
// 	// for tail.next != nil {
// 	// 	tail = tail.next
// 	// }
// 	//맨 마지막에 새 노드 추가
// 	node := &Node{val: val}
// 	tail.next = node
// 	return node

// }


// func RemoveNode(node *Node,root *Node, tail *Node ) (*Node, *Node) {
// 	//내가 지우고자 하는 노드가 맨 앞인 경우
// 	if node == root {
// 		//새로운 루트는 기존 루트의 다음이 됨
// 		root = root.next
// 		if root == nil {
// 			tail= nil
// 		}
// 		return root, tail

// 	}
// 	//이전 노드 다음이 현재 지우고자 하는 노드가 아니면
// 	//이전을 이전다음으로 보내. 맞게 되면 for문 빠져나가 
// 	prev := root
// 	for prev.next != node {
		
// 		prev = prev.next

// 	}
// 	if node== tail {
// 		prev.next = nil
// 		tail = prev
// 	} else {
// 		prev.next= prev.next.next
// 	}

// 	return root, tail
// }

// func PrintNodes(root *Node) {
// 	//그다음 노드가 없을때까지 전진
// 	node := root
// 	for node.next != nil {
// 		fmt.Printf("%d -> ",node.val)
// 		node = node.next
// 	}
// 	fmt.Printf("%d\n",node.val)
// }