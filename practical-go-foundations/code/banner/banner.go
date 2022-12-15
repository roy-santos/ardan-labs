package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
	banner("G☺", 6)

	s := "G☺"
	fmt.Println("len:", len(s))
	// rune = code points ~= unicode character

	for i, r := range s {
		fmt.Println(i, r)
		if i == 0 {
			fmt.Printf("%c of type %T\n", r, r)
			// rune (int32)
		}
	}

	b := s[0]
	fmt.Printf("%c of type %T\n", b, b)
	// byte (uint8)

	x, y := 1, "1"
	fmt.Printf("x=%v, y=%v\n", x, y)
	fmt.Printf("x=%#v, y=%#v\n", x, y) // use #v in debug/log, shows type

	fmt.Printf("%20s!\n", s) // prints 20 spaces then s variable

	fmt.Println(isPalindrome("g"))
	fmt.Println(isPalindrome("go"))
	fmt.Println(isPalindrome("gog"))
	fmt.Println(isPalindrome("g☺g"))
}

// isPalindrome("g") -> true
// isPalindrome("go") -> false
// isPalindrome("gog") -> true
// isPalindrome("gogo") -> true
func isPalindrome(s string) bool {
	rs := []rune(s) // get slice of rune, makes code unicode aware
	for i := 0; i < (len(s))/2; i++ {
		if rs[i] != rs[len(rs)-1-i] {
			return false
		}
	}
	return true
}

func banner(text string, width int) {
	padding := (width - utf8.RuneCountInString(text)) / 2
	// padding := (width - len(text)) / 2 / 2 // BUG: len is in bytes
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
