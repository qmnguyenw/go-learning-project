package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
type deck []string

func initDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// _ can replace index if do not output index
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) printDeck() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func dealDeck(d deck, handSize int) (deck, deck) {
	// fmt.Println(len(d))
	// return nil, nil
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveDeck(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func loadDeck(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()

		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffleDeck() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
