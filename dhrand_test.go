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
		261049867304784443,
		605394647632969758,
		894385949183117216,
		1443635317331776148,
		1874068156324778273,
		1905388747193831650,
		1976235410884491574,
		2015796113853353331,
		2610529275472644968,
		2703387474910584091,
		2703501726821866378,
		2740103009342231109,
		2775422040480279449,
		2933568871211445515,
		3328451335138149956,
		3510942875414458836,
		3916589616287113937,
		4037200794235010051,
		4324745483838182873,
		4751997750760398084,
		4831389563158288344,
		5263531936693774911,
		5577006791947779410,
		6129484611666145821,
		6263450610539110790,
		6334824724549167320,
		6426100070888298971,
		6941261091797652072,
		7504504064263669287,
		7955079406183515637,
		7981306761429961588,
		8674665223082153551,
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
