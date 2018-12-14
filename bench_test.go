package dhrand_test

import (
	"math/rand"
	"strconv"
	"testing"
)

import (
	"github.com/dhui/dhrand"
)

func Benchmark(b *testing.B) {
	srcInfos := []struct {
		name string
		src  rand.Source
	}{
		{name: "default", src: rand.NewSource(testSeed)},
		{name: "dhrand", src: dhrand.NewLockedSource(testSeed)},
	}

	sizes := []int{1, 4, 8, 16, 32, 64, 128}

	for _, srcInfo := range srcInfos {
		r := rand.New(srcInfo.src)
		for _, s := range sizes {
			buf := make([]byte, s)
			b.Run(srcInfo.name+strconv.Itoa(s), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					r.Read(buf)
				}
			})
		}
	}
}
