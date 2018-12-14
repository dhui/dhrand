package dhrand_test

import (
	"math/rand"
	"strconv"
	"testing"
)

import (
	"github.com/dhui/dhrand"
)

var benchSizes = []int{1, 4, 8, 16, 32, 64, 128}

func Benchmark(b *testing.B) {
	srcInfos := []struct {
		name string
		src  rand.Source
	}{
		{name: "default", src: rand.NewSource(testSeed)},        // Not safe to run in goroutines
		{name: "dhrand", src: dhrand.NewLockedSource(testSeed)}, // Safe to run in goroutines
	}

	for _, srcInfo := range srcInfos {
		r := rand.New(srcInfo.src)
		for _, s := range benchSizes {
			buf := make([]byte, s)
			b.Run(srcInfo.name+strconv.Itoa(s), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					r.Read(buf)
				}
			})
		}
	}
}

func BenchmarkModule(b *testing.B) {
	for _, s := range benchSizes {
		buf := make([]byte, s)
		b.Run(strconv.Itoa(s), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Safe to run in goroutines
				rand.Read(buf) // nolint:gosec
			}
		})
	}
}
