package main

import (
	g "codecentric.de/gol-performance-demo/pkg/gameoflife"
	"fmt"
	"time"
)

func main() {
	var blinker = g.Board{
		{g.Dead, g.Living, g.Dead, g.Dead},
		{g.Dead, g.Living, g.Dead, g.Dead},
		{g.Dead, g.Living, g.Dead, g.Dead},
		{g.Dead, g.Dead, g.Dead, g.Dead},
	}

	for i := 0; i < 16; i++ {
		fmt.Printf("\u001B[H\u001B[2JIteration: %d \n%v", i, blinker.StringPreallocate())
		blinker = blinker.Next()
		time.Sleep(100 * time.Millisecond)
	}
}
