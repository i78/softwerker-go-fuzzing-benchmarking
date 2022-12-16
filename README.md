# Fuzz Testing und Benchmarking in Go

Codebeispiele zum Softwerker-Artikel "Fuzz Testing und Benchmarking in Go".


## Testing
Lediglich die Tests ausführen

```shell
go test -v
```
```
=== RUN   TestLivingNeighbors
=== RUN   TestLivingNeighbors/should_return_0_Living_on_board_with_no_living_cells
=== RUN   TestLivingNeighbors/should_return_all_corners_as_Living_on_board_with_only_living_cells
--- PASS: TestLivingNeighbors (0.00s)
    --- PASS: TestLivingNeighbors/should_return_0_Living_on_board_with_no_living_cells (0.00s)
    --- PASS: TestLivingNeighbors/should_return_all_corners_as_Living_on_board_with_only_living_cells (0.00s)
PASS
ok      codecentric.de/gol-performance-demo/gameoflife  0.168s
```

### Coverage

```shell
go test --cover
```
```
PASS
coverage: 100.0% of statements
ok      codecentric.de/gol-performance-demo/gameoflife  0.294s
```

```shell
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```


## Fuzz Testing

Lediglich die durch den Seed-Korpus definierten Tests wie "normale" Tests ausführen. Dies wird ebenfalls bereits aufgezeichnete Findings erneut ausführen.

```shell
go test . -run=FuzzLivingNeighbors
``` 
```
ok      codecentric.de/gol-performance-demo/gameoflife  0.342s
```

Fuzzer ausführen

```shell
go test -v -fuzz=. -run=FuzzLivingNeighbors
```
```
fuzz: elapsed: 0s, gathering baseline coverage: 0/9 completed
fuzz: elapsed: 0s, gathering baseline coverage: 1/9 completed
--- FAIL: FuzzLivingNeighbors (0.03s)
```

Fuzzer ausführen und nach Ablauf des Zeitlimits beenden
```shell
go test -v -fuzz=. -run=FuzzLivingNeighbors -fuzztime 1m30s  
```

## 3. Benchmarking

Benchmarks für die drei Varianten ausführen

```shell
go test --benchmem --bench=ToString
```
```
BenchmarkToStringWithPreallocate-12                 3202            349857 ns/op          360490 B/op          1 allocs/op
BenchmarkToStringNaive-12                              1        3688076471 ns/op        22130258984 B/op          122029 allocs/op
BenchmarkToStringWithoutPreallocate-12              1809            746986 ns/op         1981265 B/op         29 allocs/op
BenchmarkToStringDifferentSizes/method=naive_board_size=8-12             2347096               516.0 ns/op           504 B/op          6 allocs/op
BenchmarkToStringDifferentSizes/method=without_preallocate_board_size=8-12               2361716               508.6 ns/op           504 B/op          6 allocs/op
BenchmarkToStringDifferentSizes/method=with_preallocate_board_size=8-12                  4726977               257.9 ns/op           208 B/op          1 allocs/op
BenchmarkToStringDifferentSizes/method=naive_board_size=128-12                             14726             81661 ns/op          211705 B/op         21 allocs/op
BenchmarkToStringDifferentSizes/method=without_preallocate_board_size=128-12               14637             82702 ns/op          211704 B/op         21 allocs/op
BenchmarkToStringDifferentSizes/method=with_preallocate_board_size=128-12                  24180             49878 ns/op           57344 B/op          1 allocs/op
BenchmarkToStringDifferentSizes/method=naive_board_size=1024-12                              166           7314546 ns/op        16792355 B/op         38 allocs/op
BenchmarkToStringDifferentSizes/method=without_preallocate_board_size=1024-12                164           7309139 ns/op        16792353 B/op         38 allocs/op
BenchmarkToStringDifferentSizes/method=with_preallocate_board_size=1024-12                   390           3029826 ns/op         3153927 B/op          1 allocs/op
BenchmarkToStringDifferentSizes/method=naive_board_size=2048-12                               46          25400299 ns/op        65469217 B/op         44 allocs/op
BenchmarkToStringDifferentSizes/method=without_preallocate_board_size=2048-12                 43          25196946 ns/op        65469219 B/op         44 allocs/op
BenchmarkToStringDifferentSizes/method=with_preallocate_board_size=2048-12                    96          12124180 ns/op        12591109 B/op          1 allocs/op

```



## Referenzen

### Links
- https://go.dev/doc/diagnostic
- https://go.dev/security/fuzz/
