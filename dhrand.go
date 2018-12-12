package dhrand

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
	"sync"
)

// LockedSource implements the rand.Source interface but is safe for usage with goroutines.
// Inspired by: golang.org/x/exp/rand
type LockedSource struct {
	sync.Mutex
	src rand.Source
}

// Int63 wraps the original source Int63 call in a mutex
func (ls *LockedSource) Int63() int64 {
	ls.Lock()
	defer ls.Unlock()
	return ls.src.Int63()
}

// Seed wraps the original source Seed call in a mutex
func (ls *LockedSource) Seed(seed int64) {
	ls.Lock()
	defer ls.Unlock()
	ls.src.Seed(seed)
}

// NewLockedSource creates a new LockedSource
func NewLockedSource(seed int64) *LockedSource {
	return &LockedSource{src: rand.NewSource(seed)}
}

// GenSeed generates a seed value using a cryptographically secure random number generator (CSRNG).
// Generated seed values should only be used to generate a truly random seed for a
// pseudo random number generator (PRNG), instead of using the program's initialization time.
// e.g. generated seed values do not make an arbitrary PRNG cryptographically secure.
func GenSeed() (int64, error) {
	seed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return 0, err
	}
	return seed.Int64(), nil
}
