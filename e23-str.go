package main

import . "fmt"

func main() {

	var a string = "aa"

	type myString string
	var b myString = "aa"

	//b = a - cannot use a (type string) as type myString in assignment
	b = myString(a)

	Printf("%s %T", a, a) //aa string
	Println()
	Printf("%s %T", b, b) //aa main.myString
}
