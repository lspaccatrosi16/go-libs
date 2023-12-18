package cartesian

import (
	"fmt"
	"testing"
)

func TestDijkstra(t *testing.T) {
	ipt := [][]int{
		{1, 1, 1},
		{10, 10, 1},
		{10, 10, 1},
	}

	grid := CoordinateGrid[int]{}

	for y, l := range ipt {
		for x, i := range l {
			grid.Add(Coordinate{x, y}, i)
		}
	}

	path, len := RunDijkstraConsec(&grid, Coordinate{0, 0}, Coordinate{2, 2}, 3)
	fmt.Println(path)
	fmt.Println(len)
}
