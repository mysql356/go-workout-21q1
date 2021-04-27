package main

import (
	. "fmt"
)

func main() {

	s := []string{"aa", "bb"}
	msg(s...)
	Println(s)

}

func msg(s ...string) {

	//	s = append(s, "play", "game") // - not work
	s[0] = "Go"
}
