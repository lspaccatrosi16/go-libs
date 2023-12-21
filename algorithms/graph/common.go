package graph

import (
	"math"

	"github.com/lspaccatrosi16/go-libs/structures/graph"
	"github.com/lspaccatrosi16/go-libs/structures/mpq"
)

func runSearch(start, end graph.GraphNode, g *graph.Graph, maxDepth int, st graph.RunType) graph.GraphRun {
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
		}
	}

	dist[start.Ident()] = 0
	queue.Add(start, 0)

	var dijkstraData graph.DijkstraRun

	switch st {
	case graph.Bfs:
		bfsLogic(&queue, &visited, g, &dist, maxDepth)
	case graph.Dijkstra:
		dijkstraData = dijkstraLogic(&queue, &visited, g, &dist, &prev, start, end, &identMap)
	}

	visitedArr := []graph.GraphNode{}

	for k, v := range visited {
		if v {
			visitedArr = append(visitedArr, identMap[k])
		}
	}

	return graph.GraphRun{
		Visited:      visitedArr,
		DijkstraData: dijkstraData,
		Dist:         dist,
		IdentMap:     identMap,
		Type:         st,
	}
}
