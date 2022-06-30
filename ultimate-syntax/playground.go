package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

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

// user is a struct type that declares user information.
type user struct {
	ID   int
	Name string
}

// updateStats provides update stats.
type updateStats struct {
	Modified int
	Duration float64
	Success  bool
	Message  string
}

const (
	// Max integer value on 64 bit architecture.
	maxInt = 9223372936854775807

	// Much larger value than int64
	bigger = 9223372036854775808543522345

	// Will NOT compile
	// biggerInt int64 = 9223372036854775808543522345
)

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
	fmt.Println()

	/*
	 Literal Struct - a type that doesn't have a name
	*/

	// Declare a variable of anonymous type set to its zero value.
	var e3 struct {
		flag    bool
		counter int16
		pi      float32
	}

	// Display the value.
	fmt.Printf("%+v\n", e3)

	// Declare a variable of an anonymous type and init using a struct literal
	e4 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the values.
	fmt.Printf("%v\n", e4)
	fmt.Println("Flag", e4.flag)
	fmt.Println("Counter", e4.counter)
	fmt.Println("Pi", e4.pi)

	/*
	 Constants
	*/

	// Untyped Constants.
	const ui = 12345    // kind: integer
	const uf = 3.141592 // kind: floating point

	// Typed Constants still use the constant type system but their precision is restricted
	const ti int = 12345        // type: int
	const tf float64 = 3.141592 // type: float64

	// ./constants.go:XX: constant 1000 overflows uint8
	// const mUint8 uint8 = 1000

	// Constant arithmetic supports different kinds.
	// Kind Promotion is used to determine kind in these scenarios.

	// Variable answer will of type float64.
	// var answer = 3 * 0.333 // KindFloat(3) * KindFloat(0.333)

	// Constnat third will be of kind floating point.
	const third = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)

	// Constant zero will be of kind integer.
	const zero = 1 / 3 // KindInt(1) / KindInt(3)

	// This is an example of constant arithmetic between typed and  untyped constants.
	// Must have like types to perform math.
	const one int8 = 1
	const two = 2
	result := one + two

	fmt.Printf("const result \t %T [%v]\n", result, result) // type is same as the typed const.

	// const block so you dont have to keep declaring each with "const"
	// can also work with other types, usually used for imports
	const (
		A1 = iota // 0 : Start at 0
		B1 = iota // 1 : Increment by 1
		C1 = iota // 2 : Increment by 1
	)

	fmt.Println("1:", A1, B1, C1)

	// shorthand use of iota
	const (
		A2 = iota // 0 : Start at 0
		B2        // 1 : Increment by 1
		C2        // 2 : Increment by 1
	)

	fmt.Println("2:", A2, B2, C2)

	// start iota at 1 or any other value, use math
	const (
		A3 = iota + 1 // 0 : Start at 0 + 1
		B3            // 1 : Increment by 1
		C3            // 2 : Increment by 1
	)

	fmt.Println("3:", A3, B3, C3)

	// bitwise opperations when using as "flags", as seen in log package
	const (
		Ldate         = 1 << iota // 1  : Shift 1 to the left 0. 0000 0001
		Ltime                     // 2  : Shift 1 to the left 1. 0000 0010
		Lmicroseconds             // 4  : Shift 1 to the left 2. 0000 0100
		Llongfile                 // 8  : Shift 1 to the left 3. 0000 1000
		Lshortfile                // 16 : Shift 1 to the left 4. 0001 0000
		LUTC                      // 32 : Shift 1 to the left 5. 0010 0000
	)

	fmt.Println("Log:", Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC)

	// Avoid following a code design like below, do not create enum type, this can
	// be tricked by the implicit Kind conversion (you can pass a kind of int instead of an expected
	// duration which will be accepted by the fxn since the fxn accepts a kind of duration)
	/*
		// A duration represents the elapsed time between two instants as an int64
		// nanosecond count. The representation limits the largest representable
		// duration  to approximately 290 years.

		type Duration int64

		// Common durations. There is no definition for units of Day or larger
		// to avoid confusion across daylight savings time zone transitions.
		const (
			Nanosecond  Duration = 1
			Microsecond          = 1000 * Nanosecond
			Millisecond          = 1000 * Microsecond
			Second               = 1000 * Millisecond
			Minute               = 60 * Second
			Hour                 = 60 * Minute
		)

		// do not do: can pass a kind of int here and not a duration, compiler will allow, this just adds extra code
		// that does not protect your application
		func printduration(d duration) {
			fmt.Println(d)
		}

		// do: leave as an int64 instead of duration, you'll have to check the value passed in, this is a tradeoff
		func printduration(d int64) {
			fmt.Println(d)
		}
	*/

	/*
	 Blank Identifier
	*/

	u := user{
		ID:   1,
		Name: "Troyaikmen",
	}

	// update the user name. Don't care about the update stats.
	// blank identifier used. no variable needed since we don't necessarily care about the update stats
	if _, err := updateUser(&u); err != nil {
		fmt.Println(err)
		return
	}

	// Display the update was successful.
	fmt.Println("Updated user record for ID", u.ID)

	/*
	 Function Idioms - see function definitions for notes, basically, Go has some quirks with fxn calls
	*/

	user, err := retrieveUser("sally")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the user profile.
	fmt.Printf("%+v\n", *user)

	/*
	 Literal Functions
	*/

	var n int

	// Declare an anonymous function and call it.
	func() {
		fmt.Println("Direct:", n)
	}()

	// Declare an anonymous function and assign it to a variable.
	f := func() {
		fmt.Println("Variable:", n)
	}

	// Call the anonymous function through the variable
	f()

	// Defer the call to the anonymous function till after main returns.
	// Go has closures which means the value of n, will still be accessible to the defer fxn
	// even though the defer function is called after the main function finishes executing.
	defer func() {
		fmt.Println("Defer 1:", n)
	}()

	// Set the value of n to 3 before the return.
	n = 3

	// Call the anonymous function through the variable.
	f()

	// Defer the call to the anonymous function till after main returns.
	// Defers are called in reverse order of when they are defined. This defer executes before the one above.
	defer func() {
		fmt.Println("Defer 2:", n)
	}()

	/*
	 Short Variable Declaration Operator
	*/

	var err1 error

	// The short variable declaration operator will declare u and redeclare err1.
	user1, err1 := getUser2()
	if err1 != nil {
		return
	}

	fmt.Println(user1)

	// The short variable declartion operator will redeclare user1 and declare err2.
	user1, err2 := getUser2()
	if err2 != nil {
		return
	}

	fmt.Println(user1)

	// NOTE: short variable declaration operators only need to declare 1 new variable, anything else
	// can be a reused variable.

	/*
	 Array Basics
	*/

	// Declare an array of five strings that is initialized
	var fruits [5]string // compiler needs to know length of array at compile time. Cant use variable length. Consts only!
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// NOTE: strings are composed of 2 words, pointer and size. Moving strings around is efficient because you only need to copy the 16 bits
	// that are used to hold pointer and size data (regardless of the size of data the pointer points to)

	// Iterate over the array of strings. "value semantic" form
	for i, fruit := range fruits { // for range loop, helps prevent going out of bounds
		fmt.Println(i, fruit)
	}

	// "pointer semantic" form
	for i := range fruits { // for range loop, helps prevent going out of bounds
		fmt.Println(i, fruits[i])
	}

	// Declare an array of 4 integers that is initialized with some values
	numbers := [4]int{10, 20, 30, 40} // literal construction syntax,

	// go allows you to also specify the index (not that useful)
	// example: numbers1 :=[4]int{0:10, 2:20, 1:30, 3:40}

	// go allows you to use ellipses for size of array when using literal construction syntax
	// example: numbers2 :=[...]int{1, 2, 3, 4, 5, 6}

	// Iterate over the array of numbers
	for i := 0; i < len(numbers); i++ { // traditional forloop syntax
		fmt.Println(i, numbers[i])
	}

	// NOTE: go has no while loops, only for loops
	// example "while" in go: for i<len(numbers) {...code}

	// NOTE: length of an array is part of its type information. Can't assign an array
	// of size 4 to an array of size 5, youll get a type mismatch error
}

// (pointers) increment declares count as a pointer variable whose value is always an address and points to alues of type int.
func incrementValue(inc int) {
	inc++
	fmt.Println("inc: \tValue of[", inc, "]\tAddr Of[", &inc, "]")
}

// (pointers) increment declares count as a pointer variable whose value is always an address and points to alues of type int.
func incrementPointer(inc *int) {
	// Change the value of the data located at the inc address
	*inc = *inc + 1
	fmt.Println("inc: \tValue of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}

// (blank identifier) updateUser updates teh specified user document.
func updateUser(u *user) (*updateStats, error) {

	// response simulates a JSON response
	response := `{"Modified" : 1, "Duration" : 0.005, "Success" : true, "Message" : "Error occurred!"}`

	//Unmarshal the json document into a value of the userStats struct type.
	var us updateStats
	if err := json.Unmarshal([]byte(response), &us); err != nil {
		return nil, err
	}

	// Check the update status to verify the update was successful
	if us.Success != true {
		return nil, errors.New(us.Message)
	}

	return &us, nil
}

// retrieveUser retrieves the user document for the specified user and returns a pointer to a user type value.
func retrieveUser(name string) (*user, error) {

	// Make a call to get the user in a json response.
	// can't use inline scope here because we are using hte response returned from the getUser fxn
	r, err := getUser(name)
	if err != nil {
		return nil, err
	}

	// Unmarshal the json document into a value of the user struct type.
	// err is scoped within the if statement, not outside that
	var u user
	if err := json.Unmarshal([]byte(r), &u); err != nil {
		return nil, err
	}
	return &u, nil
}

// getUser simulates a web call that returns a json document for the specified user.
func getUser(name string) (string, error) {

	response := `{"id" : 1432, "name" : "sally"}`
	return response, nil // using zero value for the return value that we aren't passing (zero value for error type is nil)
}

func getUser2() (*user, error) {
	return &user{1432, "Betty"}, nil
}
