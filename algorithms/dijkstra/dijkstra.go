package dijkstra

import (
	"math"
	"slices"

	"github.com/lspaccatrosi16/go-libs/structures/mpq"
)

type GraphNode interface {
	Weight() int
	Ident() string
	RegisterPrevious(p GraphNode)
}

type Graph struct {
	Nodes []GraphNode
	Edges map[GraphNode][]Edge
}

type Edge struct {
	Node   GraphNode
	Weight int
}

type Vertex struct {
	Node     GraphNode
	Distance int
}

func (g *Graph) AddNode(n GraphNode) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(n1, n2 GraphNode, weight int) {
	if g.Edges == nil {
		g.Edges = map[GraphNode][]Edge{}
	}

	e1 := Edge{
		Node:   n2,
		Weight: weight,
	}

	g.Edges[n1] = append(g.Edges[n1], e1)

	e2 := Edge{
		Node:   n1,
		Weight: weight,
	}

	g.Edges[n2] = append(g.Edges[n2], e2)
}

type DijkstraRun struct {
	Visited      []GraphNode
	PathDistance int
}

func RunDijkstra(start, end GraphNode, graph *Graph) DijkstraRun {
	visited := map[string]bool{}
	dist := map[string]int{}
	prev := map[string]string{}

	identMap := map[string]GraphNode{}

	for _, n := range graph.Nodes {
		identMap[n.Ident()] = n
	}

	queue := mpq.Queue[Vertex]{}

	startVertex := Vertex{
		Node:     start,
		Distance: 0,
	}

	for _, node := range graph.Nodes {
		dist[node.Ident()] = math.MaxInt
	}

	dist[start.Ident()] = 0

	queue.Add(startVertex, 0)

	for queue.Len() != 0 {
		v := queue.Pop()
		name := v.Node.Ident()
		if visited[name] {
			continue
		}

		visited[name] = true
		edges := graph.Edges[v.Node]

		for _, edge := range edges {
			eName := edge.Node.Ident()
			if !visited[eName] {
				edge.Node.RegisterPrevious(v.Node)
				provWeight := edge.Node.Weight()
				if provWeight != math.MaxInt && dist[eName]+provWeight < dist[name] {
					newDist := dist[name] + provWeight
					new := Vertex{
						Node:     edge.Node,
						Distance: newDist,
					}
					dist[eName] = newDist
					prev[eName] = name
					queue.Add(new, newDist)
				}
			}
		}
	}

	pv := end.Ident()

	order := []GraphNode{}

	for pv != start.Ident() {
		order = append(order, identMap[pv])
		pv = prev[pv]
	}

	order = append(order, identMap[start.Ident()])

	slices.Reverse(order)

	return DijkstraRun{
		Visited:      order,
		PathDistance: dist[end.Ident()],
	}
}
