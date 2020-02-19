package main

import (
	"fmt"

	"github.com/goker"
)

const n, m int = 2, 5

var p1, p2 goker.Player
var table goker.Hand

func main() {
	deck := goker.NewDeck()

	//draw cards
	p1.TwoCards, deck = goker.GetHand(n, deck)
	p2.TwoCards, deck = goker.GetHand(n, deck)
	table, deck = goker.GetHand(m, deck)

	fmt.Printf(">> p1 cards: ")
	goker.PrintCards(p1.TwoCards)
	fmt.Printf(">> p2 cards: ")
	goker.PrintCards(p2.TwoCards)
	fmt.Printf(">> table cards: ")
	goker.PrintCards(table)

	winner, handType, hand := goker.GetWinner(p1, p2, table)
	fmt.Printf("Player %v won due to higher hand: %v - ", winner, handType)
	goker.PrintCards(hand)
}
