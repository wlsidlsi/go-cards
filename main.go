package main

import (
	"fmt"

	"example.com/pkg/deck"
)

func War() {
	// creates a new deck
	deck := deck.NewDeck()

	// shuffles the deck
	deck.Shuffle()
	fmt.Println()

	fmt.Println("Warrrrrrrr!!!!!!")

	fmt.Println()

	// picks a card from the deck
	// prints "Your Card is {Card}"
	card, err := deck.PickACard()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Your Card is", deck.Title((card)))
	}
	fmt.Println()

	// shuffles the deck
	deck.Shuffle()
	fmt.Println()

	// picks a card from the deck
	// prints "Your Card is {Card}"
	card, err = deck.PickACard()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("My Card is", deck.Title(card))
	}

	// iterates over teh remaining cards in the deck
	// prints "Remaining Cards: {Number}"
	for deck.RemainingCards() > 0 {
		card, err = deck.PickACard()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("My Card", deck.Title(card), deck.IdentifyCard(deck.Title(card)))
		}
	}
}

func main() {
	// creates a new deck
	deck := deck.NewDeck()

	// shuffles the deck
	deck.Shuffle()
	fmt.Println()

	fmt.Println("Warrrrrrrr!!!!!!")

	fmt.Println()

	// picks a card from the deck
	// prints "Your Card is {Card}"
	card, err := deck.PickACard()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Your Card is", deck.Title((card)))
	}
	fmt.Println()

	// shuffles the deck
	deck.Shuffle()
	fmt.Println()

	// picks a card from the deck
	// prints "Your Card is {Card}"
	card, err = deck.PickACard()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("My Card is", deck.Title(card))
	}

	// iterates over teh remaining cards in the deck
	// prints "Remaining Cards: {Number}"
	for deck.RemainingCards() > 0 {
		card, err = deck.PickACard()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("My Card", deck.Title(card), deck.IdentifyCard(deck.Title(card)))
		}
	}
}

// Disable cgo checks for pointer size
