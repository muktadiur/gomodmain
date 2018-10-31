package main

import (
	"fmt"

	"github.com/muktadiur/gomodone"
	"github.com/muktadiur/gomodtwo"
)

func main() {
	fmt.Println(gomodone.Version())
	fmt.Println(gomodone.ModTwoVersion())
	fmt.Println(gomodtwo.Version())
	fmt.Println("This is from main module")
}
