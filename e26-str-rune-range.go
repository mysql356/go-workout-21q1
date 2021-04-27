package main

import . "fmt"

func main() {

	a := "mñ"

	for i, v := range a {
		Printf("Index %d, Char %c, Hex %x \n", i, v, v) //i-int v-rune
	}
}
