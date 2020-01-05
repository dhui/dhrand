package dhrand_test

import (
	"math/rand"
	"sort"
	"sync"
	"testing"
)

import (
	"github.com/dhui/dhrand"
	"github.com/stretchr/testify/assert"
)

func TestGenSeed(t *testing.T) {
	if _, err := dhrand.GenSeed(); err != nil {
		t.Error("Failed to generate seed:", err)
	}
}

const testSeed = 1

func TestLockedSourceInt63(t *testing.T) {
	ls := dhrand.NewLockedSource(testSeed)
	expected := []int64{
		-9116395095176460495,
		-8643109948232620495,
		-8135501939754691733,
		-7820069852243076544,
		-7396325266844067250,
		-6878602122688291767,
		-6012567541828693485,
		-5401922988856412949,
		-4976694246060841440,
		-4396220517320143611,
		-4139081590241127977,
		-2828042836786027649,
		-2767187248150682804,
		-2688628175992252777,
		-1502087948929427476,
		-1472860532316763732,
		-618038397637989719,
		-476427912178062506,
		406519965772129914,
		867649948573917593,
		1439618508153543461,
		3189049640734126895,
		3252925369347426176,
		3479915565668951127,
		3812638179211989351,
		4286067568638998091,
		5176744881242059097,
		5798067479865859744,
		5956029175780683524,
		6829542927065255669,
		7795536541508947529,
		8842206116831026702,
	}
	actual := make([]int64, len(expected))

	var wg sync.WaitGroup
	for i := 0; i < len(expected); i++ {
		wg.Add(1)
		go func(i int) {
			actual[i] = ls.Int63()
			wg.Done()
		}(i)
	}
	wg.Wait()

	// Need to sort the slice since the order in which the goroutines run is not deterministic
	sort.Slice(actual, func(i, j int) bool { return actual[i] < actual[j] })
	assert.Equal(t, expected, actual, "Expect to be able to use LockedSource in goroutines")
}

func TestLockedSourceSeed(t *testing.T) {
	ls := dhrand.NewLockedSource(testSeed)
	numGoroutines := 32
	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			ls.Seed(int64(i))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestLockedSourceSatisfiesSourceInterface(t *testing.T) {
	rand.New(dhrand.NewLockedSource(testSeed))
}
