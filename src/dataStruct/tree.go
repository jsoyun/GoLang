package dataStruct

type TreeNode struct {
	Val    int
	Childs []*TreeNode
}
type Tree struct {
	Root *TreeNode
}

func (t *Tree) AddNode(cal int) {
	if t.Root == nil {
		t.Root = &TreeNode{Val: val}
	} else {
		t.Root.Childs = append(t.Root.Childs, &TreeNode{Val: val})
	}
}