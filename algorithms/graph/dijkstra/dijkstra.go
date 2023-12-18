package dijkstra

import (
	"fmt"
	"math"
	"slices"

	"github.com/lspaccatrosi16/go-libs/structures/mpq"
)

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

	queue := mpq.Queue[GraphNode]{}

	for _, node := range graph.Nodes {
		if node.Ident() != start.Ident() {
			dist[node.Ident()] = math.MaxInt
			queue.Add(node, math.MaxInt)
		}
	}

	dist[start.Ident()] = 0
	queue.Add(start, 0)

	for queue.Len() != 0 {
		v := queue.Pop()
		name := v.Ident()
		if visited[name] {
			continue
		}

		neighbors := graph.Edges[v]

		for _, neighbor := range neighbors {
			nName := neighbor.Ident()
			alt := dist[name] + neighbor.Weight()
			if alt < dist[nName] {
				dist[nName] = alt
				prev[nName] = name
			}
		}
	}

	// var ok bool

	order := []GraphNode{}

	fmt.Println(dist[end.Ident()])
	pv := end.Ident()

	for pv != start.Ident() {
		order = append(order, identMap[pv])
		fmt.Println(pv)
		pv = prev[pv]

		// if !ok {
		// 	panic("not found")
		// }
	}

	order = append(order, identMap[start.Ident()])

	slices.Reverse(order)

	fmt.Println("finished")

	return DijkstraRun{
		Visited:      order,
		PathDistance: dist[end.Ident()],
	}
}
