package tools

import (
	"math/rand"
)

// Generates a 128-bit CID
func GenCID() []byte{
	// max := uint64(math.Pow(2, 32)) - 1
	return []byte{
		byte(rand.Uint32()),
		byte(rand.Uint32()),
		byte(rand.Uint32()),
		byte(rand.Uint32()),
	}
}