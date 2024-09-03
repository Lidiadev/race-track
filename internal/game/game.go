package game

import (
	"race-track/internal/grid"
)

type TestCase struct {
	X, Y       int
	Start, End Point
	Obstacles  []grid.Obstacle
}

type Point struct {
	X, Y int
}

type State struct {
	X, Y, VX, VY int
}

func Play(tc TestCase) int {
	g := grid.NewGrid(tc.X, tc.Y, tc.Obstacles)
	queue := []State{{tc.Start.X, tc.Start.Y, 0, 0}}
	visited := make(map[State]int)
	visited[State{tc.Start.X, tc.Start.Y, 0, 0}] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		hops := visited[current]

		if current.X == tc.End.X && current.Y == tc.End.Y {
			return hops
		}

		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				newVx := current.VX + dx
				newVy := current.VY + dy
				if newVx >= -3 && newVx <= 3 && newVy >= -3 && newVy <= 3 {
					newX := current.X + newVx
					newY := current.Y + newVy
					newState := State{newX, newY, newVx, newVy}
					if !g.IsOccupied(newX, newY) {
						if _, seen := visited[newState]; !seen {
							visited[newState] = hops + 1
							queue = append(queue, newState)
						}
					}
				}
			}
		}
	}

	return -1
}
