package deck

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

const CARD_COUNT = 15

// Deck is a deck of cards with field types initialized and assigned to the types of cards
type Deck struct {
	cards     []Card
	faceCards []string
	types     []string
}

// NewDeck creates a new deck of cards
func NewDeck(types ...string) *Deck {
	// If no types are provided, use default types
	if len(types) == 0 {
		types = []string{
			"Spades", "Diamonds", "Clubs", "Hearts",
		}
	}

	deck := &Deck{
		cards: make([]Card, 0, len(types)*CARD_COUNT),
		faceCards: []string{
			"Jack", "Queen", "King", "Ace", "Joker",
		},
		types: types,
	}

	for _, suit := range deck.types {
		for j := 1; j <= CARD_COUNT; j++ {
			deck.cards = append(deck.cards, NewCard(j, suit))
		}
	}
	return deck
}

// PlayableDeck interface for decks that can be shuffled
type PlayableDeck interface {
	Shuffle()
}

// Title returns the title of the card, including its suit
func (d *Deck) Title(c *Card) string {
	if c.ID < 10 {
		return fmt.Sprintf("%d of %s", c.ID, c.Suit)
	}
	return fmt.Sprintf("%s of %s", d.faceCards[(c.ID-1)%len(d.faceCards)], c.Suit)
}

// Shuffle shuffles the deck using the Fisher-Yates algorithm
func (d *Deck) Shuffle() {
	fmt.Println("Shuffling Deck")
	// Implement Fisher-Yates shuffle
	for i := len(d.cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// RemainingCards returns the number of cards left in the deck
func (d *Deck) RemainingCards() int {
	return len(d.cards)
}

// PickACard picks a card from the deck and returns it as a string
func (d *Deck) PickACard() (card *Card, err error) {
	if len(d.cards) == 0 {
		return nil, errors.New("no cards left in the deck")
	}
	card = &d.cards[0]
	d.cards = d.cards[1:] // Remove the card from the deck
	return card, nil
}

// Helper Functions

// Parse string for title and ID by splitting string with " of "
func parseTitle(title string) (string, string) {
	parts := strings.Split(title, " of ")
	return parts[0], parts[1]
}

// IdentifyCard identifies the card by its suit and ID
func (d *Deck) IdentifyCard(title string) map[string]interface{} {
	cardTitle, suit := parseTitle(title)
	id := 0
	switch cardTitle {
	case "Jack":
		id = 11
	case "Queen":
		id = 12
	case "King":
		id = 13
	case "Ace":
		id = 14
	case "Joker":
		id = 15
	default:
		// convert string to int
		id, _ = strconv.Atoi(cardTitle)
	}
	return map[string]interface{}{
		"id":   id,
		"suit": suit,
	}
}
