package grid

type Obstacle struct {
	X1, X2, Y1, Y2 int
}

type Grid struct {
	cells [][]bool
	X, Y  int
}

func NewGrid(x, y int, obstacles []Obstacle) *Grid {
	g := &Grid{
		cells: make([][]bool, y),
		X:     x,
		Y:     y,
	}
	for i := range g.cells {
		g.cells[i] = make([]bool, x)
	}
	for _, obs := range obstacles {
		for x := obs.X1; x <= obs.X2; x++ {
			for y := obs.Y1; y <= obs.Y2; y++ {
				g.cells[y][x] = true
			}
		}
	}
	return g
}

func (g *Grid) IsOccupied(x, y int) bool {
	return x < 0 || x >= g.X || y < 0 || y >= g.Y || g.cells[y][x]
}
