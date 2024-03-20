package graph

import (
	"fmt"
	"slices"

	"github.com/lspaccatrosi16/go-libs/algorithms/sequences"
	"github.com/lspaccatrosi16/go-libs/structures/graph"
	"github.com/lspaccatrosi16/go-libs/structures/mpq"
)

func RunDijkstra(start, end graph.GraphNode, g *graph.Graph) (graph.GraphRun, error) {
	return runSearch(start, end, g, -1, graph.Dijkstra)
}

func dijkstraLogic(queue *mpq.Queue[graph.GraphNode], visited *map[string]bool, g *graph.Graph, dist *map[string]int, prev *map[string]string, start, end graph.GraphNode, identMap *map[string]graph.GraphNode) (graph.DijkstraRun, error) {
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
			alt := (*dist)[name] + neighbor.Weight()
			if alt < (*dist)[nName] {
				(*dist)[nName] = alt
				(*prev)[nName] = name
				queue.Add(neighbor, alt)
			}
		}
	}

	order := []graph.GraphNode{}
	orderIdents := []string{}

	var ok bool
	pv := end.Ident()

	for pv != start.Ident() {
		order = append(order, (*identMap)[pv])
		orderIdents = append(orderIdents, pv)
		pv, ok = (*prev)[pv]

		cycle := sequences.FindCycle(orderIdents)

		if !ok || cycle.RepeatLength > 1 {
			return graph.DijkstraRun{}, fmt.Errorf("could not find path from start to end")
		}
	}

	order = append(order, (*identMap)[start.Ident()])

	slices.Reverse(order)

	return graph.DijkstraRun{
		Path:     order,
		Distance: (*dist)[end.Ident()],
	}, nil
}
