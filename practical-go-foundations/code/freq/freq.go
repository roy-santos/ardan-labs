package main

// Q: What is the most common word (ignoring case) in sherlock.txt?
// Word frequency

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("freq/sherlock.txt")

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	defer file.Close()

	w, err := mostCommon(file, 3)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(w)
	// mapDemo()

	/*
		// path := "C:\to\new\report.csv" // in this format, "\" is an escape character, "\t", "\n", "\r" get interpreted
		path := `C:\to\new\report.csv` // raw string, \ is just a \
		fmt.Println(path)
	*/
}

func mostCommon(r io.Reader, n int) ([]string, error) {
	freqs, err := wordFrequency(r)
	if err != nil {
		return nil, err
	}

	// instantiate a slice of keys and populate with all keys
	keys := make([]string, 0, len(freqs))
	for key := range freqs {
		keys = append(keys, key)
	}

	// sort keys slice based on its value in the freqs map
	sort.SliceStable(keys, func(i, j int) bool {
		return freqs[keys[i]] > freqs[keys[j]]
	})

	// instantiate slice of n common words and fill with top n words from keys slice
	nCommonStrings := make([]string, n)
	for i := 0; i < n; i++ {
		nCommonStrings[i] = keys[i]
	}

	return nCommonStrings, nil
}

/*	You can use raw strings to create multi line strings
var request = `GET /ip HTTP/1.1
Host: httpbin.org
Connection: Close

`
*/

// "Who's on first?" -> [Who s on first]
var wordRe = regexp.MustCompile(`[a-zA-Z]+`) // package level variable, executed BEFORE main fxn

/* Will run before main as well
func init() {
	// ...
}
*/

/*
func mapDemo() {
	var stocks map[string]float64 // word -> count
	sym := "TTWO"
	price := stocks[sym]
	fmt.Printf("%s -> $%.2f\n", sym, price)

	if price, ok := stocks[sym]; ok {
		fmt.Printf("%s -> $%.2f\n", sym, price)
	} else {
		fmt.Printf("%s not found\n", sym)
	}


	// stocks = make(map[string]float64)
	// stocks[sym] = 136.73


	stocks = map[string]float64{
		sym:    136.73,
		"AAPL": 172.35,
	}

	if price, ok := stocks[sym]; ok {
		fmt.Printf("%s -> $%.2f\n", sym, price)
	} else {
		fmt.Printf("%s not found\n", sym)
	}

	for k := range stocks { // keys
		fmt.Println(k)
	}

	for k, v := range stocks { // key & value
		fmt.Println(k, "->", v)
	}

	for _, v := range stocks { // values
		fmt.Println(v)
	}

	delete(stocks, "AAPL")
	fmt.Println(stocks)
	delete(stocks, "AAPL") // trying to delete unexisting key, nothing happens, NO PANIC

}
*/

func maxWord(freqs map[string]int) (string, error) {
	if len(freqs) == 0 {
		return "", fmt.Errorf("empty map")
	}

	maxN, maxW := 0, ""
	for word, count := range freqs {
		if count > maxN {
			maxN, maxW = count, word
		}
	}

	return maxW, nil
}

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freqs := make(map[string]int) // word -> count
	// lnum := 0
	for s.Scan() {
		words := wordRe.FindAllString(s.Text(), -1) // current line
		for _, w := range words {
			freqs[strings.ToLower(w)]++ // if key doesnt exist, returns 0
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	// fmt.Println("num lines:", lnum)
	return freqs, nil
}
