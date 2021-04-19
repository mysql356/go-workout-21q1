// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	url := "http://localhost/input/json/api/" //{"a":10, "b":20}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", b) //{"a":10, "b":20}
	fmt.Println(b)        //in byte [123 34 97 34 58 49 48 44 32 34 98 34 58 50 48 125]

	//json-unmarshal
	type val struct {
		Aaa int `json:"a"`
		Bbb int `json:"b"`
	}

	//by val
	var c val
	if err := json.Unmarshal(b, &c); err != nil {
		fmt.Printf("JSON unmarshaling failed - by val: %s", err)
	}
	fmt.Println(c.Aaa, c.Bbb) // 10, 20
	fmt.Println(c)

	//by ref
	d := new(val)
	if err := json.Unmarshal(b, d); err != nil {
		fmt.Printf("JSON unmarshaling failed - by ref: %s", err)
	}
	fmt.Println(d.Aaa, d.Bbb) // 10, 20
	fmt.Println(d)

}
