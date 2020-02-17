package main

import (
	"fmt"

	"github.com/goker"
)

const n, m int = 3, 2
var p1, p2, table goker.Hand

func main() {
	deck := goker.NewDeck()

	//draw cards
	p1, deck = goker.GetHand(n, deck)
	p2, deck = goker.GetHand(n, deck)
	table, deck = goker.GetHand(m, deck)
}
