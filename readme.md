# goker
Poker in Go.
[inspiration](https://github.com/notnil/joker)
[ui](https://jspoker.net/Room/Index/3)

## itr#1. 
Single Player-No bet, CLI

### Task
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
        
- [] Write Tests & Make all Green

## itr#2.