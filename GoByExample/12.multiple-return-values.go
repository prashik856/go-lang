package main

import "fmt"

// Go has built-in support for multiple return values. This feature is used often in idiomatic Go, for example to return both result and error values from a function.

// The (int, int) in this function signature shows that the function returns 2 ints.
func vals() (int, int) {
	return 3, 7
}

func main() {

	// Here we use the 2 different return values from the call with multiple assignment.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// If you only want a subset of the returned values, use the blank identifier _.
	_, c := vals()
	fmt.Println(c)
	// Accepting a variable number of arguments is another nice feature of Go functions; we’ll look at this next.
}

// go run multiple-return-values.go
// runs on g1.22
