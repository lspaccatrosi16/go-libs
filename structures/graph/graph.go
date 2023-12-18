package graph

type GraphNode interface {
	Weight() int
	Ident() string
}

type Graph struct {
	Nodes []GraphNode
	Edges map[GraphNode][]GraphNode
}

func (g *Graph) AddNode(n GraphNode) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(n1, n2 GraphNode, weight int) {
	if g.Edges == nil {
		g.Edges = map[GraphNode][]GraphNode{}
	}

	g.Edges[n1] = append(g.Edges[n1], n2)
	g.Edges[n2] = append(g.Edges[n2], n1)
}
