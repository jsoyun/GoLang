package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	var root *Node
	var tail *Node
	//노드의 주소를 root로 가지고 있음
	root = &Node{val: 0}
	//맨처음꺼 하나만 있을 때는 tail은 root와 같음
	tail= root

	for i:=1; i<10; i++ {
		tail= AddNode(tail,i)
	}
	
	PrintNodes(root)

	root,tail = RemoveNode(root.next, root , tail)

	PrintNodes(root)

	root,tail = RemoveNode(root, root, tail)
	PrintNodes(root)

	root,tail = RemoveNode(tail, root, tail)
	PrintNodes(root)
	fmt.Printf("tail:%d\n", tail.val)

}

//노드 추가하는 거를 함수로 만들게요 //새로 추가된 노드로 반환값있어야함
func AddNode(tail *Node, val int) *Node {

	// var tail *Node
	// tail = root
	// for tail.next != nil {
	// 	tail = tail.next
	// }
	//맨 마지막에 새 노드 추가
	node := &Node{val: val}
	tail.next = node
	return node

}


func RemoveNode(node *Node,root *Node, tail *Node ) (*Node, *Node) {
	//내가 지우고자 하는 노드가 맨 앞인 경우
	if node == root {
		//새로운 루트는 기존 루트의 다음이 됨
		root = root.next
		if root == nil {
			tail= nil
		}
		return root, tail

	}



	//이전 노드 다음이 현재 지우고자 하는 노드가 아니면
	//이전을 이전다음으로 보내. 맞게 되면 for문 빠져나가 
	prev := root
	for prev.next != node {
		
		prev = prev.next

	}
	if node== tail {
		prev.next = nil
		tail = prev
	} else {
		prev.next= prev.next.next
	}

	return root, tail
}

func PrintNodes(root *Node) {
	//그다음 노드가 없을때까지 전진
	node := root
	for node.next != nil {
		fmt.Printf("%d -> ",node.val)
		node = node.next
	}
	fmt.Printf("%d\n",node.val)
}