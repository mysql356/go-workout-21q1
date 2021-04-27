package main

import (
	"fmt"
)

func main() {

	a := []int{10, 20, 30, 40, 50}

	x := a[0:0]
	fmt.Println(len(x), cap(x)) //0 5

	x1 := a[1:1]
	fmt.Println(len(x1), cap(x1)) //0 4

	x2 := a[2:2]
	fmt.Println(len(x2), cap(x2)) //0 3

	x3 := a[3:3]
	fmt.Println(len(x3), cap(x3)) //0 2

	x4 := a[4:4]
	fmt.Println(len(x4), cap(x4)) //0 1

	x5 := a[5:5]
	fmt.Println(len(x5), cap(x5)) //0 0

	//x6 := a[6:6]
	//fmt.Println(len(x6), cap(x6)) //panic: runtime error: slice bounds out of range [:6] with capacity 5

	x7 := a[:]
	fmt.Println(len(x7), cap(x7)) //5 5

	x8 := a[1:]
	fmt.Println(len(x8), cap(x8)) //4 4

	x9 := a[:5]
	fmt.Println(len(x9), cap(x9)) //5 5

}
