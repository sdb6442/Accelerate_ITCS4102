// Term Project Assignment
// We elected to create one program that satisfies the three programming concepts.
// The following program uses the language 'Go' to create a deck of cards and manipulate the deck to satisfy
// the assignment ruberic.
// Group Accelerate: Brian Samuels, Samuel Baynes, Tristan Hall, Jasmine Kingg, Alex Johnston
// Date: 10/19/2023

// Program 1:
//TODO:

package main

import (
	"os"
	"fmt"
	"strings"
	//"math"
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

	// Print deck of cards
	fmt.Println("Deck of cards: ")
	fmt.Println(deck)

	// 8 Methods showcasing data manipulation

	// Two int Methods:
	// 1. Max card value when comparing two cards - max()
	card1 := deck[1]
	card2 := deck[13]
	maxCard := max(card1.value, card2.value)
	fmt.Println("\nCard 1: ", card1, " Card 2: ", card2, " Value of highest card: ", maxCard)

	// 2. Min card value when comparing two cards - min()
	minCard := min(card1.value, card2.value)
	fmt.Println("\nCard 1: ", card1, " Card 2: ", card2, " Value of lowest card: ", minCard)

	// Two String Methods:
	// 1. Count the number of Heart cards by creating a string consisting of each occurance of
	//    "Heart", then counting substrings of "Heart" using strings.Count() method
	var heartString string
	for i := 0; i < len(deck); i++ {
		if deck[i].suit == "Heart" {
			heartString += deck[i].suit
		}
	}
	heartCount := strings.Count(heartString, "Heart")
	fmt.Println("\nNumber of Heart cards: ", heartCount)


	// 2. Compare two cards suits and return if theyre similar or not.
	// If they are different, it will move onto the next card until there is a match.
	var currentSuit string
	var nextSuit string
	var suitMatcher int
	for i := 0; i < len(deck); i++ {
		currentSuit = deck[i].suit
		for j := i+1; j < len(deck);{
			nextSuit = deck[j].suit
			fmt.Println("Card 1:", currentSuit)
			fmt.Println("Card 2:", nextSuit)
			suitMatcher = strings.Compare(currentSuit, nextSuit)
			if suitMatcher != 0 {
				fmt.Println("These suits are different. Finding next card...")
					j++
			}
			if suitMatcher == 0{
				 fmt.Println("We found a match!")
				os.exit(0)
			}
		}
	}
	
	
	

	// Two Boolean Methods:
	// 1.
	// 2.

	// Two Rune Methods:
	// 1.
	// 2.

}

func exceptionHandling error{
	//Jasmine
	// if {
	// 	return errors.Errorf("Has failed")
	// }
	// else{

	// }
}

