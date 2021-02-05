package main

import "fmt"

func main() {

	e1 := []struct {
		fn  string
		age int
	}{
		{"aa", 30},
		{"bb", 50},
	}

	fmt.Println(e1)

}
