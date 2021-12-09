package dataStruct

import "fmt"

//이진 검색 트리

type BinaryTreeNode struct {
	Val int

	left  *BinaryTreeNode
	right *BinaryTreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree(v int) *BinaryTree {
	tree := &BinaryTree{}
	tree.Root = &BinaryTreeNode{Val: v}
	return tree
}

func (n *BinaryTreeNode) AddNode(v int) *BinaryTreeNode {
	//현재값보다 새로들어오는 값이 작은경우
	if n.Val > v {
		if n.left == nil {
			//왼쪽에 넣는다
			n.left = &BinaryTreeNode{Val: v}
			return n.left
		} else {
			return n.left.AddNode(v)
		}

	} else {
		if n.right == nil {
			n.right = &BinaryTreeNode{Val: v}
			return n.right
		} else {
			return n.right.AddNode(v)
		}
	}

}

type depthNode struct {
	depth int
	node  *BinaryTreeNode
}

//출력할건데 bfc로 출력, 큐로~
func (t *BinaryTree) Print() {
	q := []depthNode{}
	q = append(q, depthNode{depth: 0, node: t.Root})
	currentDepth := 0
	for len(q) > 0 {
		var first depthNode
		first, q = q[0], q[1:]

		if first.depth != currentDepth {
			fmt.Println()
			currentDepth= first.depth
		}
		fmt.Print(first.node.Val," ")

		if first.node.left != nil {
			q = append(q, depthNode{depth:currentDepth+1, node: first.node.left})
		}

		if first.node.right != nil {
			q = append(q, depthNode{depth:currentDepth+1,node: first.node.right})
		}

	}

}
                    //찾아야하는 값
func (t *BinaryTree) Search(v int) (bool,int) {
	return t.Root.Search(v,1)


}

func (n *BinaryTreeNode) Search(v int,cnt int) (bool,int) {
	if n.Val == v {
		return true,cnt
	} else if n.Val >v {
		if n.left != nil {
			return n.left.Search(v,cnt+1)
		}
		return false, cnt
	} else {
		if n.right != nil{
			return n.right.Search(v,cnt+1)
		}
		return false,cnt
	}

}