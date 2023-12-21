package graph

import (
	"github.com/lspaccatrosi16/go-libs/structures/graph"
	"github.com/lspaccatrosi16/go-libs/structures/mpq"
)

// Run a BFS on the defined graph. Set `searchDepth` to -1 to search full graph
func RunBfs(start graph.GraphNode, g *graph.Graph, searchDeptch int) (graph.GraphRun, error) {
	return runSearch(start, start, g, searchDeptch, graph.Bfs)

}

func bfsLogic(queue *mpq.Queue[graph.GraphNode], visited *map[string]bool, g *graph.Graph, dist *map[string]int, maxDepth int) {
	for queue.Len() != 0 {
		v := queue.Pop()
		name := v.Ident()

		if (*visited)[name] {
			continue
		}

		(*visited)[name] = true

		neighbors := g.Edges[v]

		for _, neighbor := range neighbors {
			nName := neighbor.Ident()
			alt := (*dist)[name] + 1
			if maxDepth < 0 || alt <= maxDepth {
				queue.Add(neighbor, 1)
			}
			if alt < (*dist)[nName] {
				(*dist)[nName] = alt
			}
		}
	}
}
