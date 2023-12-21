package graph

type GraphNode interface {
	Weight() int
	Ident() string
	Exists() bool
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

type RunType int

const (
	Bfs RunType = iota
	Dijkstra
)

type GraphRun struct {
	Visited      []GraphNode
	DijkstraData DijkstraRun
	Dist         map[string]int
	IdentMap     map[string]GraphNode
	Type         RunType
}

func (g *GraphRun) DistanceFromStart(n GraphNode) (int, bool) {
	if _, ok := g.IdentMap[n.Ident()]; ok {
		return g.Dist[n.Ident()], true

	} else {
		return 0, false
	}
}

type DijkstraRun struct {
	Path     []GraphNode
	Distance int
}
