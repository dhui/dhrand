package dhrand

import (
	"bytes"
	"testing"
)

func TestGenSeed(t *testing.T) {
	origCryptoRandReader := cryptoRandReader
	cryptoRandReader = bytes.NewReader(nil)
	t.Run("Error", func(t *testing.T) {
		_, err := GenSeed()
		if err == nil {
			t.Error("Expected an error generating seed")
		}
	})
	cryptoRandReader = origCryptoRandReader
}
