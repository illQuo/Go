package main

func main() {
	//cards := newDeck()
	// writing new deck of cards to file
	//cards.saveToFile("card_out.txt")
	deck := newDeckFromFile("card_out.txt")
	//deck.print()
	deck.shuffleDeck()
	deck.print()
	//deck.print()
	//cards.print()
	//cards.deal(8)
	//cards.print()
}
