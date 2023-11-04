package main

import "fmt"

func main() {
	a := 2
	// The [1] in %[1]v specifies that it should refer to the first argument in the argument list, which is a.
	fmt.Printf("type: %T value:%[1]v\n", a)
}
