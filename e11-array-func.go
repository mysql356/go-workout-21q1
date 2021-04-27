package main

import (
	. "fmt"
)

func main() {
	a := [3]string{"aa", "bb", "cc"}

	//Wrong way :  It'll work, but don't use.
	arrPrint(&a)

	//Right way : Use slice instead of array.
	b := a[:]
	slicePrint(b)
}

func arrPrint(a *[3]string) {
	for k, v := range a {
		Printf("%d ==== %s\n", k, v)
	}
}

func slicePrint(a []string) {
	for k, v := range a {
		Printf("%d ==== %s\n", k, v)
	}
}
