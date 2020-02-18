package main

import (
	"fmt"

	"github.com/goker"
)

const n, m int = 5, 2

var p1, p2 goker.Player
var table goker.Hand

func main() {
	deck := goker.NewDeck()

	//draw cards
	p1.FiveCards, deck = goker.GetHand(n, deck)
	p2.FiveCards, deck = goker.GetHand(n, deck)
	table, deck = goker.GetHand(m, deck)

	winner := goker.GetWinner(p1, p2, table)
	fmt.Printf("Player %v won!\n", winner)
}
