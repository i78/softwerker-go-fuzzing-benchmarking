package main

import (
	gol "codecentric.de/gol-performance-demo/pkg/gameoflife"
	"fmt"
	"time"
)

func main() {
	var stable = gol.Board{
		{gol.Dead, gol.Dead, gol.Dead, gol.Dead},
		{gol.Dead, gol.Living, gol.Living, gol.Dead},
		{gol.Dead, gol.Living, gol.Living, gol.Dead},
		{gol.Dead, gol.Dead, gol.Dead, gol.Dead},
	}

	for i := 0; i < 16; i++ {
		fmt.Printf("\u001B[H\u001B[2JIteration: %d \n%v", i, stable.StringPreallocate())
		stable = stable.Next()
		time.Sleep(100 * time.Millisecond)
	}
}
