package main

import (
	. "fmt"
)

func main() {

	s := []string{"aa", "bb", "cc"}
	s = append(s, "dd")
	Println(s)

	a := []string{"aa", "bb"}
	b := []string{"xx", "yy"}

	c := append(a, b...)
	Println(c)

	d := append([]string{}, "xx")
	e := append(d, "yy", "zz")
	e = append(e, s...)
	Println(e)
}
