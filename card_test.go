package goker

import (
	"fmt"
	"testing"
)

// NOTE: Test Score()
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

	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	results := Score(p)
	score := results[0].Score
	fmt.Println("::::::: The following matches occured => ")
	for _, r := range results {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	if score < valuehandType[RoyalFlush ]{
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

	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	results := Score(p)
	score := results[0].Score
	fmt.Println("::::::: The following matches occured => ")
	for _, r := range results {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	if score < valuehandType[StraightFlush] || score > valuehandType[RoyalFlush ]{
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

	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	results := Score(p)
	score := results[0].Score
	fmt.Println("::::::: The following matches occured => ")
	for _, r := range results {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	if score < valuehandType[StraightFlush] || score > valuehandType[RoyalFlush ]{
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

	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	results := Score(p)
	score := results[0].Score
	fmt.Println("::::::: The following matches occured => ")
	for _, r := range results {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	if score < valuehandType[StraightFlush] || score > valuehandType[RoyalFlush ]{
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

	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	results := Score(p)
	score := results[0].Score
	fmt.Println("::::::: The following matches occured => ")
	for _, r := range results {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	if score < valuehandType[FourOfAKind] || score > valuehandType[StraightFlush ]{
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

	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	results := Score(p)
	score := results[0].Score
	fmt.Println("::::::: The following matches occured => ")
	for _, r := range results {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	if score < valuehandType[FourOfAKind] || score > valuehandType[StraightFlush ]{
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

	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	results := Score(p)
	score := results[0].Score
	fmt.Println("::::::: The following matches occured => ")
	for _, r := range results {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	if score < valuehandType[FullHouse] || score > valuehandType[FourOfAKind ]{
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

	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	results := Score(p)
	score := results[0].Score
	fmt.Println("::::::: The following matches occured => ")
	for _, r := range results {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	if score < valuehandType[FullHouse] || score > valuehandType[FourOfAKind ]{
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

	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	results := Score(p)
	score := results[0].Score
	fmt.Println("::::::: The following matches occured => ")
	for _, r := range results {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	if score < valuehandType[Flush] || score > valuehandType[FullHouse ]{
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

	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	results := Score(p)
	score := results[0].Score
	fmt.Println("::::::: The following matches occured => ")
	for _, r := range results {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	if score < valuehandType[Straight] || score > valuehandType[Flush ]{
		t.Error("Not landed in correct hand")
	}
}

func TestStraightWithAceAsOne(t *testing.T) {
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

	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	results := Score(p)
	score := results[0].Score
	fmt.Println("::::::: The following matches occured => ")
	for _, r := range results {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	if score < valuehandType[Straight] || score > valuehandType[Flush ]{
		t.Error("Not landed in correct hand")
	}
}

func TestStraightWithAceAsFourteen(t *testing.T) {
	hand := []Card{
		Card{
			Suit: Heart,
			Rank: Queen,
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
			Rank: Jack,
		},
		Card{
			Suit: Diamond,
			Rank: Ten,
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

	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	results := Score(p)
	score := results[0].Score
	fmt.Println("::::::: The following matches occured => ")
	for _, r := range results {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	if score < valuehandType[Straight] || score > valuehandType[Flush ]{
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

	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	results := Score(p)
	score := results[0].Score
	fmt.Println("::::::: The following matches occured => ")
	for _, r := range results {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	if score < valuehandType[ThreeOfAKind] || score > valuehandType[Straight ]{
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

	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	results := Score(p)
	score := results[0].Score
	fmt.Println("::::::: The following matches occured => ")
	for _, r := range results {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	if score < valuehandType[TwoPair] || score > valuehandType[ThreeOfAKind ]{
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

	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	results := Score(p)
	score := results[0].Score
	fmt.Println("::::::: The following matches occured => ")
	for _, r := range results {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	if score < valuehandType[OnePair] || score > valuehandType[TwoPair ]{
		t.Error("Not landed in correct hand")
	}
}

func TestHighCard(t *testing.T) {
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
			Rank: Eight,
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

	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	results := Score(p)
	score := results[0].Score
	fmt.Println("::::::: The following matches occured => ")
	for _, r := range results {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	if score < valuehandType[HighCard] || score > valuehandType[OnePair ]{
		t.Error("Not landed in correct hand")
	}
}

//Testing few dubious(not-so-random) scenarios
func TestScenario1(t *testing.T) {
	hand := []Card{
		Card{
			Suit: Diamond,
			Rank: Queen,
		},
		Card{
			Suit: Diamond,
			Rank: Queen,
		},
	}
	table := []Card{
		Card{
			Suit: Diamond,
			Rank: Four,
		},
		Card{
			Suit: Heart,
			Rank: Jack,
		},
		Card{
			Suit: Club,
			Rank: Jack,
		},
		Card{
			Suit: Diamond,
			Rank: Five,
		},
		Card{
			Suit: Heart,
			Rank: Six,
		},
	}
	p := Player{
		TwoCards:   hand,
		SevenCards: append(table, hand...),
	}

	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	results := Score(p)
	score := results[0].Score
	fmt.Println("::::::: The following matches occured => ")
	for _, r := range results {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	if score < valuehandType[HighCard] || score > valuehandType[RoyalFlush ]{
		t.Error("Not landed in correct hand")
	}
}

func TestScenario2(t *testing.T) {
	hand := []Card{
		Card{
			Suit: Heart,
			Rank: Seven,
		},
		Card{
			Suit: Spade,
			Rank: Two,
		},
	}
	table := []Card{
		Card{
			Suit: Diamond,
			Rank: Four,
		},
		Card{
			Suit: Heart,
			Rank: Jack,
		},
		Card{
			Suit: Club,
			Rank: Jack,
		},
		Card{
			Suit: Diamond,
			Rank: Five,
		},
		Card{
			Suit: Heart,
			Rank: Six,
		},
	}
	p := Player{
		TwoCards:   hand,
		SevenCards: append(table, hand...),
	}

	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	fmt.Println("===========Cards=================")
	PrintCards(p.SevenCards)
	fmt.Println("===========Cards=================")
	results := Score(p)
	score := results[0].Score
	fmt.Println("::::::: The following matches occured => ")
	for _, r := range results {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	if score < valuehandType[HighCard] || score > valuehandType[RoyalFlush]{
		t.Error("Not landed in correct hand")
	}
}

//NOTE: Test GetWinner()
func TestGetWinner1(t *testing.T) {
	hand1 := []Card{
		Card{
			Suit: Diamond,
			Rank: Queen,
		},
		Card{
			Suit: Diamond,
			Rank: Queen,
		},
	}
	fmt.Printf(">> p1 cards: ")
	PrintCards(hand1)
	hand2 := []Card{
		Card{
			Suit: Heart,
			Rank: Seven,
		},
		Card{
			Suit: Spade,
			Rank: Two,
		},
	}
	fmt.Printf(">> p2 cards: ")
	PrintCards(hand2)
	table := []Card{
		Card{
			Suit: Diamond,
			Rank: Four,
		},
		Card{
			Suit: Heart,
			Rank: Jack,
		},
		Card{
			Suit: Club,
			Rank: Jack,
		},
		Card{
			Suit: Diamond,
			Rank: Five,
		},
		Card{
			Suit: Heart,
			Rank: Six,
		},
	}
	fmt.Printf(">> table cards: ")
	PrintCards(table)
	
	var p1, p2 Player
	p1.TwoCards = hand1
	p2.TwoCards = hand2

	r1 := Score(p1)
	r2 := Score(p2)

	fmt.Println(":::::::Matches for p1 => ")
	for _, r := range r1 {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	
	fmt.Println(":::::::Matches for p2 => ")
	for _, r := range r2 {
		fmt.Printf(">> HandType: %v \t >> Score: %v \t >> Cards: ", r.HandType, r.Score)
		PrintCards(r.Cards)
	}
	fmt.Printf("\n\n\n")

	w, _, _ := GetWinner(p1, p2, table)
	if w != 1 {
		t.Errorf("Player %v should've won\n", w)
	}
}
