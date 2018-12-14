# dhrand

[![Build Status](https://img.shields.io/travis/dhui/dhrand/master.svg)](https://travis-ci.org/dhui/dhrand) [![Code Coverage](https://img.shields.io/codecov/c/github/dhui/dhrand.svg)](https://codecov.io/gh/dhui/dhrand) [![GoDoc](https://godoc.org/github.com/dhui/dhrand?status.svg)](https://godoc.org/github.com/dhui/dhrand) [![Go Report Card](https://goreportcard.com/badge/github.com/dhui/dhrand)](https://goreportcard.com/report/github.com/dhui/dhrand) [![GitHub Release](https://img.shields.io/github/release/dhui/dhrand/all.svg)](https://github.com/dhui/dhrand/releases) ![Supported Go versions](https://img.shields.io/badge/Go-1.11-lightgrey.svg)

dhrand is a Golang package that exports a `LockedSource` similar to [golang.org/x/exp/rand](https://godoc.org/golang.org/x/exp/rand)

## Benchmarks

```
$ go test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/dhui/dhrand
Benchmark/default1-8         	100000000	        10.5 ns/op	       0 B/op	       0 allocs/op
Benchmark/default4-8         	100000000	        16.0 ns/op	       0 B/op	       0 allocs/op
Benchmark/default8-8         	100000000	        22.9 ns/op	       0 B/op	       0 allocs/op
Benchmark/default16-8        	30000000	        40.3 ns/op	       0 B/op	       0 allocs/op
Benchmark/default32-8        	20000000	        67.8 ns/op	       0 B/op	       0 allocs/op
Benchmark/default64-8        	10000000	       133 ns/op	       0 B/op	       0 allocs/op
Benchmark/default128-8       	 5000000	       252 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand1-8          	100000000	        19.7 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand4-8          	30000000	        57.5 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand8-8          	20000000	       103 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand16-8         	10000000	       194 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand32-8         	 5000000	       355 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand64-8         	 2000000	       683 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand128-8        	 1000000	      1343 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/dhui/dhrand	24.354s
```
