package main

import "fmt"

func main() {
	//cards := newDeck()
	// writing new deck of cards to file
	//cards.saveToFile("card_out.txt")
	deck := newDeckFromFile("card_out.txt")
	fmt.Println(deck)
	//cards.print()
	//cards.deal(8)
	//cards.print()
}
