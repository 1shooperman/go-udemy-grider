package main

func main() {
	cards := newDeck()

	// cards.print()
	// hand, remainingDeck := deal(cards, 5)

	// remainingDeck.print()
	// hand.print()

	// fmt.Println(cards.toString())

	// cards.saveToFile("my_cards")

	// card2 := newDeckFromFile("my_cards")

	// card2.print()

	cards.shuffle()
	cards.print()
}
