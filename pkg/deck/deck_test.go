// FILEPATH: /Users/chris/Documents/Go/1/pkg/deck/deck_test.go

package deck

import (
	"strconv"
	"testing"
)

func TestNewDeck(t *testing.T) {
	types := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	d := NewDeck(types...)

	// Check if the deck has the correct number of cards
	expectedCardCount := len(types) * CARD_COUNT
	if len(d.Cards) != expectedCardCount {
		t.Errorf("Expected deck to have %d cards, but got %d", expectedCardCount, len(d.Cards))
	}

	// Check if the deck has the correct number of face cards
	expectedFaceCardCount := 5
	if len(d.FaceCards) != expectedFaceCardCount {
		t.Errorf("Expected deck to have %d face cards, but got %d", expectedFaceCardCount, len(d.FaceCards))
	}

	// Check if the deck has the correct number of types
	if len(d.Types) != len(types) {
		t.Errorf("Expected deck to have %d types, but got %d", len(types), len(d.Types))
	}

	// Check if the deck has the correct types
	for i, typ := range types {
		if d.Types[i] != typ {
			t.Errorf("Expected deck type at index %d to be %s, but got %s", i, typ, d.Types[i])
		}
	}
}

func TestShuffle(t *testing.T) {
	types := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	d := NewDeck(types...)

	// Shuffle the deck
	d.Shuffle()

	// Check if the deck has the same number of cards
	expectedCardCount := len(types) * CARD_COUNT
	if len(d.Cards) != expectedCardCount {
		t.Errorf("Expected deck to have %d cards after shuffling, but got %d", expectedCardCount, len(d.Cards))
	}

	// Check if the deck has the same number of face cards
	expectedFaceCardCount := 5
	if len(d.FaceCards) != expectedFaceCardCount {
		t.Errorf("Expected deck to have %d face cards after shuffling, but got %d", expectedFaceCardCount, len(d.FaceCards))
	}

	// Check if the deck has the same number of types
	if len(d.Types) != len(types) {
		t.Errorf("Expected deck to have %d types after shuffling, but got %d", len(types), len(d.Types))
	}

	// Check if the deck has the same types
	for i, typ := range types {
		if d.Types[i] != typ {
			t.Errorf("Expected deck type at index %d to be %s after shuffling, but got %s", i, typ, d.Types[i])
		}
	}
}

func TestRemainingCards(t *testing.T) {
	types := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	d := NewDeck(types...)

	// Check if the remaining cards are initially equal to the total number of cards
	expectedRemainingCards := len(types) * CARD_COUNT
	if d.RemainingCards() != expectedRemainingCards {
		t.Errorf("Expected remaining cards to be %d, but got %d", expectedRemainingCards, d.RemainingCards())
	}

	// Deal some cards from the deck
	n := 5
	_, err := d.Deal(n)
	if err != nil {
		t.Errorf("Failed to deal cards: %v", err)
	}

	// Check if the remaining cards are updated correctly
	expectedRemainingCards -= n
	if d.RemainingCards() != expectedRemainingCards {
		t.Errorf("Expected remaining cards to be %d after dealing %d cards, but got %d", expectedRemainingCards, n, d.RemainingCards())
	}
}
func TestPickACard(t *testing.T) {
	types := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	d := NewDeck(types...)

	// Pick a card from the deck
	card, err := d.PickACard()
	if err != nil {
		t.Errorf("Failed to pick a card: %v", err)
	}

	// Check if the picked card is not nil
	if card == nil {
		t.Errorf("Expected a non-nil card, but got nil")
	}

	// Check if the picked card is removed from the deck
	if len(d.Cards) != len(types)*CARD_COUNT-1 {
		t.Errorf("Expected deck to have %d cards after picking a card, but got %d", len(types)*CARD_COUNT-1, len(d.Cards))
	}

	// Check if the picked card is not present in the deck
	for _, c := range d.Cards {
		if c == *card {
			t.Errorf("Picked card is still present in the deck")
		}
	}
}
func TestDeal(t *testing.T) {
	types := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	d := NewDeck(types...)

	// Deal some cards from the deck
	n := 5
	dealCards, err := d.Deal(n)
	if err != nil {
		t.Errorf("Failed to deal cards: %v", err)
	}

	// Check if the number of dealt cards is correct
	if len(dealCards) != n {
		t.Errorf("Expected to deal %d cards, but got %d", n, len(dealCards))
	}

	// Check if the dealt cards are removed from the deck
	expectedRemainingCards := len(types)*CARD_COUNT - n
	if d.RemainingCards() != expectedRemainingCards {
		t.Errorf("Expected remaining cards to be %d after dealing %d cards, but got %d", expectedRemainingCards, n, d.RemainingCards())
	}

	// Check if the dealt cards are not present in the deck
	for _, c := range dealCards {
		for _, card := range d.Cards {
			if c == &card {
				t.Errorf("Dealt card is still present in the deck")
			}
		}
	}
}
func TestIdentifyCard(t *testing.T) {
	types := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	d := NewDeck(types...)

	// Pick a card from the deck
	card, err := d.PickACard()
	if err != nil {
		t.Errorf("Failed to pick a card: %v", err)
	}

	// Identify the picked card
	identifiedCard := d.IdentifyCard(d.Title(card))

	// Check if the identified card has the correct suit
	if identifiedCard["suit"] != card.Suit {
		t.Errorf("Expected identified card suit to be %s, but got %s", card.Suit, identifiedCard["suit"])
	}

	// Check if the identified card has the correct ID
	if identifiedCard["id"] != card.ID {
		t.Errorf("Expected identified card ID to be %s, but got %s", strconv.Itoa(card.ID), identifiedCard["id"])
	}
}

func TestParseTitle(t *testing.T) {
	title := "Ace of Spades"
	expectedSuit := "Spades"
	expectedTitle := "Ace"

	title, suit := parseTitle(title)

	if suit != expectedSuit {
		t.Errorf("Expected suit to be %s, but got %s", expectedSuit, suit)
	}

	if title != expectedTitle {
		t.Errorf("Expected ID to be %s, but got %s", expectedTitle, title)
	}
}
