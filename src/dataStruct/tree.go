package dataStruct

import "fmt"

type TreeNode struct {
	Val    int
	Childs []*TreeNode
}
type Tree struct {
	Root *TreeNode
}

func (t *TreeNode) AddNode(val int) {
	t.Childs = append(t.Childs, &TreeNode{Val: val})
}

func (t *Tree) AddNode(val int) {
	if t.Root == nil {
		t.Root = &TreeNode{Val: val}
	} else {
		t.Root.Childs = append(t.Root.Childs, &TreeNode{Val: val})
	}
}

//DFS 깊이 우선탐색 1. 재귀호출로 만드는 법
//트리의 메소드
func (t *Tree) DFS1(){
	DFS1(t.Root)
}

//얘는 일반함수
//함수가 자기 함수 다시 콜하는, 재귀호출~로  DFS만들어봄
func DFS1(node *TreeNode) {
	fmt.Printf("%d->", node.Val)

	for i := 0; i < len(node.Childs); i++ {
		DFS1(node.Childs[i])
	}

}


//DFS 깊이 우선탐색 2. 스택으로 만드는법
func (t *Tree) DFS2() {
	//트리 노드 갖는 슬라이스 만들기
	s:= []*TreeNode{}
	//트리노드의 루트를 슬라이스 맨끝에다가 추가
	s = append(s, t.Root)
	
	for len(s)>0 {
		var last * TreeNode
		       //last은 맨 마지막가져오고, s 스택은 하나를 줄임,맨마지막거를 없앤 슬라이스
		last, s = s[len(s)-1], s[:len(s)-1]

		fmt.Printf("%d->",last.Val)

		for i := 0; i<len(last.Childs); i++{
			//스택에 자식노드 맨뒤로 집어넣어
			s= append(s,last.Childs[i])

		}
	}

}

	//순서를 거꾸로 해보고 싶다면 스택에서 이렇게
	//3번째
	func (t *Tree)DFS3(){

		s:=[]*TreeNode{}
		s= append(s, t.Root)
		for len(s)>0 {
			var last * TreeNode
			last, s = s[len(s)-1], s[:len(s)-1]
			fmt.Printf("%d->",last.Val)
			//거꾸로~
			for i := len(last.Childs)-1; i>=0; i--{
				s= append(s, last.Childs[i])
			}

		}
	}

