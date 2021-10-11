package main

import "fmt"

func main() {
	const beef, c1 = "eat", 2

	// 枚举
	const (
		Monday, Tuesday = 1, 2
	)

	const (
		Monday1  = 1
		Tuesday2 = 2
	)

	const a1 = iota // 0

	const (
		a = iota //0
		b = iota //1
		c        //2
		d = iota //3
	)

	fmt.Println(beef)
	fmt.Println(c1)

	fmt.Println(a1)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println()

}
