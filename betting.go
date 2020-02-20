package goker

type Action uint8

const (
	_ Action = iota
	PutBlind	
	SitOut
	Raise		//need value
	Call		//Depends on Raise
	Check
	Fold
	AllIn		
)
var actionName = map[Action]string{
	PutBlind : "PutBlind",
	SitOut : "SitOut",
	Call : "Call",
	Check : "Check",
	Raise : "Raise",
	Fold : "Fold",
	AllIn : "AllIn",
}

type Event uint8

const (
	_ Event = iota
	Ante
	BeforePocketDraw
	PocketDraw
	TableDraw0 
	TableDraw1
	TableDraw2
	BeforeShowdown
	Showdown
	Same  //when returning from wrong Action, back to same event for another input
)
var eventName = map[Event]string{
	Ante : "Ante",
	BeforePocketDraw : "BeforePocketDraw",
	PocketDraw : "PocketDraw",
	TableDraw0 : "TableDraw0",
	TableDraw1 : "TableDraw1",
	TableDraw2 : "TableDraw2",
	BeforeShowdown : "BeforeShowdown",
	Showdown : "Showdown",
	Same : "Same",
}

var validAction = map[Event][]Action{
	Ante : []Action{PutBlind, SitOut},
	BeforePocketDraw : []Action{Check, Raise, Call,Fold, AllIn},
	TableDraw0 : []Action{Check, Raise, Call,Fold, AllIn},
	TableDraw1 : []Action{Call, Raise, Check,Fold, AllIn},
	TableDraw2 : []Action{Call, Raise, Check,Fold, AllIn},
	BeforeShowdown: []Action{Call, Raise, Check,Fold, AllIn},
	Same : []Action{Call, Raise, Check,Fold, AllIn},
}

type Money int