package boringwoz

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"math/rand/v2"
)

var randomness *rand.Rand

func initRandomness() {
	var a uint64
	var b uint64
	{
		// We store 16 random bytes here.
		//
		// This array is 16 bytes long because a uint64 is 8 bytes long,
		// and we need 2 random uint64. So, we need 8 random bytes for the first uint64,
		// and another 8 random bytes for the second uint64. So:
		//
		//	16 = 2 × 8 = 8 + 8
		var bytes [16]byte

		cryptorand.Read(bytes[:])

		// This seems to be the most performant way of loading bytes into a uint64.
		a = binary.NativeEndian.Uint64(bytes[0:8])
		b = binary.NativeEndian.Uint64(bytes[8:16])
	}

	randomness = rand.New( rand.NewPCG(a, b) )
}

