package main

import . "fmt"

func main() {

	a := "mñ"

	//s := []byte(a)
	s := []rune(a)
	for i := 0; i < len(s); i++ {
		Printf("%c --- %x \n", s[i], s[i])
	}
}
