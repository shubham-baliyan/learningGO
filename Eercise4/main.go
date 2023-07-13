package main

import "fmt"

func main() {
	a, b, c := 747, 911, 90210

	fmt.Printf("%v\t %d\t %b\t %x\n", a, a, a, a)
	fmt.Printf("%v\t %d\t %b\t %x\n", b, b, b, b)
	fmt.Printf("%v\t %d\t %b\t %x\n	", c, c, c, c)
}
