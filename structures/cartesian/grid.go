package cartesian

import (
	"fmt"

	"github.com/lspaccatrosi16/go-libs/algorithms/dijkstra"
)

type CoordinateGrid[T any] map[int]map[int]T

// Add a new coordinate to the grid
func (cg *CoordinateGrid[T]) Add(c Coordinate, val T) {
	if cg == nil {
		cg = new(CoordinateGrid[T])
	}

	if v, ok := (*cg)[c[0]]; ok {
		v[c[1]] = val
	} else {
		(*cg)[c[0]] = map[int]T{c[1]: val}
	}
}

// Swap two points round
func (cg *CoordinateGrid[T]) Swap(c1, c2 Coordinate) {
	if cg == nil {
		return
	}
	(*cg)[c1[0]][c1[1]], (*cg)[c2[0]][c2[1]] = (*cg)[c2[0]][c2[1]], (*cg)[c1[0]][c1[1]]
}

// Get the entry at the coordinate c
func (cg *CoordinateGrid[T]) Get(c Coordinate) T {
	if cg == nil {
		return *new(T)
	}

	return (*cg)[c[0]][c[1]]
}

func (cg *CoordinateGrid[T]) rows() [][]T {
	if cg == nil {
		return nil
	}

	maxX := 0
	maxY := 0

	for x, l := range *cg {
		if x > maxX {
			maxX = x
		}
		for y := range l {
			if y > maxY {
				maxY = y
			}
		}
	}

	lines := make([][]T, maxY+1)

	for i := 0; i < len(lines); i++ {
		lines[i] = make([]T, maxX+1)
	}

	for x, l := range *cg {
		for y, v := range l {
			lines[y][x] = v
		}
	}

	return lines
}

func (cg *CoordinateGrid[T]) cols() [][]T {
	if cg == nil {
		return nil
	}

	lines := [][]T{}

	for _, l := range *cg {
		thisCol := []T{}
		for _, c := range l {
			thisCol = append(thisCol, c)
		}
		lines = append(lines, thisCol)
	}

	return lines
}

// Pretty format the grid
func (cg *CoordinateGrid[T]) String() string {
	lines := cg.rows()

	outStr := ""
	for _, l := range lines {
		s := ""
		for _, v := range l {
			s += fmt.Sprint(v)
		}
		outStr += s + "\n"
	}
	return outStr
}

// Produce a hashable representation of the grid
func (cg *CoordinateGrid[T]) Hash() string {
	arrs := cg.rows()

	hashStr := ""

	for _, l := range arrs {
		for _, v := range l {
			hashStr += fmt.Sprint(v)
		}
	}

	return hashStr
}

// Get the grid's columns
func (cg *CoordinateGrid[T]) GetRows() [][]T {
	return cg.rows()
}

// Get the grid's rows
func (cg *CoordinateGrid[T]) GetCols() [][]T {
	return cg.cols()
}

// Find the shortest path between two elements
func RunDijkstra(cg *CoordinateGrid[int], start, end Coordinate) ([]Coordinate, int) {
	graph := dijkstra.Graph{}

	nm := map[Coordinate]*dijkstraGridPoint{}

	rows := cg.GetRows()

	edges := map[Coordinate][]Coordinate{}

	for y, r := range rows {
		for x, i := range r {
			coord := Coordinate{x, y}
			gp := &dijkstraGridPoint{
				Point: coord,
				W:     i,
			}

			if y+1 < len(rows) {
				edges[coord] = append(edges[coord], coord.Transform(0, 1))
			}

			if x+1 < len(r) {
				edges[coord] = append(edges[coord], coord.Transform(1, 0))
			}

			graph.AddNode(gp)
			nm[coord] = gp
		}
	}

	for c, e := range edges {
		for _, edge := range e {
			graph.AddEdge(nm[c], nm[edge], cg.Get(edge))
		}
	}

	runProf := dijkstra.RunDijkstra(nm[start], nm[end], &graph)

	order := []Coordinate{}

	for _, gn := range runProf.Visited {
		order = append(order, gn.(*dijkstraGridPoint).Point)
	}

	return order, runProf.PathDistance
}

type dijkstraGridPoint struct {
	Point Coordinate
	W     int
}

func (d *dijkstraGridPoint) Ident() string {
	return fmt.Sprintf("%d_%d", d.Point[0], d.Point[1])

}
func (d *dijkstraGridPoint) Weight() int {
	return d.W
}
