package main

import "fmt"

// Create a new type (custom data type)
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
