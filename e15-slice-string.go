package main

import (
	. "fmt"
)

func main() {

	s := [...]string{"aa", "bb", "cc", "dd", "ee"}

	s1 := s[2:4]
	Println("len", len(s1)) // 4-2 = 2
	Println("cap", cap(s1)) // 5-2 = 3
	Println(s1)             //s[2], s[3] : [cc dd]

}
