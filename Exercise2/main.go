package main

import "fmt"

func main() {

	//short hand
	a := 1
	//var with specificity
	var b float64 = 3.4
	//zero value
	var c int
	var g bool
	var h string

	// multiple initialization

	d, e, f := 1, "@", false

	const (
		_ = iota
		i
		j
		k
	)
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k)

}
