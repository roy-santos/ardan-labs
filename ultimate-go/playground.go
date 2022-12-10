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

	// Struct types - user defined types
	// example type consists of 8 bytes, one of which is for "alignment"
	// alignments ensure that values fall within 1 8 byte word, 2 byte values need 2 byte alignment
	// 4 byte value needs 4 byte alignment, and so on...
	// can optimize amount of padding by ordering fields by size, DO NOT DO THIS UNLESS NECESSARY
	type example struct { // composite type
		flag    bool
		counter int16
		pi      float32
	}

	// Declare a variable of type example set to its zero value.
	var e1 example // can also use empty literal construction `e1 := example{}`

	// Display the value.
	fmt.Printf("%+v\n", e1)

	// Declare a variable of type example and init using a struct literal
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the field values.
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("pi", e2.pi)

	// Declare a variable of an anonymous type and init to its zero value.
	var e3 struct {
		flag    bool
		counter int16
		pi      float32
	}

	// Display the value
	fmt.Printf("%+v\n", e3)

	// Declare a variable of an anonymous type and init using a struct literal.
	e4 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    false,
		counter: 10,
		pi:      3.141592,
	}

	// Display the field values.
	fmt.Printf("%+v\n", e4)
	fmt.Println("Flag", e4.flag)
	fmt.Println("Counter", e4.counter)
	fmt.Println("pi", e4.pi)

	// Two named types that are pretty much identical except they are named differently,
	// thus they are 2 different types.
	type bill struct {
		flag    bool
		counter int16
		pi      float32
	}

	type nancy struct {
		flag    bool
		counter int16
		pi      float32
	}

	// Declare zero value for each type
	var b1 bill
	var n1 nancy
	// can not assign b1 = n1, implicit conversion between NAMED types is not allowed in go even if they are identical/compatible
	fmt.Println(b1, n1)

	// must explicitly convert compatible types. This is allowed
	b1 = bill(n1)

	// implicit conversion between named and unamed type is allowed
	b1 = e4 // this is OK!

	// Declare variable of type int with a value of 10.
	count := 10

	/*
		Value Semantics - Pass by value, passes copy of the value across program boundaries
	*/

	// Display the "value of" and "address of" count.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "value of" the count.
	incrementByValue(count)

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	/*
		Pointer Semantics - Pass address of the value, passes memory location of the value across program boundaries
	*/

	// Declare variable of type int with a value of 10.
	count2 := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue Of[", count2, "]\tAddr Of[", &count2, "]")

	// Pass the "address of" count.
	incrementByReference(&count2)

	println("count:\tValue Of[", count2, "]\tAddr Of[", &count2, "]")

	/*
		Escape Analysis -  a type of static code analysis done by compiler, reads code at compile time and determines whether
		a value should be constructed on the stack or on the heap (an escape) based on how the value is shared
	*/

	u1 := createUserV1() // copy of value is received from the fxn call
	u2 := createUserV2() // fxn call returns address of the value located on the heap

	// note: values created on the heap have to get managed by garbage collector, anything on the stack is self cleaning
	// there is a tradeoff here. Efficiency vs data integrity

	print("u1", &u1, "u2", u2)

	/*
		Stack Growth -
	*/
}

// increment declares count as a pointer variable whose value is always an address and points to values of type int.
func incrementByValue(inc int) {
	inc++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")

}

// increment declares count as a pointer variable whose value is always an address and points to values of type int.
// this is still "pass by value" where the value being passed is the address.
func incrementByReference(inc *int) { // *int type is address of the int, not the int itself
	// Increment the "value of" count that the "pointer points to".
	// this is indirect access to the original value
	// direct access to a value can only be done if the value is declared within the "frame (in this case, the fxn boundary)"
	*inc++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")

}

// user represents a user in the system.
type user struct {
	name  string
	email string
}

// createUserV1 creates a user value and passes a copy back to the caller
// value semantic function, sends COPY of the user value back
func createUserV1() user {
	userV1 := user{
		name:  "John",
		email: "john@gmail.com",
	}

	println("V1", &userV1)

	return userV1
}

// createUserV2 creates a user value and shares the value with the caller
// pointer semantic function, shares original copy of the user value w/ caller
func createUserV2() *user {
	userV2 := user{ // this value gets created on the heap, represents a value on the heap
		name:  "John",
		email: "john@gmail.com",
	}

	println("V2", &userV2)

	return &userV2 // shares value up call stack, compiler knows to construct value in the heap (escape analysis)
}
