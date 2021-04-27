package main

import . "fmt"

func main() {

	a := map[string]int{
		"xx": 10,
		"yy": 20,
		"zz": 30,
	}

	//before delete
	if v, ok := a["yy"]; ok {
		Println(v, ok)
	}
	for k, v := range a {
		Println(k, v)
	}

	//delete
	delete(a, "yy")

	//after delete
	if v, ok := a["yy"]; ok {
		Println(v, ok)
	}

	for k, v := range a {
		Println(k, v)
	}

}
