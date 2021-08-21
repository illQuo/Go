package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// init sets initial values for variables used in the package.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, fmt.Sprintf("%v of %v", value, suit))
		}
	}

	return cards

}

func (d deck) deal(cardsToDeal int) {
	randomCardsToDeal := []string{}

	// We want to choose a random card from the deck
	for i := 0; len(randomCardsToDeal) < cardsToDeal; i++ {
		cardExists := false
		randomCard := d[rand.Intn(len(d))]

		for j := range randomCardsToDeal {
			if randomCard == randomCardsToDeal[j] {
				cardExists = true
				break
			}
		}

		if cardExists != true {
			randomCardsToDeal = append(randomCardsToDeal, randomCard)

			// after we have 'dealt' a card, we remove it from our original deck
			for k, v := range d {
				if randomCard == v {
					// copy last element to index of k
					d[k] = d[len(d)-1]
					// wipe last element
					d[len(d)-1] = ""
					// truncate slice
					d = d[:len(d)-1]
					break
				}
			}
		}
	}

	fmt.Println(randomCardsToDeal)
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	deckString := strings.Join([]string(d), ",")
	return deckString
}

func (d deck) saveToFile(fileLocation string) error {
	deckString := d.toString()
	return ioutil.WriteFile(fileLocation, []byte(deckString), 0644)
}
