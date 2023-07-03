package main

import "fmt"

func main() {
	// variable h of type int DECLARATION
	//is assigned a value of 5 INITIALISATION

	var h int = 5
	// short hand for this is
	a := 5

	// multiple assignment
	c, d, _, f := 0, 7, 2, "shubham"

	fmt.Println(h, a, c, d, f)

	// printing values as binary and hexadecimal
	fmt.Printf("%b and %x", c, c)
	fmt.Printf("%b and %x", d, d)

	// CONVERSIONS
	z := 4.2 // float64
	var m float32 = 3.2

	// we cant assigne z=m different types
	z = float64(m)
	// now its is called coversion
	fmt.Printf("%v and %T ", z, z)
	// short hand declarations can only be declared at function level or inside the function

}
