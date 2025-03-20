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

/*
 * Identifies the card by its suit and ID
 * returns id and suit as a JSON
 */
func (c *Card) Identify() map[string]interface{} {
	return map[string]interface{}{
		"id":   c.ID,
		"suit": c.Suit,
	}
}
