package gameoflife

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLivingNeighbors(t *testing.T) {
	t.Run("should return 0 Living on board with no living cells", func(t *testing.T) {
		var fakeBoard = &Board{
			{Dead, Dead, Dead, Dead},
			{Dead, Dead, Dead, Dead},
			{Dead, Dead, Dead, Dead},
		}

		assert.Equal(t, 0, fakeBoard.livingNeighborsCount(1, 1))
	})

	t.Run("should return all corners as Living on board with only living cells", func(t *testing.T) {
		var fakeBoard = &Board{
			{Living, Living, Living, Living},
			{Living, Living, Living, Living},
			{Living, Living, Living, Living},
		}

		assert.Equal(t, 8, fakeBoard.livingNeighborsCount(1, 1))
	})
}

func TestNextField(t *testing.T) {
	t.Run("when received a field with a single cell", func(t *testing.T) {
		t.Run("it should return an empty field", func(t *testing.T) {
			board := Board{
				{Dead, Dead, Dead, Dead},
				{Dead, Living, Dead, Dead},
				{Dead, Dead, Dead, Dead},
			}

			expected := Board{
				{Dead, Dead, Dead, Dead},
				{Dead, Dead, Dead, Dead},
				{Dead, Dead, Dead, Dead},
			}

			next := board.Next()

			assert.Equal(t, expected, next)
		})
	})
}
