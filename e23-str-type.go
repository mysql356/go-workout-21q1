package main

import (
	. "fmt"
)

func main() {
	compare()
	types()
}
func compare() {
	var a string = "aa"
	type myString string
	var b myString = "aa"
	//b = a - cannot use a (type string) as type myString in assignment
	b = myString(a)
	Printf("%s %T \n", a, a) //aa string
	Printf("%s %T \n", b, b) //aa main.myString
}

//string
func types() {
	byteSlice := []byte("Welcome")
	Println(byteSlice)
	Printf("%s, %T \n", byteSlice, byteSlice)

	uint8Slice := []uint8("Welcome")
	Println(uint8Slice)

	runeSlice := []rune("Welcome")
	Println(runeSlice)

	int32Slice := []int32("Welcome")
	Println(int32Slice)
}
