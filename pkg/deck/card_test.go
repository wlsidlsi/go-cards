// FILEPATH: /Users/chris/Documents/Go/1/pkg/deck/card_test.go

package deck

import "testing"

func TestNewCard(t *testing.T) {
	id := 1
	suit := "Spades"

	card := NewCard(id, suit)

	if card.ID != id {
		t.Errorf("Expected card ID to be %d, but got %d", id, card.ID)
	}

	if card.Suit != suit {
		t.Errorf("Expected card suit to be %s, but got %s", suit, card.Suit)
	}
}
