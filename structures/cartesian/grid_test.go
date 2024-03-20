package cartesian

import (
	"fmt"
	"testing"

	"github.com/lspaccatrosi16/go-libs/algorithms/graph"
)

func TestRunes(t *testing.T) {
	input := [][]rune{
		{'a', 'c', 'e', 'd', 'd'},
		{'f', 'm', 'u', 'o', 's'},
		{'e', 'y', 'a', 'o', 'm'},
		{'e', 'c', 'h', 'o', 'p'},
		{'q', 'v', 'r', 'e', 'l'},
	}

	grid := CoordinateGrid[rune]{}

	for y, l := range input {
		for x, c := range l {
			grid.Add(Coordinate{x, y}, c)
		}
	}

	g, nm := grid.CreateGraph(false, []rune{'a', 'c', 'e', 'i', 'o', 'u'}, false)

	start := (*nm)[Coordinate{0, 0}]
	end := (*nm)[Coordinate{3, 4}]

	searchRes, err := graph.RunDijkstra(start, end, g)

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	repr := grid.GraphSearchRepresentation(searchRes)

	fmt.Println(repr)

}
