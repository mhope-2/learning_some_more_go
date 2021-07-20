package main

import "fmt"

// deck struct
type Deck []string

func (d Deck) Print(){
	for _,card := range d {
		fmt.Println(card)
	}
}