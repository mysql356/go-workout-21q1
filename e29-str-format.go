package main

import . "fmt"

func main() {
	var a string = "ab"

	Printf("%s %T \n", a, a)

	for i := 0; i < len(a); i++ {
		Printf("%T %c %x %o %b %v", a[i], a[i], a[i], a[i], a[i], a[i])
		Println()
	}
}
