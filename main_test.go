package main

import (
	"strconv"
	"strings"
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
	t.Run("Each suite should have cards numbered 2 to 10, then Jack, Queen, King, Ace.", func(t *testing.T) {

		Expected := 4
		JackCount := 0
		QueenCount := 0
		KingCount := 0
		AceCount := 0

		for _, card := range NewDeck() {

			// Check if Atoi returns an error. If so, then we know it's not a number card, and proceed accordingly. This is ok because strings.Split does not return an error.

			number, err := strconv.Atoi(strings.Split(card.name, " ")[0])
			if err != nil {
				actualCard := strings.Split(card.name, " ")[0]

				switch actualCard {
				case "Jack":
					JackCount++
				case "Queen":
					QueenCount++
				case "King":
					KingCount++
				case "Ace":
					AceCount++
				default:
					t.Errorf("Not a numbered card, and not an accepted picture card. Got %v instead.", actualCard)
				}

			}

			// If the test reaches this point, either it is an accepted picture card (which would have meant err != nil) or it is a number card.
			// err being nil means it is definitely a number card
			if err == nil && (number < 2 || number > 10) {
				t.Errorf("Number must be between 2 and 10. Got %v. Card was %v", number, card.name)
			}

		}
		if JackCount != Expected {
			t.Errorf("Expected %v Jacks, got %v", Expected, JackCount)
		}
		if QueenCount != Expected {
			t.Errorf("Expected %v Queens, got %v", Expected, QueenCount)
		}
		if KingCount != Expected {
			t.Errorf("Expected %v King, got %v", Expected, KingCount)
		}
		if AceCount != Expected {
			t.Errorf("Expected %v Aces, got %v", Expected, AceCount)
		}

	})

	t.Run("Values should be the same as the card number, then 10 for Jack, Queen and King, and 11 for Ace.", func(t *testing.T) {
		pictureValue := 10
		aceValue := 11
		for _, card := range NewDeck() {

			// Check if Atoi returns an error. If so, then we know it's not a number card, and proceed accordingly. This is ok because strings.Split does not return an error.

			number, err := strconv.Atoi(strings.Split(card.name, " ")[0])
			if err != nil {
				actualCard := strings.Split(card.name, " ")[0]

				switch actualCard {
				case "Jack":
					if card.value != pictureValue {
						t.Errorf("Jack should have value of %v. Got %v", pictureValue, card.value)
					}

				case "Queen":
					if card.value != pictureValue {
						t.Errorf("Queen should have value of %v. Got %v", pictureValue, card.value)
					}
				case "King":
					if card.value != pictureValue {
						t.Errorf("King should have value of %v. Got %v", pictureValue, card.value)
					}
				case "Ace":
					if card.value != aceValue {
						t.Errorf("Ace should have value of %v. Got %v", aceValue, card.value)
					}

				}
				// If picture card with no error, or number card, then...
				if err == nil && (number != card.value) {
					t.Errorf("Card value is %v but card number is %v. These should be the same.", card.value, number)
				}
			}
		}
	})

}
