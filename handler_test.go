package goker

import (
	"fmt"
	"testing"
)

func TestRoyalFlush(t *testing.T) {
	hand := []Card{
		Card{
			Suit: Diamond,
			Rank: Ace,
		},
		Card{
			Suit: Diamond,
			Rank: Ten,
		},
	}
	table := []Card{
		Card{
			Suit: Heart,
			Rank: Two,
		},
		Card{
			Suit: Diamond,
			Rank: Jack,
		},
		Card{
			Suit: Diamond,
			Rank: King,
		},
		Card{
			Suit: Diamond,
			Rank: Queen,
		},
		Card{
			Suit: Heart,
			Rank: Three,
		},
	}
	p := Player{
		TwoCards:   hand,
		SevenCards: append(table, hand...),
	}

	fmt.Printf(">>>> Score: %v\n", Score(p))
	if Score(p) < RoyalFlush {
		t.Error("Not landed in correct hand")
	}
}

func TestStraightFlushLow(t *testing.T) {
	hand := []Card{
		Card{
			Suit: Diamond,
			Rank: Five,
		},
		Card{
			Suit: Diamond,
			Rank: Three,
		},
	}
	table := []Card{
		Card{
			Suit: Heart,
			Rank: Two,
		},
		Card{
			Suit: Diamond,
			Rank: Two,
		},
		Card{
			Suit: Diamond,
			Rank: Four,
		},
		Card{
			Suit: Diamond,
			Rank: Six,
		},
		Card{
			Suit: Heart,
			Rank: Three,
		},
	}
	p := Player{
		TwoCards:   hand,
		SevenCards: append(table, hand...),
	}

	fmt.Printf(">>>> Score: %v\n", Score(p))
	if Score(p) < StraightFlush || Score(p) > RoyalFlush {
		t.Error("Not landed in correct hand")
	}
}

func TestStraightFlushHigh(t *testing.T) {
	hand := []Card{
		Card{
			Suit: Diamond,
			Rank: Nine,
		},
		Card{
			Suit: Diamond,
			Rank: Jack,
		},
	}
	table := []Card{
		Card{
			Suit: Heart,
			Rank: Two,
		},
		Card{
			Suit: Diamond,
			Rank: King,
		},
		Card{
			Suit: Diamond,
			Rank: Ten,
		},
		Card{
			Suit: Diamond,
			Rank: Queen,
		},
		Card{
			Suit: Heart,
			Rank: Three,
		},
	}
	p := Player{
		TwoCards:   hand,
		SevenCards: append(table, hand...),
	}

	fmt.Printf(">>>> Score: %v\n", Score(p))
	if Score(p) < StraightFlush || Score(p) > RoyalFlush {
		t.Error("Not landed in correct hand")
	}
}

func TestStraightFlushWithAce(t *testing.T) {
	hand := []Card{
		Card{
			Suit: Diamond,
			Rank: Ace,
		},
		Card{
			Suit: Diamond,
			Rank: Two,
		},
	}
	table := []Card{
		Card{
			Suit: Heart,
			Rank: Two,
		},
		Card{
			Suit: Diamond,
			Rank: Three,
		},
		Card{
			Suit: Diamond,
			Rank: Four,
		},
		Card{
			Suit: Diamond,
			Rank: Five,
		},
		Card{
			Suit: Heart,
			Rank: Three,
		},
	}
	p := Player{
		TwoCards:   hand,
		SevenCards: append(table, hand...),
	}

	fmt.Printf(">>>> Score: %v\n", Score(p))
	if Score(p) < StraightFlush || Score(p) > RoyalFlush {
		t.Error("Not landed in correct hand")
	}
}

func TestFourOfAKindHigh(t *testing.T) {
	hand := []Card{
		Card{
			Suit: Diamond,
			Rank: Ace,
		},
		Card{
			Suit: Diamond,
			Rank: Two,
		},
	}
	table := []Card{
		Card{
			Suit: Heart,
			Rank: Ace,
		},
		Card{
			Suit: Diamond,
			Rank: King,
		},
		Card{
			Suit: Club,
			Rank: Ace,
		},
		Card{
			Suit: Diamond,
			Rank: Queen,
		},
		Card{
			Suit: Spade,
			Rank: Ace,
		},
	}
	p := Player{
		TwoCards:   hand,
		SevenCards: append(table, hand...),
	}

	fmt.Printf(">>>> Score: %v\n", Score(p))
	if Score(p) < FourOfAKind || Score(p) > StraightFlush {
		t.Error("Not landed in correct hand")
	}
}

func TestFourOfAKindLow(t *testing.T) {
	hand := []Card{
		Card{
			Suit: Heart,
			Rank: Ace,
		},
		Card{
			Suit: Diamond,
			Rank: Two,
		},
	}
	table := []Card{
		Card{
			Suit: Heart,
			Rank: Two,
		},
		Card{
			Suit: Diamond,
			Rank: King,
		},
		Card{
			Suit: Club,
			Rank: Two,
		},
		Card{
			Suit: Diamond,
			Rank: Queen,
		},
		Card{
			Suit: Spade,
			Rank: Two,
		},
	}
	p := Player{
		TwoCards:   hand,
		SevenCards: append(table, hand...),
	}

	fmt.Printf(">>>> Score: %v\n", Score(p))
	if Score(p) < FourOfAKind || Score(p) > StraightFlush {
		t.Error("Not landed in correct hand")
	}
}

func TestFullHouseLow(t *testing.T) {
	hand := []Card{
		Card{
			Suit: Heart,
			Rank: Two,
		},
		Card{
			Suit: Diamond,
			Rank: Four,
		},
	}
	table := []Card{
		Card{
			Suit: Heart,
			Rank: Three,
		},
		Card{
			Suit: Diamond,
			Rank: Three,
		},
		Card{
			Suit: Club,
			Rank: Three,
		},
		Card{
			Suit: Diamond,
			Rank: Queen,
		},
		Card{
			Suit: Spade,
			Rank: Two,
		},
	}
	p := Player{
		TwoCards:   hand,
		SevenCards: append(table, hand...),
	}

	fmt.Printf(">>>> Score: %v\n", Score(p))
	if Score(p) < FullHouse || Score(p) > FourOfAKind {
		t.Error("Not landed in correct hand")
	}
}

func TestFullHouseHigh(t *testing.T) {
	hand := []Card{
		Card{
			Suit: Heart,
			Rank: King,
		},
		Card{
			Suit: Club,
			Rank: Ace,
		},
	}
	table := []Card{
		Card{
			Suit: Heart,
			Rank: Ace,
		},
		Card{
			Suit: Diamond,
			Rank: Ace,
		},
		Card{
			Suit: Diamond,
			Rank: Four,
		},
		Card{
			Suit: Diamond,
			Rank: Queen,
		},
		Card{
			Suit: Spade,
			Rank: King,
		},
	}
	p := Player{
		TwoCards:   hand,
		SevenCards: append(table, hand...),
	}

	fmt.Printf(">>>> Score: %v\n", Score(p))
	if Score(p) < FullHouse || Score(p) > FourOfAKind {
		t.Error("Not landed in correct hand")
	}
}

func TestFlush(t *testing.T) {
	hand := []Card{
		Card{
			Suit: Heart,
			Rank: Ace,
		},
		Card{
			Suit: Heart,
			Rank: King,
		},
	}
	table := []Card{
		Card{
			Suit: Heart,
			Rank: Jack,
		},
		Card{
			Suit: Heart,
			Rank: Ten,
		},
		Card{
			Suit: Heart,
			Rank: Seven,
		},
		Card{
			Suit: Diamond,
			Rank: Queen,
		},
		Card{
			Suit: Spade,
			Rank: King,
		},
	}
	p := Player{
		TwoCards:   hand,
		SevenCards: append(table, hand...),
	}

	fmt.Printf(">>>> Score: %v\n", Score(p))
	if Score(p) < Flush || Score(p) > FullHouse {
		t.Error("Not landed in correct hand")
	}
}

func TestStraight(t *testing.T) {
	hand := []Card{
		Card{
			Suit: Heart,
			Rank: Two,
		},
		Card{
			Suit: Spade,
			Rank: Three,
		},
	}
	table := []Card{
		Card{
			Suit: Heart,
			Rank: Six,
		},
		Card{
			Suit: Heart,
			Rank: Three,
		},
		Card{
			Suit: Heart,
			Rank: Five,
		},
		Card{
			Suit: Diamond,
			Rank: Two,
		},
		Card{
			Suit: Spade,
			Rank: Four,
		},
	}
	p := Player{
		TwoCards:   hand,
		SevenCards: append(table, hand...),
	}

	fmt.Printf(">>>> Score: %v\n", Score(p))
	if Score(p) < Straight || Score(p) > Flush {
		t.Error("Not landed in correct hand")
	}
}

func TestStraightWithAce(t *testing.T) {
	hand := []Card{
		Card{
			Suit: Heart,
			Rank: Two,
		},
		Card{
			Suit: Spade,
			Rank: Three,
		},
	}
	table := []Card{
		Card{
			Suit: Heart,
			Rank: King,
		},
		Card{
			Suit: Heart,
			Rank: Four,
		},
		Card{
			Suit: Heart,
			Rank: Five,
		},
		Card{
			Suit: Diamond,
			Rank: Two,
		},
		Card{
			Suit: Spade,
			Rank: Ace,
		},
	}
	p := Player{
		TwoCards:   hand,
		SevenCards: append(table, hand...),
	}

	fmt.Printf(">>>> Score: %v\n", Score(p))
	if Score(p) < Straight || Score(p) > Flush {
		t.Error("Not landed in correct hand")
	}
}

func TestThreeOfAKind(t *testing.T) {
	hand := []Card{
		Card{
			Suit: Heart,
			Rank: Ace,
		},
		Card{
			Suit: Spade,
			Rank: Three,
		},
	}
	table := []Card{
		Card{
			Suit: Heart,
			Rank: Six,
		},
		Card{
			Suit: Heart,
			Rank: King,
		},
		Card{
			Suit: Heart,
			Rank: Five,
		},
		Card{
			Suit: Diamond,
			Rank: Ace,
		},
		Card{
			Suit: Spade,
			Rank: Ace,
		},
	}
	p := Player{
		TwoCards:   hand,
		SevenCards: append(table, hand...),
	}

	fmt.Printf(">>>> Score: %v\n", Score(p))
	if Score(p) < ThreeOfAKind || Score(p) > Straight {
		t.Error("Not landed in correct hand")
	}
}

func TestTwoPair(t *testing.T) {
	hand := []Card{
		Card{
			Suit: Heart,
			Rank: Ace,
		},
		Card{
			Suit: Spade,
			Rank: King,
		},
	}
	table := []Card{
		Card{
			Suit: Heart,
			Rank: Six,
		},
		Card{
			Suit: Heart,
			Rank: King,
		},
		Card{
			Suit: Heart,
			Rank: Five,
		},
		Card{
			Suit: Diamond,
			Rank: Ace,
		},
		Card{
			Suit: Spade,
			Rank: Jack,
		},
	}
	p := Player{
		TwoCards:   hand,
		SevenCards: append(table, hand...),
	}

	fmt.Printf(">>>> Score: %v\n", Score(p))
	if Score(p) < TwoPair || Score(p) > ThreeOfAKind {
		t.Error("Not landed in correct hand")
	}
}

func TestOnePair(t *testing.T) {
	hand := []Card{
		Card{
			Suit: Heart,
			Rank: Ace,
		},
		Card{
			Suit: Spade,
			Rank: King,
		},
	}
	table := []Card{
		Card{
			Suit: Heart,
			Rank: Six,
		},
		Card{
			Suit: Heart,
			Rank: Jack,
		},
		Card{
			Suit: Heart,
			Rank: Five,
		},
		Card{
			Suit: Diamond,
			Rank: Ace,
		},
		Card{
			Suit: Spade,
			Rank: Nine,
		},
	}
	p := Player{
		TwoCards:   hand,
		SevenCards: append(table, hand...),
	}

	fmt.Printf(">>>> Score: %v\n", Score(p))
	if Score(p) < OnePair || Score(p) > TwoPair {
		t.Error("Not landed in correct hand")
	}
}
