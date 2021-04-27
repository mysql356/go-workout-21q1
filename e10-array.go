package main

import "fmt"

func main() {

	//1
	var a [5]int
	a[0] = 10
	a[1] = 20
	a[2] = 30
	a[3] = 40
	a[4] = 50

	fmt.Println(a)

	for k, v := range a {
		fmt.Println(k, "---", v)
	}

	//2
	var a1 [3]int = [3]int{10, 20, 30}
	fmt.Println("a1===============", a1)

	//3
	b := [3]int{10, 20, 30}
	fmt.Println("b==============", b)

	//4
	c := [...]int{10, 20, 30}
	fmt.Println("c===================", c)

	//5
	arrPrint(b)
}

func arrPrint(a [3]int) {

	for k, v := range a {
		fmt.Println(k, "===", v)
	}

}
