package testdata

import "fmt"

// Regular comment explaining the function
func TestFunction() {
	fmt.Println("Hello, World!")
	// This is a normal comment
	fmt.Println("Debug information")
	x := 42
}

// Another regular comment
func AnotherFunction() {
	// TODO: This is fine - not a bark comment
	x := 42
	fmt.Println(x)
}
