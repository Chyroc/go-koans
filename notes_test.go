package go_koans

import (
	"testing"
	"fmt"
)

func TestNote(t *testing.T) {
	/*
	5 = 101b // in binary, but Go doesn't have binary literals
	2 = 010b
	XOR:
	7 = 111b
	*/
	// basic
	assert(5^2 == 7)

	// string
	assert("abc"[0] == uint8('a'))                                                             // their contents are reminiscent of C
	assert(fmt.Sprintf("hello %q", "world") == "hello \"world\"")                              // although it can be done more easily
	assert(fmt.Sprintf("your balance: %d and %0.2f", 3, 4.5589) == "your balance: 3 and 4.56") // "the root of all evil" is actually a misquotation, by the way

}
