package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create type deck
type deck []string

var name string

func newDeck() deck {
	newDeck := []string{}
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four"}

	for _, value := range values {
		for _, suit := range suits {
			newDeck = append(newDeck, value+" of "+suit)
		}
	}

	return newDeck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// parm 0666 mean anyone can read and write
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	strSlices := strings.Split(string(bs), ",")
	return deck(strSlices)
}

func (d deck) shuffling() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for idx := range d {
		newIdx := r.Intn(len(d) - 1)
		d[idx], d[newIdx] = d[newIdx], d[idx]
	}
}
