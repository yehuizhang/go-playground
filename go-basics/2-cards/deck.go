package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
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
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteArray, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return deck(strings.Split(string(byteArray), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(source)
	rand.Seed(510)
	for i := len(d); i > 0; i-- {
		newPosition := randGenerator.Intn(i)
		d[newPosition], d[i-1] = d[i-1], d[newPosition]
	}
}
