package main

import "fmt"

func main() {

	//1
	e := struct {
		fn  string
		age int
	}{
		"aa", 30,
	}

	fmt.Println(e)
	//return 

	//2
	e1 := []struct {
		fn  string
		age int
	}{
		{"aa", 30},
		{"bb", 50},
		{"cc", 70},
	}

	fmt.Println(e1)
	fmt.Println(e1[0].fn, e1[1].age,e1[2].fn)
	//return

	//3
	for i:=0; i< len(e1) ; i++ {
		fmt.Println(e1[i].fn, e1[i].age)
	}
}
