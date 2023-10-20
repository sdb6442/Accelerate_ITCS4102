// Term Project Assignment
// We elected to create one program that satisfies the three programming concepts.
// The following program uses the language 'Go' to create a deck of cards and manipulate the deck to satisfy
// the assignment ruberic.
// Group Accelerate: Brian Samuels, Samuel Baynes, Tristan Hall, Jasmine Kingg, Alex Johnston
// Date: 10/19/2023

// Program 1:
// The 8 methods use different functions from the standard library in the main method of this program.
//Program 2:
//The assignment contains various for loops/if else statements, and uses struct for the cards, and an array for the deck of cards
//Program 3:
//The method findCardIndexByRune handles errors and a new play initialization that handles errors.

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// PROGRAM 2: Creates deck of cards (Brian)
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
				color:      rune(cardColor),
			}
			deck = append(deck, card)
		}
	}
	return deck
}

//PROGRAM 3: Card User exception handling (Tristan)

type cardMaker struct {
	value      int
	suit       string
	isFaceCard bool
	color      rune
}
type player struct {
	name   string
	hand   []cardMaker
	status bool
}
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

type hand []player
type Deck []cardMaker

func New(text string) error {
	return &errorString{text}
}
func newPlayer() (Nplayer player) {
	fmt.Println("Enter Name: ")
	fmt.Scanln(&Nplayer.name)
	if Nplayer.name == "" {
		errors.New("Name is empty")
	}
	Nplayer.status = true
	return Nplayer
}
func cngStatus(player player) {
	if player.status == true {
		player.status = false
	} else {
		player.status = true
	}
	fmt.Println(player.name, " Player status:", player.status)
}

func main() {
	deck := newDeck()
	player1 := newPlayer()
	player1.hand = make([]cardMaker, 0, 7)

	// Print deck of cards
	fmt.Println("Data Structures(Deck of cards, struct and arrays): ")
	fmt.Println(deck)

	// PROGRAM 1: 8 Methods showcasing data manipulation

	// Two int Methods:
	// 1. (Samuel) - Max card value when comparing two cards - max()
	card1 := deck[1]
	card2 := deck[13]
	maxCard := max(card1.value, card2.value)

	fmt.Println("\nInt Method 1(comparing max card values,max()):")
	fmt.Println("Card 1: ", card1, " Card 2: ", card2, " Value of highest card: ", maxCard)

	// 2. (Samuel) - Min card value when comparing two cards - min()
	minCard := min(card1.value, card2.value)
	fmt.Println("\nInt Method 2(comparing min card values, min()):")
	fmt.Println("Card 1: ", card1, " Card 2: ", card2, " Value of lowest card: ", minCard)

	// Two String Methods:
	// 1. (Samuel) - Count the number of Heart cards by creating a string consisting of each occurance of
	//    "Heart", then counting substrings of "Heart" using strings.Count() method
	var heartString string
	for i := 0; i < len(deck); i++ {
		if deck[i].suit == "Heart" {
			heartString += deck[i].suit
		}
	}
	heartCount := strings.Count(heartString, "Heart")
	fmt.Println("\nString Method 1(counting Heart cards, Count()):")
	fmt.Println("Number of Heart cards: ", heartCount)

	// 2. (Jasmine)  -  Compare() two cards suits and return if theyre similar or not.
	// If they are different, it will move onto the next card until there is a match.
	var currentSuit string
	var nextSuit string
	var suitMatcher int

	fmt.Println("\nString Method 2(comparing card suits, Compare()):")
	for i := 0; i < len(deck); i++ {
		currentSuit = deck[i].suit
		for j := i + 1; j < len(deck); {
			nextSuit = deck[j].suit
			fmt.Println("Card 1:", currentSuit)
			fmt.Println("Card 2:", nextSuit)
			suitMatcher = strings.Compare(currentSuit, nextSuit)
			if suitMatcher != 0 {
				fmt.Println("These suits are different. Finding next card...")
				j++
			}
			if suitMatcher == 0 {
				fmt.Println("We found a match!")
				break
			}
		}
		break
	}

	// Two Boolean Methods:
	// 1. (Jasmine)  - Returns boolean if it contains() the red card indicator or not
	var isRed bool
	fmt.Println("\nBoolean Method 1(checking if card is red or black, Contains()):")
	for i := 0; i < len(deck); i++ {
		isRed = strings.Contains(string(deck[i].color), "R")
		fmt.Println(deck[i])
		if isRed == true {
			fmt.Println("This card is red.")
			break
		}
		if isRed == false {
			fmt.Println("This card is black.")
		}

	}

	// 2. (Alex) - Picks a random card and determines if it is a face card. Uses a random function from the Standard library.
	//The method i used is located under the main method.
	randomCard := deck.getRandomCard()

	fmt.Println("\nBoolean Method 2(checking if card is a face card, Intn()):")
	if randomCard.isFaceCard {
		fmt.Println("This card is a face card")
	} else {
		fmt.Println("This card is not a face card")
	}

	// Two Rune Methods:
	// (Alex) 1. Determines what the index of the first card of a selected rune. Uses B (Black) and C (Invalid input) as examples
	//Uses the findcardindexbyrune method, which is also located under the main mehthod. uses HasPrefix() function from standard library
	targetRune := 'B' //Blue

	index, err := deck.findCardIndexByRune(targetRune)

	fmt.Println("\nRune Method 1(determine index of card rune, HasPrefix()):")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The first card of the color '%c' is at index %d\n", targetRune, index)
	}
	//Showing the error handling when an invalid rune is used
	targetRune = 'C' //Cyan

	fmt.Println("\nexception handling(find card index by rune):")
	index, err = deck.findCardIndexByRune(targetRune)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The first card of the color '%c' is at index %d\n", targetRune, index)
	}

	// 2. (Alex) - Picks 5 random cards using time.now and rand.new(source)
	// then displays the 5 associated color runes
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomColors := []rune{}
	for i := 0; i < 5; i++ {
		randomIndex := r.Intn(len(deck))
		randomColors = append(randomColors, deck[randomIndex].color)
	}
	fmt.Println("\nRune Method 2(display random card rune color, Intn()):")
	fmt.Printf("Randomly selected 5 card colors: %c %c %c %c %c\n", randomColors[0], randomColors[1], randomColors[2], randomColors[3], randomColors[4])

}

// (Alex) getRandomCard() Generates a random card for boolean method #2

func (d Deck) getRandomCard() cardMaker {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomIndex := r.Intn(len(d))
	fmt.Println("\nexception handling(generate random card):")
	print("The index of the random card was: ", randomIndex, "\n")
	return d[randomIndex]
}

// (Alex) findCardIndexByRune method for the rune methods. Contains exception handling for program 3.
func (d Deck) findCardIndexByRune(targetRune rune) (int, error) {
	targetRuneStr := string(targetRune) // Convert targetRune to a string
	for i, card := range d {
		cardColorStr := string(card.color) // Convert card.color to a string
		if strings.HasPrefix(cardColorStr, targetRuneStr) {
			return i, nil
		}
	}

	return -1, fmt.Errorf("No card of color '%c' found in the deck", targetRune)
}
