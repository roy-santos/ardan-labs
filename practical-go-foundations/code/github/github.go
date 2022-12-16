package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/roy-santos")
	if err != nil {
		log.Fatalf("error: %s", err)
		/*  this is shortcut to
		log.Printf("error: %s", err)
		os.Exit(1)
		*/
	}
	if resp.StatusCode != http.StatusOK { //Go throws error if http response is invalid, for example status 400s or 500s is valid and will not be considered "error"
		log.Fatalf("error: %s", resp.Status)
	}
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	/*
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			log.Fatalf("error: can't copy - %s", err)
		}
	*/
	var r Reply
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("error: can't decode -%s", err)
	}
	fmt.Println(r)
	fmt.Printf("%#v\n", r) // this way of printing shows types of the values printed

	name, numrepos, err := githubInfo("tebeka")
	if err != nil {
		log.Fatalf("error: can't decode %s", err)
	} else {
		fmt.Printf("%#v, %#v\n", name, numrepos)
	}
}

// githubInfo returns name and number of public repos for login
func githubInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + url.PathEscape(login) // makes sure the value passed in is valid
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%#v - %s", url, resp.Status)
	}

	defer resp.Body.Close() // after content of body is read, this fxn will close the reader

	var r struct { // anonymous struct
		Name     string
		NumRepos int `json:"public_repos"`
	}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}
	return r.Name, r.NumRepos, nil
}

type Reply struct {
	Name     string
	Location string
	Admin    bool
	NumRepos int `json:"public_repos"` // this example uses json field tag so that you don't need to match field names in your struct
}

/* JSON <-> Go
true/false <-> true/false
string <-> string
null <-> nil
number <-> float64, float32, int8, int16, int32, int 64, int, uint8,...
array <-> []any ([]interface{})
object <-> map[string]any, struct

JSON -> io.Reader -> Go: json.Decoder
JSON -> []byte -> Go: json.Unmarshal
Go -> io.Writer -> JSON: json.Encoder
Go -> []byte -> JSON: json.Marshal
*/
