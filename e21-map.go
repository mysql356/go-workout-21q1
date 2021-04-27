package main

import . "fmt"

func main() {

	a := map[string]int{
		"xx": 10,
		"yy": 20,
		"zz": 30,
	}

	if v, ok := a["zz"]; ok {
		Println(v, ok)
	}

	for k, v := range a {
		Println(k, v)
	}

}
