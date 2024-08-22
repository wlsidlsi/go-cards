package deck

// Card represents a single card in the deck
type Card struct {
	ID   int
	Suit string
}

// NewCard creates a new card with an ID and suit
func NewCard(id int, suit string) Card {
	return Card{ID: id, Suit: suit}
}
