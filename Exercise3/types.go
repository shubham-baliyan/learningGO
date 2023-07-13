package main

import "fmt"

func main() {
	a, b, c := 1, 2.4, "hello"

	fmt.Printf("%v and %T", a, a)
	fmt.Printf("%v and %T", b, b)
	fmt.Printf("%v and %T", c, c)

}
