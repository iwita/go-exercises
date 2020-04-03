package blackjack

import (
	"fmt"

	deck "github.com/iwita/go-exercises/deckOfCards"
)

type AI interface {
	Play(hand []deck.Card, dealer deck.Card) Move
	Bet(shuffled bool) int
	Results(hands [][]deck.Card, dealer []deck.Card)
}

type dealerAI struct{}

func (ai dealerAI) Bet(shuffled bool) int {
	return 1
}

func (ai dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	dScore := Score(hand...)
	if dScore <= 16 || dScore == 17 && Soft(hand...) {
		return MoveHit
	} else {
		return MoveStand
	}
}

func (ai dealerAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	// noop
}

func HumanAI() AI {
	return humanAI{}
}

type humanAI struct{}

func (ai humanAI) Bet(shuffled bool) int {
	fmt.Println("What would you like to bet?")
	var bet int
	fmt.Scanf("%d\n", &bet)
	return bet
}

func (ai humanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	fmt.Println("Player: ", hand)
	fmt.Println("Dealer: ", dealer)
	fmt.Println("What will you do? (h)it or (s)tand or (d)ouble, s(p)lit")
	var input string
	for {
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		case "d":
			return MoveDouble
		case "p":
			return MoveSplit
		default:
			fmt.Println("Invalid option: ", input)
		}
	}
}

func (ai humanAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:")
	for _, h := range hands {
		fmt.Println(" ", h)
	}
	fmt.Println("Dealer: ", dealer)
}
