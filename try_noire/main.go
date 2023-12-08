package main

import (
	"fmt"
	"github.com/teacat/noire"
)

func main() {
	c := noire.NewRGB(255, 255, 255)
	hex := noire.NewHex("000000")
	html := noire.NewHTML("red")
	fmt.Println(html.Invert().Hex()) // Output: 00FFFF
	fmt.Println(html.Invert().RGB()) // Output: 0, 255, 255
	fmt.Println(hex.Invert().HTML()) // Output: White
	fmt.Println(hex.Invert().RGB())  // Output: 255, 255, 255
	fmt.Println(c.Invert().Hex())    // Output: 000000
	fmt.Println(c.Invert().HTML())   // Output: Black
	fmt.Println(c.Lighten(1).RGB())  // Output: 255, 255, 255
}
