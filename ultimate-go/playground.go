package main

import (
	"fmt"
)

func main() {

	// word - generic allocation that represents an integer or an address

	// 4 built-in types in go
	// not initialized variables get initialized with zero value
	var a int    //not specifying the precision allows compiler to choose most efficient one for the system architecture
	var b string //uses 2 words that consists of a pointer and integer that represents number of bytes for the string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b int \t %T [%v]\n", b, b)
	fmt.Printf("var c int \t %T [%v]\n", c, c)
	fmt.Printf("var d int \t %T [%v]\n", d, d)

	// short variable declaration operator declares & initializes
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := 10 \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 10 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := 10 \t %T [%v]\n", dd, dd)

	// Go has CONVERSION (allocates new bytes needed to create the casted type) over CASTING (over lays a different type structure) to avoid casting bugs
	// Conversion over casting because INTEGRITY matters more (accurate, consistent, efficient)
	// Casting still possible within the "unsafe" package
	aaa := int32(10)

	fmt.Printf("aaa := int32(10) \t %T [%v]\n", aaa, aaa)

}
