package main

import . "fmt"

func main() {

	a := []string{"aa", "bb"}
	b := []string{"xx", "yy"}

	c := append([]string{}, a...)
	c = append(c, b...)
	d := append(c, "dd")
	Println(d)

}
