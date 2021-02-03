package main

import "fmt"


func main() {
	a := [...]string{"aa", "bb", "cc"}
	b := a
	b[0] = "us"
	for k, v := range a {
		fmt.Println(k, "----", v)
	}

	fmt.Println(b)
	Test()
}
