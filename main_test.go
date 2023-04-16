package main

import (
	"testing"
)

/*
	We should probably start by defining what a deck of cards is.

- It should have 52 cards.
- All the cards should be distinct.
- It should have 4 suites, Diamonds, Hearts, Clubs and Spades, each with 13 cards.
- Clubs and Spades should be black, Diamonds and Hearts should be red.
- Each suite should have cards numbered 2 to 10, then Jack, Queen, King, Ace.
- Values should be the same as the card number, then 10 for Jack, Queen and King, and 11 for Ace.
*/
func TestNewDeck(t *testing.T) {
	t.Run("Check the deck has 52 cards", func(t *testing.T) {
		want := 52
		got := 0

		for _, _ = range NewDeck() {
			got += 1
		}
		if got != want {
			t.Errorf("Wanted %v, got %v", want, got)
		}
	})

	t.Run("Check there are no duplicate cards", func(t *testing.T) {
		want := 1
		count := 0

		ourDeck := NewDeck()
		checkDeck := NewDeck()

		for _, ourCard := range ourDeck {
			for _, checkCard := range checkDeck {
				if checkCard == ourCard {
					count++
				}
			}
			// Once we have looped through checking our card against each card in the check deck, we then check if the counter went above 1. If so we have a duplicate and return false
			if count != want {
				t.Errorf("Wanted %v, got %v", want, count)
			}
			// We reset the "got" counter for the next card
			count = 0
		}
	})

	t.Run("It should have 4 suites, Diamonds, Hearts, Clubs and Spades, each with 13 cards", func(t *testing.T) {
		DiamondsWanted := 13
		HeartsWanted := 13
		ClubsWanted := 13
		SpadesWanted := 13

		DiamondsActual := 0
		HeartsActual := 0
		ClubsActual := 0
		SpadesActual := 0

		for _, card := range NewDeck() {
			switch card.suite {
			case "Diamonds":
				DiamondsActual++

			case "Hearts":
				HeartsActual++

			case "Clubs":
				ClubsActual++

			case "Spades":
				SpadesActual++

			default:
				t.Errorf("Expected one of the four suites, got %s", card.suite)
			}
		}

		if DiamondsActual != DiamondsWanted {
			t.Errorf("Expected %v Diamonds, got %v", DiamondsWanted, DiamondsActual)
		}
		if HeartsActual != HeartsWanted {
			t.Errorf("Expected %v Hearts, got %v", HeartsWanted, HeartsActual)
		}
		if ClubsActual != ClubsWanted {
			t.Errorf("Expected %v Clubs, got %v", ClubsWanted, ClubsActual)
		}
		if SpadesActual != SpadesWanted {
			t.Errorf("Expected %v Spades, got %v", SpadesWanted, SpadesActual)
		}
	})

	t.Run("Clubs and Spades should be black, Diamonds and Hearts should be red", func(t *testing.T) {

		for _, card := range NewDeck() {
			switch card.suite {
			case "Clubs", "Spades":
				if card.colour != "Black" {
					t.Errorf("Expected Black for %v suite", card.suite)
				}
			case "Hearts", "Diamonds":
				if card.colour != "Red" {
					t.Errorf("Expected Red for %v suite", card.suite)
				}
			}
		}

	})

}
