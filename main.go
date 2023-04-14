package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	We are creating a simple deck of cards. It must have the following functionality

newDeck()
- Create a list of playing cards. Essentially an array of strings.

print()
- Log out the contents of a deck of cards.

shuffle()
- Shuffles all the cards in a deck.

deal()
- Create a 'hand' of cards.

saveToFile()
- Save a list of cards to a file on the local machine.

newDeckFromFile()
- Load a list of cards from the local machine.



Objects we will require:

 - A suite struct, with numbers of the cards, colour, and picture cards, and whether the card has been dealt
 - A card struct, which has properties including the number, suite, and colour


*/

type card struct {
	colour string
	value  int
	name   string
}

var spadeSuite = []card{
	{
		colour: "Black",
		value:  2,
		name:   "Two Of Spades",
	},

	{
		colour: "Black",
		value:  3,
		name:   "Three Of Spades",
	},

	{
		colour: "Black",
		value:  4,
		name:   "Four Of Spades",
	},

	{
		colour: "Black",
		value:  5,
		name:   "Five Of Spades",
	},

	{
		colour: "Black",
		value:  6,
		name:   "Six Of Spades",
	},

	{
		colour: "Black",
		value:  7,
		name:   "Seven Of Spades",
	},

	{
		colour: "Black",
		value:  8,
		name:   "Eight Of Spades",
	},

	{
		colour: "Black",
		value:  9,
		name:   "Nine Of Spades",
	},

	{
		colour: "Black",
		value:  10,
		name:   "10 Of Spades",
	},

	{
		colour: "Black",
		value:  10,
		name:   "Jack Of Spades",
	},

	{
		colour: "Black",
		value:  10,
		name:   "Queen Of Spades",
	},

	{
		colour: "Black",
		value:  10,
		name:   "King Of Spades",
	},

	{
		colour: "Black",
		value:  11,
		name:   "Ace Of Spades",
	},
}

type suite struct {
}

func createSpadeSuite() []card {
	var Deck []card
	for j := 0; j < 3; j++ {

	}
	for i := 1; i < 15; i++ {
		var suite string
		var colour string
		switch j {
		case 0:
			suite = "Spades"
			colour = "Black"

		case 1:
			suite = "Hearts"
			colour = "Red"

		case 2:
			suite = "Clubs"
			colour = "Black"

		case 3:
			suite = "Diamonds"
			colour = "Red"

		default:

			cardAsString := strings.ToTitle(strconv.Itoa(i))
			switch i {
			case 11:
				card := card{
					colour: colour,
					value:  10,
					name:   fmt.Sprintf("Jack Of %s", suite),
				}
				Deck = append(Deck, card)

			case 12:
				card := card{
					colour: colour,
					value:  10,
					name:   fmt.Sprintf("Queen Of %s", suite),
				}
				Deck = append(Deck, card)

			case 13:
				card := card{
					colour: colour,
					value:  10,
					name:   fmt.Sprintf("King Of %s", suite),
				}
				Deck = append(Deck, card)

			case 14:
				card := card{
					colour: colour,
					value:  11,
					name:   fmt.Sprintf("Ace Of %s", suite),
				}
				Deck = append(Deck, card)

			default:
				card := card{
					colour: colour,
					value:  i,
					name:   fmt.Sprintf("%s Of %s", cardAsString, suite),
				}
				Deck = append(Deck, card)
			}
		}

	}
	return Deck
}

func main() {

	for _, i := range createSpadeSuite() {
		fmt.Println(i)
	}
}
