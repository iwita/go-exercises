package main

import (
	"fmt"

	"github.com/iwita/go-exercises/blackjack_AI/blackjack"
)

func main() {
	game := blackjack.New()
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
