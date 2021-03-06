# goker
Poker in Go.

# How to run
Clone the repo & run `go run goker/main.go` to start playing

<p align="center">

<img src="https://cdn.shopify.com/s/files/1/2022/6883/products/cards_on_table_1024x1024@2x.jpg?v=1557787138" height= "500px" >
</p>

### Inspirations
* For Logic & good UI: [danielpaz6/Poker-Hand-Evaluator](https://github.com/danielpaz6/Poker-Hand-Evaluator)
* [joker](https://github.com/notnil/joker)
* [joker-ui](https://jspoker.net/Room/Index/3)

### References
* [Poker Terms](https://www.wsop.com/poker-terms/)

## [x] itr#1: Single Player-No bet, CLI
- [x] Draw Cards
    - [x] Deck Generation
    - [x] Deck Shuffle
    - [x] Assign cards: table, p1, p2
- [x] **Evaluator Score** 
    - Possible Ways: 
        1. 52!/(7!*45!) Look-up table 
        2. some over-the-head x13 conversion logic 
    - Approach:(inspired from [jspoker](https://github.com/danielpaz6/Poker-Hand-Evaluator))
        - The idea is to give results firstly on the basis of HandRank only.But in case of tie, we don't want to go back to that Hand & compare all the tied-hands(lol).So give each hand a score, which reflects both: HandRank & a (normalized) score(b/w 0-100) based on cards in the hand, because once the HandRank is same, all that remains to compare is who's got bigger cards.
        - Out of 7 cards(5@table + 2@player), check if any 5 exist in a Hand(in decreasing order - ofcourse), if yes; calculate score as below:
        - Score = HandRankScore + NormalizedCardsScore
        - 1. HandRankValue:
              | HandRank |Value|
              | ---------| ----|
              | HighCard| 0 |
              | OnePair | 100 |
              | TwoPair | 200 |
              | ThreeOfAKind | 300 |
              | Straight | 400 |
              | Flush | 500 |
              | FullHouse | 600 |
              | FourOfAKind  | 700 |
              | StraightFlush | 800 |
              | RoyalFlush    | 900 |
        - 2. NormalizedCardsScore: 
            - cardRank(`a`) = Rank - 2 . So `2` becomes `0` & `A` becomes `12`
            - Sort all 5 cards in ascending order: `2278K`
            - CardsScore: (i={0,4})sum(a<sub>i</sub> * 13<sup>i</sup> )
            - Normalizer = (i={0,4})sum(12 * 13<sup>i</sup> ) = 433174  (12 becomes 13 when `Joker` will be added in future)
            - NormalizedCardsScore = CardsScore/Normalizer  (Comes in [0,100])
        - ASSUMPTIONS: Comparison in Scoring(in dec order of priority): HandRankValue(100) > BiggerHandRank(90) > HandValue(10)
            - In FullHouse: BiggerHandRank(60(x3)+30(x2))
            - In TwoPair: BiggerHandRank(45+45)
            - In HighCard: BiggerHandRank = 0
        - *Edge Case which led to change in Approcah*: What if the table cards have highest HandRank, showing it's rank to user is not justice(even though you're calculting user's Hand value in score calculation, but you cant show the **next best HandType**)
        - **Solution:** Need to maintain all the Hand-matches(in descending order) into a `Queue` of `Result` struct; so that while comparing 2 players, keep on popping from queues while both are having same HandType.This is necessary because we know that our score is more than enough to give the correct result, but the user-satisfaction is not good with it- They need to SEE it to BELIEVE it.
        - ```go
            type Result struct{
                HandType string
                Score float64
                Cards []Card
            }        
- [x] Write Tests & Make all Green

## [x] itr#2 : Single-player, Betting, CLI

### Tasks
- [x] Step-wise card assigning on table:
    - [x] Support multiple user inputs as actions : {call, raise, fold} & add check if they are valid at that step.
    - Events & their valid Actions:
        - Ante : [PutBlind, SitOut]
        - BeforePocketDraw : [Check, Raise, Call,Fold, AllIn]
        - PocketDraw: nil
        - TableDraw0 : [Check, Raise, Call,Fold, AllIn]
        - TableDraw1 : [Call, Raise, Check,Fold, AllIn]
        - TableDraw2 : [Call, Raise, Check,Fold, AllIn]
        - BeforeShowdown: [Call, Raise, Check,Fold, AllIn]
        - Showdown : nil
- [x] BettingOn:
    - Integrate actions with them.
- [x] Card Redaction(for bot's cards)

## Itr#3: Multiplayer, SingleTable, API : {in progress}
- [] Money distribution logic : [idea]((https://github.com/danielpaz6/Poker-Hand-Evaluator))
- [] Add new players
- [] Multiple continuous games.
- [] Connect with APIs.