package main

func main() {
	cards := newDeck()
	// writing new deck of cards to file
	cards.saveToFile("card_out.txt")
	cards.print()
	cards.deal(8)
	cards.print()
}
