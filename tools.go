package main

import (
	"math/rand"
)

// Generates a 32-bit CID used in 
func GenCID() []byte{
	// max := uint64(math.Pow(2, 32)) - 1
	return []byte{
		byte(rand.Uint32()),
		byte(rand.Uint32()),
		byte(rand.Uint32()),
		byte(rand.Uint32()),
	}
}