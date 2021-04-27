package main

import . "fmt"

func main() {
	var a map[string]int
	a = make(map[string]int)
	a["aa"] = 10

	b := make(map[string]int)
	b["bb"] = 20

	c := map[string]int{}
	c["cc"] = 30

	var d map[string]int = map[string]int{}
	d["dd"] = 40

	var e map[string]int = map[string]int{"ee": 50}
	e["ff"] = 60

	Println(a)
	Println(b)
	Println(c)
	Println(d)
	Println(e)

}
