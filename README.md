# dhrand

[![Build Status](https://img.shields.io/travis/dhui/dhrand/master.svg)](https://travis-ci.org/dhui/dhrand) [![Code Coverage](https://img.shields.io/codecov/c/github/dhui/dhrand.svg)](https://codecov.io/gh/dhui/dhrand) [![GoDoc](https://godoc.org/github.com/dhui/dhrand?status.svg)](https://godoc.org/github.com/dhui/dhrand) [![Go Report Card](https://goreportcard.com/badge/github.com/dhui/dhrand)](https://goreportcard.com/report/github.com/dhui/dhrand) [![GitHub Release](https://img.shields.io/github/release/dhui/dhrand/all.svg)](https://github.com/dhui/dhrand/releases) ![Supported Go versions](https://img.shields.io/badge/Go-1.12%2C%201.13-lightgrey.svg)

dhrand is a Golang package that exports a `LockedSource` similar to [golang.org/x/exp/rand](https://godoc.org/golang.org/x/exp/rand)

## Benchmarks

```
$ go test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/dhui/dhrand
Benchmark/default1-8         	100000000	        10.1 ns/op	       0 B/op	       0 allocs/op
Benchmark/default4-8         	100000000	        15.7 ns/op	       0 B/op	       0 allocs/op
Benchmark/default8-8         	100000000	        23.2 ns/op	       0 B/op	       0 allocs/op
Benchmark/default16-8        	50000000	        38.5 ns/op	       0 B/op	       0 allocs/op
Benchmark/default32-8        	20000000	        69.0 ns/op	       0 B/op	       0 allocs/op
Benchmark/default64-8        	10000000	       133 ns/op	       0 B/op	       0 allocs/op
Benchmark/default128-8       	 5000000	       259 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand1-8          	100000000	        19.8 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand4-8          	30000000	        55.4 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand8-8          	20000000	       102 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand16-8         	10000000	       191 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand32-8         	 5000000	       354 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand64-8         	 2000000	       694 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand128-8        	 1000000	      1355 ns/op	       0 B/op	       0 allocs/op
BenchmarkModule/1-8          	50000000	        30.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkModule/4-8          	50000000	        32.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkModule/8-8          	50000000	        35.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkModule/16-8         	30000000	        47.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkModule/32-8         	20000000	        70.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkModule/64-8         	10000000	       124 ns/op	       0 B/op	       0 allocs/op
BenchmarkModule/128-8        	10000000	       214 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/dhui/dhrand	36.694s
```
