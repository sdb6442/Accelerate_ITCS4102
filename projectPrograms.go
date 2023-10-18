// Program 1:
//TODO:

// Program 2
//TODO:

// Program 3:
// TODO:
package main

import (
	"fmt"
)

type cardTest struct {
	value      int
	suit       string
	isFaceCard bool
	color      rune
}

type Deck []cardTest

func new() (deck Deck) {
	//card := cardTest{1, "Heart", true, 'R'}

	//jack = 11 queen=12 king= 13 ace=14
	values := [13]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	// Valid suits include Heart, Diamond, Club & Spade
	suits := [4]string{"Heart", "Diamond", "Club", "Spade"}

	//faceCards := []bool{false,true}

	//color:= [] rune{'R','B'}
	var faceCard bool
	faceCard = false
	for i := 0; i < len(values); i++ {
		for n := 0; n < len(suits); n++ {
			if values[i] > 10 {
				faceCard = true
			}
			card := cardTest{
				value:      values[i],
				suit:       suits[n],
				isFaceCard: faceCard,
				color:      'R',
			}
			deck = append(deck, card)
		}
	}
	return deck

	//fmt.Printf("%+v\n", card)

	//fmt.Println(card)
}

func main() {
	deck := new()

	fmt.Println(deck)
}
