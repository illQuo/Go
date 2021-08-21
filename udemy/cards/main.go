package main

func main() {
	cards := newDeck()
	// should take user input here
	cards.print()
	cards.deal(8)
	cards.print()
}
