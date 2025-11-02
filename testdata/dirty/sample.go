package testdata

import "fmt"

// BARK: This is a test function that needs cleanup
func TestFunction() {
	fmt.Println("Hello, World!")
	// BARK: Remove this debug code before commit
	fmt.Println("Debug information")
	// BARK plain marker without colon
	x := 42
}

// Regular comment without BARK
func AnotherFunction() {
	// BARK: TODO fix this hack
	x := 42
	// BARK
	fmt.Println(x)
}
