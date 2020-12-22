package main

func main() {
	cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")

	cards.shuffling()
	cards.print()
	// fmt.Println(cards.toString())

	// hand, remainCards := deal(cards, 5)

	// cards.print()
	// hand.print()
	// remainCards.print()
}
