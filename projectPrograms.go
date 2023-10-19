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
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type cardMaker struct {
	value      int
	suit       string
	isFaceCard bool
	color      rune
}
type player struct{
	name 	string
	hand[] cardMaker
	status 	bool
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
func newPlayer()(Nplayer player){
	fmt.Println("Enter Name: ")
	fmt.Scanln(&Nplayer.name)
	if Nplayer.name ==""{
	    errors.New("Name is empty")
	}
	Nplayer.status = true
	return Nplayer
}
func cngStatus (player player){
	if player.status == true {
		player.status = false
	}else{
		player.status = true
	}
	fmt.Println(player.name," Player status:",player.status)
}


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
				color:      rune(cardColor),
			}
			deck = append(deck, card)
		}
	}
	return deck
}

func main() {
	var turn float32 
	deck := newDeck()
	player1 := newPlayer()
	player1.hand = make([]cardMaker,0,7)

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


	// 2. Jasmine  -  Compare() two cards suits and return if theyre similar or not.
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
				 break
			}
		}
		break
	}
	
	

	// Two Boolean Methods:
	// 1. 


	// 2. Jasmine  - Returns boolean if contains the red card indicator or not




	// Two Rune Methods:
	// 1.
	targetRune := 'B' //Blue

	index, err := deck.findCardIndexByRune(targetRune)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The first card of the color '%c' is at index %d\n", targetRune, index)
	}
	//Showing the error handling when an invalid rune is used
	targetRune = 'C' //Cyan

	index, err = deck.findCardIndexByRune(targetRune)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The first card of the color '%c' is at index %d\n", targetRune, index)
	}
	// 2. Picks 5 random cards using time.now and rand.new(source) then displays the 5 associated color runes
	//
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomColors := []rune{}
	for i := 0; i < 5; i++ {
		randomIndex := r.Intn(len(deck))
		randomColors = append(randomColors, deck[randomIndex].color)
	}
	fmt.Printf("Randomly selected 5 card colors: %c %c %c %c %c\n", randomColors[0], randomColors[1], randomColors[2], randomColors[3], randomColors[4])

}

//findCardIndexByRune method. Contains exception handling for program 3.
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
