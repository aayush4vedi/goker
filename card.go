package goker

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"time"
)

type Suit uint8

const (
	_ Suit = iota
	Spade
	Club
	Diamond
	Heart
)

const (
	minSuit = Spade
	maxSuit = Heart
)

var suits = [...]Suit{Spade, Club, Diamond, Heart}

type Rank uint8

const (
	_ Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

const (
	minRank = Two
	maxRank = Ace
)

var ranks = [...]Rank{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

type HandType float64

const (
	_ HandType = iota
	RoyalFlush
	StraightFlush
	FourOfAKind
	FullHouse
	Flush
	Straight
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

type Card struct {
	Rank
	Suit
}

type Deck []Card

type Hand []Card

type Result struct {
	HandType string
	Score    float64
	Cards    []Card
}

var symbolRank = map[Rank]string{
	Two:   "2",
	Three: "3",
	Four:  "4",
	Five:  "5",
	Six:   "6",
	Seven: "7",
	Eight: "8",
	Nine:  "9",
	Ten:   "10",
	Jack:  "J",
	Queen: "Q",
	King:  "K",
	Ace:   "A",
}
var symbolSuit = map[Suit]string{
	Spade:   "♠",
	Club:    "♣",
	Diamond: "♦",
	Heart:   "♥",
}

var symbolhandType = map[HandType]string{
	RoyalFlush:    "RoyalFlush",
	StraightFlush: "StraightFlush",
	FourOfAKind:   "FourOfAKind",
	FullHouse:     "FullHouse",
	Flush:         "Flush",
	Straight:      "Straight",
	ThreeOfAKind:  "ThreeOfAKind",
	TwoPair:       "TwoPair",
	OnePair:       "OnePair",
	HighCard:      "HighCard",
}

var valuehandType = map[HandType]float64{
	RoyalFlush:    900,
	StraightFlush: 800,
	FourOfAKind:   700,
	FullHouse:     600,
	Flush:         500,
	Straight:      400,
	ThreeOfAKind:  300,
	TwoPair:       200,
	OnePair:       100,
	HighCard:      0,
}

const n, m int = 2, 5 //Number of Cards for player & table respectively

func PrintCards(cards []Card) {
	fmt.Printf("[")
	for _, card := range cards {
		fmt.Printf("{ %v%v }", symbolRank[card.Rank], symbolSuit[card.Suit])
	}
	fmt.Printf("]\n")
}

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

func GetHand(n int, deck Deck) (Hand, Deck) {
	var hand Hand
	cards := deck[:n]
	hand = append(hand, cards...)
	return hand, shuffle(deck[n:])
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	}
}

func areConsecutive(cards []Card) bool {
	for i := 0; i < len(cards)-1; i++ {
		if cards[i+1].Rank-cards[i].Rank != 1 {
			return false
		}
	}
	return true
}

func intersection(a, b []Card) []Card { //NOTE: pass the smaller slice in `a`
	m := make(map[Card]bool)
	c := []Card(nil)
	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; !ok {
			c = append(c, item)
		}
	}
	return c
}

func nOfSameSuit(h Hand, n int) ([]Card, bool) {
	m := make(map[Suit][]Card)
	for i := len(h) - 1; i >= 0; i-- {
		m[h[i].Suit] = append(m[h[i].Suit], h[i])
		if len(m[h[i].Suit]) == n {
			sort.Slice(m[h[i].Suit], Less(m[h[i].Suit]))
			return m[h[i].Suit], true
		}
	}
	return nil, false
}

func nOfSameRank(h Hand, n int) ([]Card, bool) {
	m := make(map[Rank][]Card)
	for i := len(h) - 1; i >= 0; i-- {
		m[h[i].Rank] = append(m[h[i].Rank], h[i])
		if len(m[h[i].Rank]) == n {
			sort.Slice(m[h[i].Rank], Less(m[h[i].Rank]))
			return m[h[i].Rank], true
		}
	}
	return nil, false
}
func nPair(h Hand, n int) ([]Card, bool) {
	ret := []Card(nil)
	cnt := 0
	for i := len(h) - 1; i >= 1; i-- {
		if h[i].Rank == h[i-1].Rank {
			cnt++
			ret = append(ret, h[i], h[i-1])
			i--
			if cnt == n {
				return ret, true
			}
		} else {
			cnt = 0
			ret = nil
		}
	}
	return nil, false
}

func normalizedScore(cards []Card, n int) float64 {
	cardScore, normalizer := float64(0), float64(0)
	for i, c := range cards {
		cardScore += (math.Pow(13, float64(i)) * (float64(c.Rank)))
		normalizer += math.Pow(13, float64(i)) * 14
	}
	return float64((cardScore * float64(n)) / normalizer)
}

func (h Hand) contains(card Card) bool {
	for _, c := range h {
		if (c.Suit == card.Suit) && (c.Rank == card.Rank) {
			return true
		}
	}
	return false
}

func nonRepeatingCards(h Hand) Hand {
	var ret Hand
	m := make(map[Rank]bool)
	for _, card := range h {
		if _, ok := m[card.Rank]; !ok {
			ret = append(ret, card)
			m[card.Rank] = true
		}
	}
	return ret
}

func checkRoyalFlushForSuit(h Hand, s Suit) ([]Card, bool) {
	if h.contains(Card{Suit: s, Rank: Ten}) && h.contains(Card{Suit: s, Rank: Jack}) && h.contains(Card{Suit: s, Rank: Queen}) && h.contains(Card{Suit: s, Rank: King}) && h.contains(Card{Suit: s, Rank: Ace}) {
		return []Card{
			Card{Suit: s, Rank: Ten},
			Card{Suit: s, Rank: Jack},
			Card{Suit: s, Rank: Queen},
			Card{Suit: s, Rank: King},
			Card{Suit: s, Rank: Ace},
		}, true
	}
	return []Card(nil), false
}

//900-RoyalFlush
func isRoyalFlush(h Hand) ([]Card, bool) {
	for suit := minSuit; suit < maxSuit; suit++ {
		if cards, ok := checkRoyalFlushForSuit(h, suit); ok {
			return cards, true
		}
	}
	return nil, false
}

//800-StraightFlush
func isStraightFlush(h Hand) ([]Card, bool) {
	if cards, ok := nOfSameSuit(h, 5); ok {
		if areConsecutive(cards) {
			return cards, true
		} else if cards[len(cards)-1].Rank == Ace && cards[0].Rank == Two && areConsecutive(cards[0:4]) {
			return cards, true
		}
	}
	return nil, false
}

//700-FourOfAKind
func isFourOfAKind(h Hand) ([]Card, bool) {
	return nOfSameRank(h, 4)
}

//600-FullHouse
func isFullHouse(h Hand) ([]Card, []Card, bool) {
	cards3, ok3 := nOfSameRank(h, 3)
	h = intersection(cards3, h)
	cards2, ok2 := nOfSameRank(h, 2)
	if ok2 && ok3 {
		return cards3, cards2, true
	}
	return nil, nil, false
}

//500-Flush
func isFlush(h Hand) ([]Card, bool) {
	return nOfSameSuit(h, 5)
}

//400-Straight
func isStraight(h Hand) ([]Card, bool) {
	//first form slice of non-repeating cards
	h = nonRepeatingCards(h)
	//check size >= 5; if yes take the last 5 elements
	if len(h) < 5 {
		return nil, false
	}
	//check for Ace-case
	if h[len(h)-1].Rank == Ace && h[0].Rank == Two && areConsecutive((h[:4])) {
		return append([]Card{h[len(h)-1]}, h[0:4]...), true
	}
	//return the max possible hand
	h = h[len(h)-5:]
	//check if areConsecutive
	if areConsecutive(h) {
		return h, true
	}
	return nil, false
}

//300-ThreeOfAKind
func isThreeOfAKind(h Hand) ([]Card, bool) {
	return nOfSameRank(h, 3)
}

//200-TwoPair
func isTwoPair(h Hand) ([]Card, []Card, bool) {
	cards1, ok1 := nPair(h, 1)
	h = intersection(cards1, h)
	cards2, ok2 := nPair(h, 1)
	if ok1 && ok2 {
		return cards1, cards2, true
	}
	return nil, nil, false
}

//100-OnePair
func isOnePair(h Hand) ([]Card, bool) {
	if cards, ok := nPair(h, 1); ok {
		return cards, true
	}
	return nil, false
}

func Score(p Player) []Result {
	sort.Slice(p.PocketCards, Less(p.PocketCards))
	sort.Slice(p.CommunityCards, Less(p.CommunityCards))

	ans := float64(0)
	result := []Result(nil)
	if cards, ok := isRoyalFlush(p.CommunityCards); ok {
		ans = valuehandType[RoyalFlush]
		result = append(result, Result{
			HandType: symbolhandType[RoyalFlush],
			Score:    ans,
			Cards:    cards,
		})
	}
	if cards, ok := isStraightFlush(p.CommunityCards); ok {
		ans = valuehandType[StraightFlush] + normalizedScore(cards, 90)
		result = append(result, Result{
			HandType: symbolhandType[StraightFlush],
			Score:    ans,
			Cards:    cards,
		})
	}
	if cards, ok := isFourOfAKind(p.CommunityCards); ok {
		ans = valuehandType[FourOfAKind] + normalizedScore(cards, 90)
		result = append(result, Result{
			HandType: symbolhandType[FourOfAKind],
			Score:    ans,
			Cards:    cards,
		})
	}
	if cards3, cards2, ok := isFullHouse(p.CommunityCards); ok {
		ans = valuehandType[FullHouse] + normalizedScore(cards3, 60) + normalizedScore(cards2, 30)
		handCards := []Card(nil)
		handCards = append(handCards, cards3...)
		handCards = append(handCards, cards2...)
		result = append(result, Result{
			HandType: symbolhandType[FullHouse],
			Score:    ans,
			Cards:    handCards,
		})
	}
	if cards, ok := isFlush(p.CommunityCards); ok {
		ans = valuehandType[Flush] + normalizedScore(cards, 90)
		result = append(result, Result{
			HandType: symbolhandType[Flush],
			Score:    ans,
			Cards:    cards,
		})
	}
	if cards, ok := isStraight(p.CommunityCards); ok {
		ans = valuehandType[Straight] + normalizedScore(cards, 90)
		result = append(result, Result{
			HandType: symbolhandType[Straight],
			Score:    ans,
			Cards:    cards,
		})
	}
	if cards, ok := isThreeOfAKind(p.CommunityCards); ok {
		ans = valuehandType[ThreeOfAKind] + normalizedScore(cards, 90)
		result = append(result, Result{
			HandType: symbolhandType[ThreeOfAKind],
			Score:    ans,
			Cards:    cards,
		})
	}
	if cards1, cards2, ok := isTwoPair(p.CommunityCards); ok {
		ans = valuehandType[TwoPair] + normalizedScore(cards1, 45) + normalizedScore(cards2, 45)
		handCards := []Card(nil)
		handCards = append(handCards, cards1...)
		handCards = append(handCards, cards2...)
		result = append(result, Result{
			HandType: symbolhandType[TwoPair],
			Score:    ans,
			Cards:    handCards,
		})
	}
	if cards, ok := isOnePair(p.CommunityCards); ok {
		ans = valuehandType[OnePair] + normalizedScore(cards, 90)
		result = append(result, Result{
			HandType: symbolhandType[OnePair],
			Score:    ans,
			Cards:    cards,
		})
	}
	// ans += normalizedScore(p.PocketCards, 10)
	ans = normalizedScore(p.PocketCards, 10) //Change of approach, might be reverted
	result = append(result, Result{
		HandType: symbolhandType[HighCard],
		Score:    ans,
		Cards:    []Card{p.PocketCards[1]},
	})
	result = append(result, Result{
		HandType: symbolhandType[HighCard],
		Score:    ans,
		Cards:    []Card{p.PocketCards[0]},
	})
	return result
}

func GetWinner(p1, p2 Player, table Hand) (int, string, []Card) {
	p1.CommunityCards = append(p1.PocketCards, table...)
	p2.CommunityCards = append(p2.PocketCards, table...)

	r1 := Score(p1)
	r2 := Score(p2)

	for r1[0].Score == r2[0].Score && len(r1) > 0 && len(r2) > 0 {
		r1 = r1[1:]
		r2 = r2[1:]
	}

	if r1[0].Score > r2[0].Score {
		return 1, r1[0].HandType, r1[0].Cards
	} else {
		return 2, r2[0].HandType, r2[0].Cards
	}
}

func containsAction(actions []Action, x Action) (int, bool) {
	for i, a := range actions {
		if x == a {
			return i, true
		}
	}
	return 0, false
}

func (game *Game) ActionHandler(action Action, val *Money) {
	switch action {
	case PutBlind:
		game.MainPot.value += BLIND * Money(len(game.Players))
		for i := range game.Players {
			game.Players[i].chips -= BLIND
		}
	case SitOut:
		//NOTE: assuming p1 by default for now
		fmt.Printf("=>> Sittin you out in this game. You have %v chips remaining\n", game.Players[0].chips)
		os.Exit(0)
	case Call:
		if game.isRaised {
			//NOTE: assuming p1 by default for now
			for i := 1; i < len(game.Players); i++ {
				if game.Players[i].chips >= *val {
					game.Players[i].chips -= *val
					game.MainPot.value += *val
				} else {
					game.Players[i].chips = 0
					game.MainPot.value += game.Players[i].chips
				}
			}
			game.isRaised = false
		} else {
			fmt.Printf("There is no Raise by anyone.Wrong move. Retry input.\n")
			game.waitUserInput(Same)
		}
	case Check:
		if game.isRaised {
			fmt.Printf("There is a Raise in game.Can't Check. Retry input.\n")
			game.waitUserInput(Same)
		}
	case Raise:
		if game.Players[0].chips < *val {
			fmt.Printf("Not enough money.Go for a lower raise or All-In. Retry input.\n")
			game.waitUserInput(Same)
		} else {
			game.MainPot.value += *val
			//NOTE: assuming p1 by default for now
			game.Players[0].chips -= *val
			// game.isRaised = true  //Uncomment when multiplayer mode is active
			//Calling `Call` .Fix when multiplater
			for i := 1; i < len(game.Players); i++ {
				if game.Players[i].chips >= *val {
					game.Players[i].chips -= *val
					game.MainPot.value += *val
				} else {
					game.Players[i].chips = 0
					game.MainPot.value += game.Players[i].chips
				}
			}
		}
	case Fold:
		//NOTE: assuming p1 by default for now
		fmt.Printf("=>> Folding you out in this game. You have %v chips remaining\n", game.Players[0].chips)
		os.Exit(1)
	case AllIn:
		// game.isRaised = true  //Uncomment when multiplayer mode is active
		game.MainPot.value += game.Players[0].chips
		*val = game.Players[0].chips
		game.Players[0].chips = 0
	}
	fmt.Printf("==== Pot Value: %v\n", game.MainPot.value)
}

func (game *Game) waitUserInput(e Event) {
	var input Action
	var val Money = 0
	fmt.Printf("=== Step: %v ===\n", eventName[e])
	fmt.Printf("Possible Actions:(Press Accordingly) . Always press `0` to quit\n")
	for _, a := range validAction[e] {
		fmt.Printf("# For `%v` press `%v`\n", actionName[a], a)
	}
Here:
	fmt.Scanf("%d\n", &input)
	if input == 0 {
		fmt.Println(" -_- -_- -_- Fuck you.Bye! -_- -_- -_-")
		os.Exit(1)
	} else if _, ok := containsAction(validAction[e], input); !ok {
		fmt.Println("Sorry, I didn't understand that input.Try again")
		goto Here
	} else if input == Raise {
		fmt.Println(" Raise by how much?")
		fmt.Scanf("%d\n", &val)
	}
	//Action Handler
	game.ActionHandler(input, &val)
}

//TODO: make init()for this
var p1 = Player{
	chips: STAKES,
}
var p2 = Player{
	chips: STAKES,
}

var table Hand

var game = Game{
	MainPot: MainPot{
		value: 0,
	},
	Players:  []Player{p1, p2},
	isRaised: false,
}

func Play() {
	deck := NewDeck()

	fmt.Println("Playing as Player1")
	fmt.Printf("STARTING GAME.You have %v chips\n", game.Players[0].chips)
	//NOTE: assue we are playins as p1 for now

	game.waitUserInput(Ante)

	game.waitUserInput(BeforePocketDraw)

	p1.PocketCards, deck = GetHand(n, deck)
	p2.PocketCards, deck = GetHand(n, deck)

	fmt.Printf(">> p1 cards: ")
	PrintCards(p1.PocketCards)
	// fmt.Printf(">> p2 cards:  *  * ")  //Redacting P2's cards
	PrintCards(p2.PocketCards)

	game.waitUserInput(TableDraw0)
	table, deck = GetHand(3, deck)
	fmt.Printf(">> table cards: ")
	PrintCards(table)

	game.waitUserInput(TableDraw1)
	var t []Card
	t, deck = GetHand(1, deck)
	table = append(table, t...)
	fmt.Printf(">> table cards: ")
	PrintCards(table)

	game.waitUserInput(TableDraw2)
	t, deck = GetHand(1, deck)
	table = append(table, t...)
	fmt.Printf(">> table cards: ")
	PrintCards(table)

	game.waitUserInput(BeforeShowdown)

	winner, handType, hand := GetWinner(p1, p2, table)
	fmt.Printf("Player %v won due to higher hand: %v - ", winner, handType)
	PrintCards(hand)
	game.Players[winner-1].chips += game.MainPot.value
	game.MainPot.value = 0
	fmt.Printf("GAME ENDED.You(p1) have %v chips\n", game.Players[0].chips)
	fmt.Printf("GAME ENDED.P2 has %v chips\n", game.Players[1].chips)
}
