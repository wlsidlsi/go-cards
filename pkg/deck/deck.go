package deck

import (
	"compress/gzip"
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const CARD_COUNT = 15

// Deck is a deck of cards with field types initialized and assigned to the types of cards
type Deck struct {
	Cards     []Card
	FaceCards []string
	Types     []string
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
		Cards: make([]Card, 0, len(types)*CARD_COUNT),
		FaceCards: []string{
			"Jack", "Queen", "King", "Ace", "Joker",
		},
		Types: types,
	}

	for _, suit := range deck.Types {
		for j := 1; j <= CARD_COUNT; j++ {
			deck.Cards = append(deck.Cards, NewCard(j, suit))
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
	return fmt.Sprintf("%s of %s", d.FaceCards[(c.ID-1)%len(d.FaceCards)], c.Suit)
}

// Shuffle shuffles the deck using the Fisher-Yates algorithm
func (d *Deck) Shuffle() {
	fmt.Println("Shuffling Deck")
	// Implement Fisher-Yates shuffle
	for i := len(d.Cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

// RemainingCards returns the number of cards left in the deck
func (d *Deck) RemainingCards() int {
	return len(d.Cards)
}

// PickACard picks a card from the deck and returns it as a string
func (d *Deck) PickACard() (card *Card, err error) {
	if len(d.Cards) == 0 {
		return nil, errors.New("no cards left in the deck")
	}
	card = &d.Cards[0]
	d.Cards = d.Cards[1:] // Remove the card from the deck
	return card, nil
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

// Save deck state by writing the Deck object to a file
func (d *Deck) Save() (string, error) {
	filePath := "./deck.bin" // Replace with the actual file path

	// Open the file for writing
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create an encoder
	encoder := gob.NewEncoder(file)

	// Encode and write the deck object to the file
	err = encoder.Encode(d)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

/* Save deck state.
 * Write the deck object to a file using compress/gzip package
 * return the file path and an error */
func (d *Deck) SaveGzip() (string, error) {
	filePath := "./deck.bin.gz" // Replace with the actual file path

	// Open the file for writing
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create a gzip writer
	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	// Create an encoder
	encoder := gob.NewEncoder(gzipWriter)
	// Encode and write the deck object to the file
	err = encoder.Encode(d)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

// Load deck state by reading the Deck object from a file
// return a deck object and an error after decoding
// with gob.NewDecoder(file)
func Load(filePath string) (*Deck, error) {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	// Decode and read the deck object from the file
	var deck Deck
	err = decoder.Decode(&deck)
	if err != nil {
		return nil, err
	}
	return &deck, nil
}

/* Load deck state.
 * Read the deck object to a file using compress/gzip package
 * return the deck object and an error */
func LoadGzip(filePath string) (*Deck, error) {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a gzip reader
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	// Create a decoder
	decoder := gob.NewDecoder(gzipReader)
	// Decode and read the deck object from the file
	var deck Deck
	err = decoder.Decode(&deck)
	if err != nil {
		return nil, err
	}
	return &deck, nil
}

// Helper Functions

// Parse string for title and ID by splitting string with " of "
func parseTitle(title string) (string, string) {
	parts := strings.Split(title, " of ")
	return parts[0], parts[1]
}
