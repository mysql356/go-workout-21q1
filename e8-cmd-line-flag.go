package main

import (
	"flag"
	"fmt"
)

func main() {

	//case-1
    wordPtr := flag.String("word", "foo", "a string")
    numbPtr := flag.Int("numb", 10, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

	//case-2
    var strVar string
    flag.StringVar(&strVar, "svar", "bar", "a string var")

	var intVar int64
	flag.Int64Var(&intVar, "ivar", 20, "a string var")

    flag.Parse()

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", strVar)
	fmt.Println("ivar:", intVar)
    fmt.Println("tail:", flag.Args())
}