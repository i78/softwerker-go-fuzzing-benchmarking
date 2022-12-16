package gameoflife

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func createFakeBoard(lenX int, lenY int) *Board {
	board := make(Board, lenY)
	for i := range board {
		board[i] = make([]CellState, lenX)
	}
	return &board
}

func FuzzLivingNeighbors(f *testing.F) {
	var fakeBoard = createFakeBoard(4, 4)

	testcases := []struct {
		x int
		y int
	}{
		{
			x: 1,
			y: 1,
		}, {
			x: 2,
			y: 2,
		},
	}

	for _, testcase := range testcases {
		f.Add(testcase.x, testcase.y)
	}

	f.Fuzz(func(t *testing.T, x int, y int) {
		result := fakeBoard.livingNeighborsCount(x, y)
		// Unpredictable output but basic assumptions
		assert.NotNil(t, result)
		assert.GreaterOrEqual(t, result, 0)
		assert.LessOrEqual(t, result, 8)
	})
}
