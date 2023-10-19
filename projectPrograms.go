
package main

import (
	"fmt"
	"errors"
	//"math"
	//"string"
)

type cardMaker struct {
	value      int
	suit       string
	isFaceCard bool
	color      string
}

type Deck []cardMaker

// Creates deck of cards
func newDeck() (deck Deck) {
	//jack = 11 queen=12 king= 13 ace=14
	values := [13]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	//suits include Heart, Diamond, Club, Spade
	suits := [4]string{"Heart", "Diamond", "Club", "Spade"}

	var cardColor rune
	var faceCard bool

	for i := 0; i < len(values); i++ {
		for j := 0; j < len(suits); j++ {
			if values[i] > 10 {
				faceCard = true
			} else {
				faceCard = false
			}
			//Switches card color based on suite. 82 = 'R'(Red)   66= 'B' (Black)
			if suits[j] == "Heart" || suits[j] == "Diamond" {
				cardColor = 'R'
			} else {
				cardColor = 'B'
			}
			//Brian-Converted rune cardColor to string for output
			card := cardMaker{
				value:      values[i],
				suit:       suits[j],
				isFaceCard: faceCard,
				color:      string(cardColor),
			}
			deck = append(deck, card)
		}
	}
	return deck
}

func main() {
	deck := newDeck()

	fmt.Println(deck)
}

func exceptionHandling error{
	//Jasmine
	if {
		return errors.Errorf("Has failed")
	}
	else{
		
	}
}

