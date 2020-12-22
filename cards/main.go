package main

func main() {
	cards := newDeck()

	hand, remainCards := deal(cards, 5)

	cards.print()
	hand.print()
	remainCards.print()
}
