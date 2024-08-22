package main

import "example.com/pkg/deck"

func main() {

	myDeck := deck.NewDeck()
	myDeck.Shuffle()
	myCards, err := myDeck.Deal(60)
	if err != nil {
		panic(err)
	}
	for _, card := range myCards {
		println(myDeck.Title(card))
	}
	// // creates a new deck
	// deck1, err := deck.LoadGzip("./deck.bin.gz")
	// if err != nil {
	// 	fmt.Println("Error Loadig:", err)
	// 	deck1 = deck.NewDeck()
	// } else {
	// 	fmt.Println(deck1.RemainingCards())
	// }

	// // shuffles the deck
	// deck1.Shuffle()
	// fmt.Println()

	// fmt.Println("Warrrrrrrr!!!!!!")

	// fmt.Println()

	// // picks a card from the deck
	// // prints "Your Card is {Card}"
	// card, err := deck1.PickACard()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Your Card is", deck1.Title((card)))
	// }
	// fmt.Println()

	// // shuffles the deck
	// deck1.Shuffle()
	// fmt.Println()

	// // picks a card from the deck
	// // prints "Your Card is {Card}"
	// card, err = deck1.PickACard()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("My Card is", deck1.Title(card))
	// }

	// _, err = deck1.SaveGzip()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // iterates over teh remaining cards in the deck
	// // prints "Remaining Cards: {Number}"
	// for deck1.RemainingCards() > 0 {
	// 	card, err = deck1.PickACard()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println("My Card", deck1.Title(card), deck1.IdentifyCard(deck1.Title(card)))
	// 	}
	// }
}

// Disable cgo checks for pointer size
