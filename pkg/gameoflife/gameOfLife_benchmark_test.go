package gameoflife

import (
	"fmt"
	"testing"
)

func BenchmarkToStringWithPreallocate(b *testing.B) {
	board := createFakeBoard(400, 300)
	for i := 0; i < b.N; i++ {
		board.StringPreallocate()
	}
}

func BenchmarkToStringNaive(b *testing.B) {
	board := createFakeBoard(400, 300)
	for i := 0; i < b.N; i++ {
		board.StringNaive()
	}
}

func BenchmarkToStringWithoutPreallocate(b *testing.B) {
	board := createFakeBoard(400, 300)
	for i := 0; i < b.N; i++ {
		board.StringBuilderWoPreallocate()
	}
}

func BenchmarkToStringDifferentSizes(b *testing.B) {
	for _, s := range []int{8, 128, 1024, 2048} {
		board := createFakeBoard(s, s)
		b.Run(fmt.Sprintf("method=naive board_size=%d", s),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					board.StringBuilderWoPreallocate()
				}
			})

		b.Run(fmt.Sprintf("method=without_preallocate board_size=%d", s),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					board.StringBuilderWoPreallocate()
				}
			})

		b.Run(fmt.Sprintf("method=with_preallocate board_size=%d", s),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					board.StringPreallocate()
				}
			})
	}
}
