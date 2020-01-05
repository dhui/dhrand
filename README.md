# dhrand

[![Build Status](https://img.shields.io/travis/dhui/dhrand/master.svg)](https://travis-ci.org/dhui/dhrand) [![Code Coverage](https://img.shields.io/codecov/c/github/dhui/dhrand.svg)](https://codecov.io/gh/dhui/dhrand) [![GoDoc](https://godoc.org/github.com/dhui/dhrand?status.svg)](https://godoc.org/github.com/dhui/dhrand) [![Go Report Card](https://goreportcard.com/badge/github.com/dhui/dhrand)](https://goreportcard.com/report/github.com/dhui/dhrand) [![GitHub Release](https://img.shields.io/github/release/dhui/dhrand/all.svg)](https://github.com/dhui/dhrand/releases) ![Supported Go versions](https://img.shields.io/badge/Go-1.12%2C%201.13-lightgrey.svg)

dhrand is a Golang package that exports a `LockedSource` similar to [golang.org/x/exp/rand](https://godoc.org/golang.org/x/exp/rand)

## Benchmarks

2020/01/05 - Go 1.13
```
$ go test -v -bench . -benchmem -run Bench
goos: darwin
goarch: amd64
pkg: github.com/dhui/dhrand
Benchmark/default1-8         	123336901	         9.74 ns/op	       0 B/op	       0 allocs/op
Benchmark/default4-8         	79965352	        15.7 ns/op	       0 B/op	       0 allocs/op
Benchmark/default8-8         	54513741	        22.1 ns/op	       0 B/op	       0 allocs/op
Benchmark/default16-8        	32836630	        37.2 ns/op	       0 B/op	       0 allocs/op
Benchmark/default32-8        	17189065	        67.1 ns/op	       0 B/op	       0 allocs/op
Benchmark/default64-8        	 9563968	       122 ns/op	       0 B/op	       0 allocs/op
Benchmark/default128-8       	 5211151	       234 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand1-8          	60485150	        18.0 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand4-8          	23568386	        47.2 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand8-8          	14561524	        83.0 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand16-8         	 8338760	       141 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand32-8         	 4525593	       277 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand64-8         	 2379886	       524 ns/op	       0 B/op	       0 allocs/op
Benchmark/dhrand128-8        	 1212876	       976 ns/op	       0 B/op	       0 allocs/op
BenchmarkModule/1-8          	49203238	        24.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkModule/4-8          	45214652	        25.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkModule/8-8          	39749751	        30.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkModule/16-8         	28504729	        43.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkModule/32-8         	18041762	        67.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkModule/64-8         	 9782745	       124 ns/op	       0 B/op	       0 allocs/op
BenchmarkModule/128-8        	 5699034	       212 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/dhui/dhrand	29.481s
```
