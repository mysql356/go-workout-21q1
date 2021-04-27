package main

import (
	"fmt"
	. "fmt"
	"unicode/utf8"
)

func main() {

	var a string = "aaa bbb"
	Println("Length =", len(a))

	word1 := "Señor"
	length(word1)       //5
	Println(len(word1)) //6

	Println("========")

	word2 := "Pets"
	length(word2)       //4
	Println(len(word2)) //4
}

func length(s string) {
	fmt.Println(utf8.RuneCountInString(s))
}
