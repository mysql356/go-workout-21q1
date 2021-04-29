package main

import (
	"fmt"
)

func main() {

	byte_string()
	byte_hex()
	byte_dec()
	runes()

}

//string
func byte_string() {
	byteSlice := []byte("Welcome")
	str := string(byteSlice)
	fmt.Println(byteSlice)
	fmt.Println(str)
}

//oct - Café
func byte_hex() {
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)
}

//dec - Café
func byte_dec() {
	byteSlice := []byte{67, 97, 102, 195, 169} //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str := string(byteSlice)
	fmt.Println(str)
}

//rune - Señor
func runes() { //
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str := string(runeSlice)
	fmt.Println(str)
}
