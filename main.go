package main

import (
	"bufio"
	"fmt"
	"os"
	"race-track/internal/game"
	"race-track/internal/grid"
	"strconv"
	"strings"
)

/*
2
5 5
4 0 4 4
1
1 4 2 3
3 3
0 0 2 2
2
1 1 0 2
0 2 1 1
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	testCases, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < testCases; i++ {
		tc := parseTestCase(scanner)
		result := game.Play(tc)
		if result == -1 {
			fmt.Print("\nNo solution.")
		} else {
			fmt.Printf("\nOptimal solution takes %d hops.", result)
		}
	}
}

func parseTestCase(scanner *bufio.Scanner) game.TestCase {
	scanner.Scan()
	dims := strings.Fields(scanner.Text())
	x := atoi(dims[0])
	y := atoi(dims[1])

	scanner.Scan()
	points := strings.Fields(scanner.Text())
	start := game.Point{X: atoi(points[0]), Y: atoi(points[1])}
	end := game.Point{X: atoi(points[2]), Y: atoi(points[3])}

	scanner.Scan()
	numObstacles := atoi(scanner.Text())
	obstacles := make([]grid.Obstacle, numObstacles)
	for i := 0; i < numObstacles; i++ {
		scanner.Scan()
		obs := strings.Fields(scanner.Text())
		x1, x2, y1, y2 := atoi(obs[0]), atoi(obs[1]), atoi(obs[2]), atoi(obs[3])
		obstacles[i] = grid.Obstacle{X1: x1, X2: x2, Y1: y1, Y2: y2}
	}

	return game.TestCase{Start: start, End: end, X: x, Y: y, Obstacles: obstacles}
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
