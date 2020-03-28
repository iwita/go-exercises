# What is a game state (GS)?

### Initial State

Deck: [A, 10, 10, 8, 2, ...]
Turn: NewGame
Player Hand: []
Dealer Hand: []

### An event

```golang
func Deal(gs GameState) GameState{
    ret := clone(gs)
    .. ret // do some work
    return ret
}
```

### Updated State
Deck: [2, ...]
Turn: HandsDealt
PlayerHand: [A, 10]
DealerHand: [10, 8]