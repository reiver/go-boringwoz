package boringwoz

import (
	"bytes"
)

var (
	forbidden          = []byte("boring_wozniak")
	nothingIsForbidden = []byte("")
)

// RandomName returns a "boring_wozniak" style human-friendly random-name.
//
// If `numDigits` is less-than or equal-to zero (0), then the returned "boring_wozniak" style **random-name** is of the following form:
//
//	<ADJECTIVE> "_" <SURNAME>
//
// For example:
//
//	* ethereal_mcdonald
//	* epic_maxwell
//	* goofy_black
//	* candid_ahan
//	* stoic_knuth
//
// Else if, `numDigits` is greater-than zero (0), then the returned "boring_wozniak" style human-friendly random-name will have an underscore (_) followed by one or more (base-10) (arabic-numeral) digits after it:
//
//	<ADJECTIVE> "_" <SURNAME> "_" <DIGITS>
//
// Where the number of digits equals the value of `numDigits`.
// So, if `numDigits` is 3, then there will be (exactly) 3 digits after it.
//
// For example:
//
//	* ethereal_mcdonald_002
//	* epic_maxwell_42
//	* goofy_black_1
//	* candid_ahan_2
//	* stoic_knuth_7
//
// Note that although [RandomNameClassic] will NEVER return "boring_wozniak", "boring_wozniak0", "boring_wozniak1", "boring_wozniak2", "boring_wozniak3", "boring_wozniak4", "boring_wozniak5", "boring_wozniak6", "boring_wozniak7", "boring_wozniak8", "boring_wozniak9",
// RandomName could.
func RandomName(numDigits int) string {
	const includeUnderscoreInSuffix bool = true
	return randomName(allAdjectives, allSurnames, nothingIsForbidden, includeUnderscoreInSuffix, numDigits)
}

// RandomNameClassic returns a "boring_wozniak" style human-friendly random-name.
//
// If `retry` is less-than or equal-to zero (0), then the returned "boring_wozniak" style **random-name** is of the following form:
//
//	<ADJECTIVE> "_" <SURNAME>
//
// For example:
//
//	* epic_maxwell
//	* goofy_black
//	* stoic_knuth
//
// Else if, `retry` is greater-than zero (0), then the returned "boring_wozniak" style human-friendly random-name will have a single (base-10) (arabic-numeral) digit after it:
//
//	<ADJECTIVE> "_" <SURNAME> <DIGIT>
//
// For example:
//
//	* epic_maxwell4
//	* goofy_black1
//	* stoic_knuth7
//
// Note that RandomNameClassic will NEVER return "boring_wozniak".
// And, will never return "boring_wozniak0", "boring_wozniak1", "boring_wozniak2", "boring_wozniak3", "boring_wozniak4", "boring_wozniak5", "boring_wozniak6", "boring_wozniak7", "boring_wozniak8", "boring_wozniak9".
// As per the comment in the original implementation of this algorithm: "Steve Wozniak is not boring" 🙂
// In fact, that is where this package ("boringwoz") gets its name from.
func RandomNameClassic(retry int) string {
	const includeUnderscoreInSuffix bool = false
	return randomName(wozAdjectives[:], wozSurnames[:], forbidden, includeUnderscoreInSuffix, retry)
}

func randomName(adjectives []string, surnames []string, forbidden []byte, includeUnderscoreInSuffix bool, numDigits int) string {
	var buffer [128]byte
	var p []byte = buffer[0:0]

	for {
		var (
			adjective = adjectives[randomness.IntN(len(adjectives))]
			surname   = surnames[randomness.IntN(len(surnames))]
		)

		p = appendBoringWozName(p, adjective, surname)

		if !bytes.Equal(forbidden, p) {
			break
		}

		p = buffer[0:0]
	}

	// Only append an underscore if we are appending digits.
	// Don't include append an underscore by itself (as the suffix).
	if includeUnderscoreInSuffix && 0 < numDigits {
		p = append(p, '_')
	}

	// The classic algorithm only appended a single digit.
	// We do the same for the classic algorithm.
	//
	// We use `includeUnderscoreInSuffix` to determine if the classic algorithm is in effect or not.
	if 0 < numDigits && !includeUnderscoreInSuffix {
		numDigits = 1
	}

	for index := 0; index < numDigits; index++ {
		var b byte = byte(randomness.IntN(10))
		b += '0'

		p = append(p, b)
	}

	return string(p)
}

func appendBoringWozName(p []byte, adjective string, surname string) []byte {
	p = append(p, adjective...)
	p = append(p, '_')
	p = append(p, surname...)

	return p
}
