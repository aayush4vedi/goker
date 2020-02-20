package goker

const BLIND Money = 10
const STAKES Money = 100

type Player struct {
	PocketCards   Hand
	CommunityCards Hand
	chips Money
}

type MainPot struct{
	value Money
}

type Game struct {
	MainPot
	Players []Player  //= 2 ;compromised for #itr2
	isRaised bool 
}