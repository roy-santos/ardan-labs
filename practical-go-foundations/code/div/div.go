package main

import (
	"fmt"
	"log"
)

func main() {
	// fmt.Println(div(1, 0)) // divide by zero error
	fmt.Println(safeDiv(1, 0))
	fmt.Println(safeDiv(7, 2))

}

// named return values
func safeDiv(a, b int) (q int, err error) {
	// q & err are local variables in safeDiv
	// (just like a & b)
	defer func() {
		// e's type is any (or interface{}) *not* error
		// this way of error handling might not be as good as allowing code to panic and return stack trace
		if e := recover(); e != nil {
			log.Println("ERROR:", e)
			err = fmt.Errorf("%v", e)
		}
	}()

	// panic("ouch!")

	/* another way to return the values
	q = a / b
	return
	*/

	return a / b, nil
}

func div(a, b int) int {
	return a / b
}
