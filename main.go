package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
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

 - A card struct, which has properties including the number, suite, and colour


*/

type card struct {
	colour string
	value  int
	name   string
	suite  string
}

func NewDeck() []card {
	var Deck []card
	for j := 0; j < 4; j++ {

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

			}

			cardAsString := strings.ToTitle(strconv.Itoa(i))
			switch i {
			case 11:
				card := card{
					colour: colour,
					suite:  suite,
					value:  10,
					name:   fmt.Sprintf("Jack Of %s", suite),
				}

				Deck = append(Deck, card)

			case 12:
				card := card{
					colour: colour,
					suite:  suite,
					value:  10,
					name:   fmt.Sprintf("Queen Of %s", suite),
				}
				Deck = append(Deck, card)

			case 13:
				card := card{
					colour: colour,
					suite:  suite,
					value:  10,
					name:   fmt.Sprintf("King Of %s", suite),
				}
				Deck = append(Deck, card)

			case 14:
				card := card{
					colour: colour,
					suite:  suite,
					value:  11,
					name:   fmt.Sprintf("Ace Of %s", suite),
				}
				Deck = append(Deck, card)

			default:
				if i >= 2 {
					card := card{
						colour: colour,
						suite:  suite,
						value:  i,
						name:   fmt.Sprintf("%s Of %s", cardAsString, suite),
					}
					Deck = append(Deck, card)
				}
			}
		}

	}

	return Deck
}

// For this we will implement the Fisher Yates shuffle.

/*
Fisher-Yates Shuffle Algorithm:
Input: An array or list of elements

1. For i from n - 1 to 1, where n is the number of elements in the array:
 1. Generate a random integer j such that 0 <= j <= i
 2. Swap the element at index i with the element at index j

Output: The shuffled array or list of elements
*/

// Let's try splitting the deck in 2, shuffling, then recombining the 2 halves
func Shuffle(Deck []card) []card {
	var outputDeck = NewDeck()

	var mid = len(Deck) / 2
	var half1 = Deck[0:mid]
	var half2 = Deck[mid:]

	var outHalf1 = outputDeck[0:25]
	var outHalf2 = outputDeck[26:51]

	n1 := len(half2)
	n2 := len(half2)

	// Get the current time as seed
	seed := time.Now().UnixNano()

	// Create a new Rand instance with the current time as seed
	r := rand.New(rand.NewSource(seed))

	// Shuffle the first half
	for i := n1 - 1; i > 0; i-- {
		// Generate a random index from the remaining unshuffled elements
		j := r.Intn(i)

		// Swap the current element with the randomly selected element
		half1[i], outHalf1[j] = outHalf1[j], half1[i]
	}

	// Shuffle the second half
	for i := n2 - 1; i > 0; i-- {
		// Generate a random index from the remaining unshuffled elements
		j := r.Intn(i)

		// Swap the current element with the randomly selected element
		half2[i], outHalf2[j] = outHalf2[j], half2[i]
	}

	// stick the two halves back together
	outputDeck = append(half1, half2...)

	return outputDeck
}

func Deal(Deck []card) []card {
	var hand []card

	for i := 0; i < 5; i++ {
		hand = append(hand, Deck[i])
	}

	return hand
}

func main() {
	var InputDeck = NewDeck
	var TestDeck = Shuffle(InputDeck())

	for _, card := range TestDeck {
		fmt.Println(card)
	}
}
