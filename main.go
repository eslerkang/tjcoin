package main

import (
	"fmt"

	"github.com/eslerkang/tjcoin/person"
)

func main() {
	nico := person.Person{}
	nico.SetDetails("kang", 213)
	fmt.Println(nico)
}