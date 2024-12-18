package graph

type DirectedGraph struct {
	nodes map[int]*Node
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{nodes: make(map[int]*Node)}
}

func (g *DirectedGraph) AddNode(nodeID int) {
	if _, exists := g.nodes[nodeID]; !exists {
		newNode := NewNode(nodeID)
		g.nodes[nodeID] = newNode
	}
}

func (g *DirectedGraph) AddEdge(parentNodeId int, childNodeId int) {
	parent := g.nodes[parentNodeId]
	child := g.nodes[childNodeId]

	parent.Children = append(parent.Children, child)
}
