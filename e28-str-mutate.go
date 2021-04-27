package main

import (
	"fmt"
)

/*
func mutate(s string)string {
    s[0] = 'a'//any valid unicode character within single quote is a rune
    return s
}

func main() {
    h := "hello"
    fmt.Println(mutate(h))
}
*/

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}
func main() {
	h := "hello"
	fmt.Println(mutate([]rune(h)))
}
