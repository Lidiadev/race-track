package game

import (
	"race-track/internal/grid"
	"testing"
)

func TestGame(t *testing.T) {
	testCases := []struct {
		name     string
		testCase TestCase
		expected int
	}{
		{
			name: "Optimal solution",
			testCase: TestCase{
				X: 5, Y: 5,
				Start: Point{X: 4, Y: 0},
				End:   Point{X: 4, Y: 4},
				Obstacles: []grid.Obstacle{
					{X1: 1, X2: 4, Y1: 2, Y2: 3},
				},
			},
			expected: 7,
		},
		{
			name: "No solution",
			testCase: TestCase{
				X: 3, Y: 3,
				Start: Point{X: 0, Y: 0},
				End:   Point{X: 2, Y: 2},
				Obstacles: []grid.Obstacle{
					{X1: 1, X2: 1, Y1: 0, Y2: 2},
					{X1: 0, X2: 2, Y1: 1, Y2: 1},
				},
			},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Play(tc.testCase)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}
