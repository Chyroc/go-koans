package go_koans

import (
	"testing"
	"fmt"
)

func TestNote(t *testing.T) {
	// basic
	assert(5^2 == 7)

	// string
	assert("abc"[0] == uint8('a'))
	assert(fmt.Sprintf("hello %q", "world") == "hello \"world\"")
	assert(fmt.Sprintf("your balance: %d and %0.2f", 3, 4.5589) == "your balance: 3 and 4.56")

	// array
	assert(cap([4]string{"apple", "orange", "mango"}) == 4) // it can hold no more
	assert(cap([]string{"apple", "orange", "mango"}) == 3)  // it can hold no more
	{
		f := []string{"apple", "orange", "mango"}
		f[0] = "change"
		assert(f[0] == "change")
	}

	// slice
	{
		p := []string{"baby", "baby", "lemon"}
		assert(cap(p) == 3)
		p = append(p, "baby!")
		assert(cap(p) == 6)
	}

	// ControlFlow
	{
		var str string
		switch {
		case false:
			str = "first"
		case true:
			str = "second"
		}
		assert(str == "second")
	}

	// struct
	{
		var bob struct {
			name string
			age  int
		}
		bob.name = "bob"
		bob.age = 30

		type person struct {
			name string
			age  int
		}
		var john = person{
			"bob",
			30,
		}
		assert(bob == john) // assuredly, bob is certainly not john.. yet
	}

	// Allocation
	{
		slice := make([]int, 3)
		assert(len(slice) == 3) // make() creates slices of a given length
	}
}
