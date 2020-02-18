package goker

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	_ Suit   = iota
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

var symbol = map[Rank]string{
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
	Spade: "♠",
	Club:  "♣",
	Diamond:"♦",
	Heart: "♥",
}

type Card struct {
	Rank
	Suit
}

type Deck []Card

type Hand []Card

type Player struct{
	FiveCards Hand
	SevenCards Hand
}

func PrintCards(cards []Card) {
	fmt.Printf("[")
	for _, card := range cards {
		fmt.Printf("{ %v%v }", symbol[card.Rank], symbol[card.Suit])
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

type HandType uint8

const (
	RoyalFlush    = 900
	StraightFlush = 800
	FourOfAKind   = 700
	FullHouse     = 600
	Flush         = 500
	Straight      = 400
	ThreeOfAKind  = 300
	TwoPair       = 200
	OnePair       = 100
	HighCard      = 0
)

func areConsecutive(cards []Card) bool {  //NOTE: function takes sorted array
	for i:= 0 ; i< len(cards) -1 ;i++{
		if cards[i].Rank != cards[i+1].Rank{
			return false
		}
	}
	return true
}

func intersection(a, b []Card) (c []Card) {
	m := make(map[Card]bool)

	for _, item := range a {
			m[item] = true
	}

	for _, item := range b {
			if _, ok := m[item]; !ok {
					c = append(c, item)
			}
	}
	return
}

func nOfSameSuit(h Hand, n int)([]Card, bool){
	m := make(map[Suit][]Card)
	for i := len(h)-1; i>=0 ;i-- {
		m[h[i].Suit] = append(m[h[i].Suit], h[i])
		
		if len(m[h[i].Suit]) == n {
			return m[h[i].Suit], false
		}
	}
	return nil, false
}

func nOfSameRank(h Hand, n int)([]Card, bool){ //Always returns the hieghst cards
	m := make(map[Rank][]Card)
	for i := len(h)-1; i>=0 ;i-- {
		m[h[i].Rank] = append(m[h[i].Rank], h[i])
		
		if len(m[h[i].Rank]) == n {
			return m[h[i].Rank], false
		}
	}
	return nil, false
}

func nPair(h Hand, n int)([]Card, bool){
	ret := []Card(nil)
	cnt := 0
	for i:= len(h)-1 ; i>= 0; i-- {
		if h[i] == h[i-1] {
			cnt ++
			ret = append(ret, h[i], h[i-1])
			i--
		}
		if cnt == n {
			return ret, true
		}
	}
	return nil, false
}


func normalizedScore(cards []Card, n int) int{
	cardScore, normalizer := 0 , 0
	for i, c := range cards{
		cardScore += math.Pow(13,i)*(c.Rank - 2)
		normalizer += math.Pow(13,i)*12
	}
	return (cardScore*n)/(normalizer*100)
}


//900-RoyalFlush
func isRoyalFlush(h Hand) bool { //TODO: make function
	if ((h.contains(Card{Suit: Spade, Rank: Ten}) && h.contains(Card{Suit: Spade, Rank: Jack}) && h.contains(Card{Suit: Spade, Rank: Queen}) && h.contains(Card{Suit: Spade, Rank: King}) && h.contains(Card{Suit: Spade, Rank: Ace}))
	|| (h.contains(Card{Suit: Club, Rank: Ten}) && h.contains(Card{Suit: Club, Rank: Jack}) && h.contains(Card{Suit: Club, Rank: Queen}) && h.contains(Card{Suit: Club, Rank: King}) && h.contains(Card{Suit: Club, Rank: Ace}))
	|| (h.contains(Card{Suit: Diamond, Rank: Ten}) && h.contains(Card{Suit: Diamond, Rank: Jack}) && h.contains(Card{Suit: Diamond, Rank: Queen}) && h.contains(Card{Suit: Diamond, Rank: King}) && h.contains(Card{Suit: Diamond, Rank: Ace}))
	|| (h.contains(Card{Suit: Heart, Rank: Ten}) && h.contains(Card{Suit: Heart, Rank: Jack}) && h.contains(Card{Suit: Heart, Rank: Queen}) && h.contains(Card{Suit: Heart, Rank: King}) && h.contains(Card{Suit: Heart, Rank: Ace}))) {
		return true
	}
	return false
}

//800-StraightFlush
func isStraightFlush(h Hand)([]Card, bool){
	//check if >=5 cards are of same suit
	ret := []Card
	m := make(map[Suit][]Card)
	for _, card := range h{
		m[card.Suit] = append(m[card.Suit], card)
	}
	for suit, count := range m{
		if count>=5 {
			ret = m[suit]
		}
	}
	if len(ret) <5 {
		return nil, false
	}
	ret =  ret[len(ret)-5:]
	
	//check for AP
	if areConsecutive(ret) {
		return ret, true
	}
	return nil, false
}

//700-FourOfAKind
func isFourOfAKind(h Hand)([]Card, bool){
	return nOfSameRank(h, 4)
}

//600-FullHouse
func isFullHouse(h Hand)([]Card, []Card, bool){
	cards3, ok3 := nOfSameRank(h, 3)
	h = intersection(h, cards3)
	cards2, ok2 := nOfSameRank(h, 2)
	if ok2 && ok3 {
		return cards3, cards2, true
	}
	return nil, nil, false
}

//500-Flush
func isFlush(h Hand)([]Card, bool){
	return nOfSameSuit(h, 5)
}

//400-Straight
func isStraight(h Hand)([]Card, bool){
	for i := 0 ;i<3; i++ {
		_h := h[len(h)-5-i:len(h)-i]
		if areConsecutive(_h){
			return _h, ture
		}
	}
	return nil, false
}


//300-ThreeOfAKind
func isThreeOfAKind(h Hand)([]Card, bool){
	return nOfSameRank(h, 3)
}

//200-TwoPair
func isTwoPair(h Hand)([]Card, bool){
	return nPair(h,2)
}


//100-OnePair
func isOnePair(h Hand)([]Card, bool){
	return nPair(h,1)
}


func Score(p Player) int { 
	sort.Slice(p1.FiveCards, Less(p1.FiveCards))
	sort.Slice(p2.FiveCards, Less(p2.FiveCards))
	sort.Slice(p1.SevenCards, Less(p1.SevenCards))
	sort.Slice(p2.SevenCards, Less(p2.SevenCards))

	ans := 0
	if ok := isRoyalFlush(p.SevenCards);ok {
		return RoyalFlush
	}
	if cards, ok := isStraightFlush(p.SevenCards);ok{
		ans =  StraightFlush + normalizedScore(cards, 90)
	}
	if cards, ok := isFourOfAKind(p.SevenCards);ok{
		ans =  FourOfAKind + normalizedScore(cards, 90)
	}
	if cards3, cards2, ok := isFullHouse(p.SevenCards);ok{
		ans =  FourOfAKind + normalizedScore(cards3, 60) + normalizedScore(cards2, 30)
	}
	if cards, ok := isFlush(p.SevenCards);ok{
		ans =  Flush + normalizedScore(cards, 90)
	}
	if cards, ok := isStraight(p.SevenCards);ok{
		ans =  Straight + normalizedScore(cards, 90)
	}
	if cards, ok := isThreeOfAKind(p.SevenCards);ok{
		ans =  ThreeOfAKind + normalizedScore(cards, 90)
	}
	if cards1, cards2, ok := isTwoPair(p.SevenCards);ok{
		ans =  TwoPair + normalizedScore(cards1, 45)+ normalizedScore(cards2, 45)
	}
	if cards, ok := isOnePair(p.SevenCards);ok{
		ans =  OnePair + normalizedScore(cards, 90)
	}
	//else HighCard

	return ans + normalizedScore(p.FiveCards, 10)
}


func GetWinner(p1, p2 Player, table Hand) int {
	p1.SevenCards = append(p1.FiveCards, table...)
	p2.SevenCards = append(p2.FiveCards, table...)

	if Score(p1) > Score(p2) {
		return 1
	} else {
		return 2
	}
}
