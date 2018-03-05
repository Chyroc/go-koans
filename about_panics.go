package go_koans

import "fmt"

func divideFourBy(i int) int {
	return 4 / i
}

const __divisor__ = 0

func aboutPanics() {
	//assert(__delete_me__) // panics are exceptional errors at runtime
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("%s\n", e)
		}
	}()

	n := divideFourBy(__divisor__)
	assert(n == 2) // panics are exceptional errors at runtime
}
