package graph

type Node struct {
	Value       int
	Children    []*Node
	HasChildren bool
}

func NewNode(value int) *Node {
	return &Node{Value: value, Children: make([]*Node, 0), HasChildren: false}
}
