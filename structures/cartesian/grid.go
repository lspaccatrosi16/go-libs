package cartesian

import (
	"fmt"

	"github.com/lspaccatrosi16/go-libs/structures/graph"
	"github.com/lspaccatrosi16/go-libs/structures/mpq"
)

type CoordinateGrid[T comparable] map[int]map[int]T

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

	maxX, maxY := cg.MaxBounds()

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

// Get the maxium values for each of X, Y
func (cg *CoordinateGrid[T]) MaxBounds() (int, int) {
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

	return maxX, maxY
}

func (cg *CoordinateGrid[T]) FloodFill(start Coordinate, border T, fill T) []Coordinate {
	queue := mpq.Queue[Coordinate]{}
	queue.Add(start, 1)

	visited := map[Coordinate]bool{}

	for queue.Len() != 0 {
		cur := queue.Pop()
		curLoc := cg.Get(cur)
		if curLoc != border && !visited[cur] {
			visited[cur] = true
			directions := []Coordinate{
				cur.TransformInDirection(North),
				cur.TransformInDirection(East),
				cur.TransformInDirection(South),
				cur.TransformInDirection(West),
			}
			for _, d := range directions {
				queue.Add(d, 1)
			}
		}
	}

	visitedArr := []Coordinate{}

	for k, v := range visited {
		if v {
			visitedArr = append(visitedArr, k)
			cg.Add(k, fill)
		}
	}

	return visitedArr
}

func (cg *CoordinateGrid[T]) CreateGraph(intWeights bool) (*graph.Graph, *map[Coordinate]*GraphGridPoint) {
	graph := graph.Graph{}

	nm := map[Coordinate]*GraphGridPoint{}
	edges := map[Coordinate][]Coordinate{}

	rows := cg.GetRows()

	for y, r := range rows {
		for x, i := range r {
			coord := Coordinate{x, y}
			var gp *GraphGridPoint

			if v, ok := any(i).(int); ok && intWeights {
				gp = &GraphGridPoint{
					Point: coord,
					W:     v,
				}
			} else {
				gp = &GraphGridPoint{
					Point: coord,
					W:     1,
				}
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
			graph.AddEdge(nm[c], nm[edge], nm[edge].W)
		}
	}

	return &graph, &nm
}

// Generate a secondary grid showing the results of a graph search
func (cg *CoordinateGrid[T]) GraphSearchRepresentation(run graph.GraphRun) *CoordinateGrid[string] {
	repreGrid := CoordinateGrid[string]{}

	mx, my := cg.MaxBounds()

	for x := 0; x <= mx; x++ {
		for y := 0; y <= my; y++ {
			repreGrid.Add(Coordinate{x, y}, " ")
		}
	}

	var visited *GridPointList

	switch run.Type {
	case graph.Bfs:
		visited = new(GridPointList).FromGraphNodes(run.Visited)
	case graph.Dijkstra:
		visited = new(GridPointList).FromGraphNodes(run.DijkstraData.Path)
	}

	for _, dgp := range *visited {
		repreGrid.Add(dgp.Point, "#")
	}

	return &repreGrid
}

type GraphGridPoint struct {
	Point Coordinate
	W     int
}

func (d *GraphGridPoint) Ident() string {
	return fmt.Sprintf("%d_%d", d.Point[0], d.Point[1])

}

func (d *GraphGridPoint) Weight() int {
	return d.W
}

type GridPointList []*GraphGridPoint

func (l *GridPointList) FromGraphNodes(run []graph.GraphNode) *GridPointList {
	l = new(GridPointList)
	for _, r := range run {
		ggp := r.(*GraphGridPoint)
		*l = append(*l, ggp)
	}
	return l
}
