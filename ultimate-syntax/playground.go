package main

import "fmt"

// example represents a user defined type with different fields.
// Declared name (example) first then type (struct)
// provides compiler with info regarding amount of memory to allocate and the data itself
type example struct {
	flag    bool
	counter int16
	pi      float32
}

type example2 struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {

	/*
	 Variables
	*/

	// Declare variable that are set to their zero value (when not instantiated).

	var a int     // zero value state is 0, not specifying size of int (32 or 64) allows compiler to pick most efficient size of int for your processor
	var b string  // zero value state is empty string ""
	var c float64 // zero value state is 0
	var d bool    // zero value is state false

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

	// Declare variables and initialize.
	// Using short variable declaration operator to declare and initialize variable.
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	// Specify type and perform a conversion.
	aaa := int32(10) //
	fmt.Printf("aaa := int32(10) %T [%v]\n\n", aaa, aaa)

	// Strings - In go, a 2 word data structure: pointer, number of bites in string
	var s string
	s = "hello world"
	fmt.Printf("%v\n\n", s)

	/*
	 Struct - (see "example" struct above)
	*/

	// Declare a variable of type example set to its zero value.
	var e1 example

	// Display the value
	fmt.Printf("%+v\n", e1) // styles of formatting (%v, %#v, %+v)
	fmt.Printf("%v\n", e1)
	fmt.Printf("%#v\n\n", e1)

	// Declare an EMPTY literal value (values are will not always initialize to zero)
	emptyLiteral := example{} // in this case, values did initialize to zero value
	fmt.Printf("%+v\n\n", emptyLiteral)

	// Declare variable of type example and init using struct literal
	// dont need to include a particular field if you intend for it to set to zero value
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the field values.
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
	fmt.Println()

	/*
	 Conversions
	*/

	// Declare a variable of an anonymous type and init using a struct literal.
	e := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Create a value of type example
	var ex example
	var ex2 example2

	// Assign the value of the unnamed struct type to the named struct type value.
	ex = e // can do this since it was a LITERAL type and memory layouts are the same, implicit conversion happens
	// ** "ex = ex2" ** cannot do this since they are different NAMED types, no implicit conversion

	// explicit conversion example2 to example
	ex = example(ex2)

	var signedInt int
	var unsignedInt uint
	// ** signedInt = unsignedInt ** cannot do this implicit conversion
	signedInt = int(unsignedInt)
	fmt.Printf("%+v\n\n", signedInt)

	// Display the values.
	fmt.Printf("%+v\n", ex) // styles of formatting (%v, %#v, %+v)
	fmt.Printf("%v\n", e)
	fmt.Printf("%#v\n", e.flag)
	fmt.Printf("%#v\n", e.counter)
	fmt.Printf("%#v\n\n", e.pi)

	/*
	 Pointers
	*/

	// 2 data semantics - value semantics (piece of code is copied when its moved around, easier to find if bug is introduced),
	// pointer semantics (one copy of the data, pointer reference to memory location is passed around, more efficient)
	// each semantic has its own pros and cons

	// Sample to show the basic concept of pass by value

	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	fmt.Println("count: \tValue of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "value of" count.
	incrementValue(count)

	fmt.Println("count: \tValue of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "address of" count.
	incrementPointer(&count)

	fmt.Println("count: \tValue of[", count, "]\tAddr Of[", &count, "]")
}

// increment declares count as a pointer variable whose value is always an address and points to alues of type int.
func incrementValue(inc int) {
	inc++
	fmt.Println("inc: \tValue of[", inc, "]\tAddr Of[", &inc, "]")
}

// increment declares count as a pointer variable whose value is always an address and points to alues of type int.
func incrementPointer(inc *int) {
	// Change the value of the data located at the inc address
	*inc = *inc + 1
	fmt.Println("inc: \tValue of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}
