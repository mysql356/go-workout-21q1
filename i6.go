package main

import "fmt"

type Tester interface {
	Test()
}

type customFloat64 float64

func (m customFloat64) Test() {
	fmt.Println(m)
}

func describe(t Tester) {
	fmt.Printf("\n%T %v", t, t)
}

func main() {
	var t Tester
	f := customFloat64(10.5)
	t = f

	describe(t)
	t.Test()

}
