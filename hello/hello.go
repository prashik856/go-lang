// Declare a main package
package main

// import fmt package which contains functions for formatting text
// A standard library package
import "fmt"

// Upgrading our code
import "rsc.io/quote"

// implement a main function to print a message to console
// A main function executes by default when we run the main package
func main() {
	// Printing to console
	fmt.Println("Hello, World!")

	// Print to console using external module
	// Check documentation to see all the available functions
	fmt.Println(quote.Go())
}
