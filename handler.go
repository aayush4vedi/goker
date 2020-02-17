package goker

import (
	"math/rand"
	"time"
)

type Suit string

const (
	Spade   = "♠"
	Club    = "♣"
	Diamond = "♦"
	Heart   = "♥"
)

var suits = [...]Suit{Spade, Club, Diamond, Heart}

type Rank string

const (
	Ace   = "A"
	Two   = "2"
	Three = "3"
	Four  = "4"
	Five  = "5"
	Six   = "6"
	Seven = "7"
	Eight = "8"
	Nine  = "9"
	Ten   = "10"
	Jack  = "J"
	Queen = "Q"
	King  = "K"
)

var ranks = [...]Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

type Card struct {
	Rank
	Suit
}

type Deck []Card

type Hand []Card

func NewDeck() Deck {
	var deck []Card
	for _, s := range suits {
		for _, r := range ranks {
			deck = append(deck, Card{Suit: s, Rank: r})
		}
	}
	return shuffle(deck)
}

func shuffle(d Deck) Deck {
	newD := make(Deck, len(d))
	rand.Seed(time.Now().UnixNano())
	perm := rand.Perm(len(d))
	for i, v := range perm {
		newD[v] = d[i]
	}
	return newD
}

func GetHand(n int, deck Deck) (Hand, Deck){
	var hand Hand
	cards := deck[:n]
	hand = append(hand,cards...)
	return hand , shuffle(deck[n:])
}
