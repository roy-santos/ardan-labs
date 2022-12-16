package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	sig, err := shaSum("http.log.gz")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)

	sig, err = shaSum("sha1.go")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)
}

/*
if file names ends with .gz

	cat http.log.gz| gunzip | shasum

else

	ca http.log.gz | shasum
*/
func shaSum(fileName string) (string, error) {
	// idiom: acquire a resource, check for error, defer release
	file, err := os.Open(fileName)
	if err != nil {
		return "", nil
	}
	// multiple defers are called in LIFO order
	defer file.Close() // this will close the file whenver the function exits, happens in fxn level

	var r io.Reader = file
	if strings.HasSuffix(fileName, ".gz") {
		gz, err := gzip.NewReader(file) // uncompress file content
		if err != nil {
			return "", err
		}
		defer gz.Close()
		r = gz
	}
	//io.CopyN(os.Stdout, r, 100)
	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), nil
}
