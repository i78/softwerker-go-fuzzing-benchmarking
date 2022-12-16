package gameoflife

type Board [][]CellState
type CellState bool

const (
	Dead   = false
	Living = true
)

const (
	DeadStringRepresentation   = " . "
	LivingStringRepresentation = " * "
)
