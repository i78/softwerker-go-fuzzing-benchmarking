package gameoflife

import (
	"github.com/samber/lo"
	"strings"
)

func (board *Board) livingNeighborsCount(cellX int, cellY int) (result int) {
	neighborOffsets := []int{-1, 0, 1}
	for _, offsetY := range neighborOffsets {
		for _, offsetX := range neighborOffsets {
			if offsetX == 0 && offsetY == 0 {
				continue
			}

			neighborX, neighborY := cellX+offsetX, cellY+offsetY

			if neighborX >= 0 &&
				neighborY >= 0 &&
				neighborY < len(*board) &&
				neighborX < len((*board)[neighborY]) &&
				(*board)[neighborY][neighborX] == Living {
				result++
			}
		}
	}
	return
}

func (board *Board) Next() Board {
	return lo.Map[[]CellState, []CellState](*board, func(it []CellState, y int) []CellState {
		return lo.Map[CellState, CellState](it, func(it CellState, x int) (result CellState) {
			numOfNeighbors := board.livingNeighborsCount(x, y)
			switch (*board)[y][x] {
			case Living:
				if numOfNeighbors >= 2 && numOfNeighbors <= 3 {
					result = Living
				} else {
					result = Dead
				}
			case Dead:
				if numOfNeighbors == 3 {
					result = Living
				} else {
					result = Dead
				}
			}
			return
		})
	})
}

func (c *CellState) String() string {
	switch *c {
	case Living:
		return LivingStringRepresentation
	default:
		return DeadStringRepresentation
	}
}

func (board *Board) StringNaive() (result string) {
	for _, row := range *board {
		for _, col := range row {
			result += col.String()
		}
		result += "\n"
	}
	return
}

func (board *Board) StringBuilderWoPreallocate() string {
	var builder strings.Builder

	for _, row := range *board {
		for _, col := range row {
			builder.WriteString(col.String())
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (board *Board) StringPreallocate() string {
	var builder strings.Builder

	if len(*board) > 0 && len((*board)[0]) > 0 {
		builder.Grow(len(*board)*len((*board)[0])*3 + len(*board))
	}

	for _, row := range *board {
		for _, col := range row {
			builder.WriteString(col.String())
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
