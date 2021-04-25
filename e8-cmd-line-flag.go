package main

import (
	"flag"
	"fmt"
)

/*
-flag
-flag=x
-flag x  // non-boolean flags only

One or two minus signs may be used; they are equivalent.
The last form is not permitted for boolean flags because the meaning of the command

cmd -x *
where * is a Unix shell wildcard, will change if there is a file called 0, false, etc.
You must use the -flag=false form to turn off a boolean flag.
Flag parsing stops just before the first non-flag argument ("-" is a non-flag argument) or after the terminator "--".
Integer flags accept 1234, 0664, 0x1234 and may be negative. Boolean flags may be:
1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False
*/

func main() {

	//case-1 : -word india -num 50 -fork=true  // --word india --num 50 --fork=true
    wordPtr := flag.String("word", "foo", "a string")
    numPtr := flag.Int("num", 10, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

	//case-2 : -svar delhi -ivar 100
    var strVar string
    flag.StringVar(&strVar, "svar", "bar", "a string var")

	var intVar int64 
	flag.Int64Var(&intVar, "ivar", 20, "a string var")

    flag.Parse()

    fmt.Println("word:", *wordPtr)
    fmt.Println("num:", *numPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", strVar)
	fmt.Println("ivar:", intVar)
    fmt.Println("tail:", flag.Args())
}