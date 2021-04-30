package main

import (
	. "fmt"
)

func main() {

	//	numeric()
	str()
}

func numeric() {
	//int
	a := 10
	b := []int{10, 20}

	find(10)         //10 []
	find(10, 10)     //10 [10]
	find(10, 10, 20) //10 [10 20]
	find(a, b...)    //10 [10 20]

	find(10, a)    //10 [10]
	find(10, b...) //10 [10 20]
}

func find(x int, y ...int) {
	Println(x, y)
}

//string
func str() {
	//1
	s := []string{"aa", "bb"}
	show(s...)

	//2
	s1 := []string{"mm", "nn"}
	show1(s1...)

	//3
	s2 := []string{"xx", "yy"}
	Println(s2, len(s2), cap(s2))
	show2(s2...)
	Println(s2, len(s2), cap(s2))

}

func show(s ...string) {
	Println(s)
}

func show1(s1 ...string) {
	s1[0] = "mmm"
	//	s1[1] = "nnn"

}

func show2(s1 ...string) {
	s1[0] = "zz"
	s1 = append(s1, "play", "game", "over")
	Println(s1, len(s1), cap(s1))
}
