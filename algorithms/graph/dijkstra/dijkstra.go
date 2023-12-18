package dijkstra

import (
	"fmt"
	"math"
	"slices"

	"github.com/lspaccatrosi16/go-libs/structures/graph"
	"github.com/lspaccatrosi16/go-libs/structures/mpq"
)

type DijkstraRun struct {
	Path         []graph.GraphNode
	PathDistance int
	NodesVisited []graph.GraphNode
}

func RunDijkstra(start, end graph.GraphNode, g *graph.Graph) DijkstraRun {
	visited := map[string]bool{}
	dist := map[string]int{}
	prev := map[string]string{}

	identMap := map[string]graph.GraphNode{}

	for _, n := range g.Nodes {
		identMap[n.Ident()] = n
	}

	queue := mpq.Queue[graph.GraphNode]{}

	for _, node := range g.Nodes {
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

		neighbors := g.Edges[v]

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

	order := []graph.GraphNode{}

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

	visitedArr := []graph.GraphNode{}

	for k, v := range visited {
		if v {
			visitedArr = append(visitedArr, identMap[k])
		}
	}

	return DijkstraRun{
		Path:         order,
		PathDistance: dist[end.Ident()],
		NodesVisited: visitedArr,
	}
}
