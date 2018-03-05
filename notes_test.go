package go_koans

import (
	"testing"
	"fmt"
)

func TestNote(t *testing.T) {
	// basic
	{
		assert(5^2 == 7)
		// 乘方是：math.Pow()，这个是按位与
	}

	// string
	{
		assert("abc"[0] == uint8('a'))
		assert(fmt.Sprintf("hello %q", "world") == "hello \"world\"")
		assert(fmt.Sprintf("your balance: %d and %0.2f", 3, 4.5589) == "your balance: 3 and 4.56")
		// %q 会加上双引号
		// %0.2f 会四舍五入
	}

	// array alice
	{
		assert(cap([4]string{"apple", "orange", "mango"}) == 4) // it can hold no more
		assert(cap([]string{"apple", "orange", "mango"}) == 3)  // it can hold no more
		// array slice在计算cap的不同

		f := []string{"apple", "orange", "mango"}
		f[0] = "change"
		assert(f[0] == "change")
		// slice会引用计算

		p := []string{"baby", "baby", "lemon"}
		assert(cap(p) == 3)
		p = append(p, "baby!")
		assert(cap(p) == 6)
		// append 会扩展一倍
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
		// switch 在没有选项的时候，默认是true
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
		// struct只要结构和数据一样，就相同，不需要使用同一个type
	}

	// Allocation
	{
		slice := make([]int, 3)
		assert(len(slice) == 3) // make() creates slices of a given length
		// make创建的array的长度是3
	}
}
