package blackjack

import (
	"errors"
	"fmt"

	deck "github.com/iwita/go-exercises/deckOfCards"
)

const (
	statePlayerTurn state = iota
	stateDealerTurn
	stateHandOver
)

type state int8

type Options struct {
	Decks           int
	Hands           int
	BlackjackPayout float64
}

func New(opts Options) Game {
	g := Game{
		state:    statePlayerTurn,
		dealerAI: dealerAI{},
		balance:  0,
	}
	if opts.Decks == 0 {
		opts.Decks = 3
	}
	if opts.Hands == 0 {
		opts.Hands = 100
	}
	if opts.BlackjackPayout == 0.0 {
		opts.BlackjackPayout = 1.5
	}
	g.nDecks = opts.Decks
	g.nHands = opts.Hands
	g.blackjackPayout = opts.BlackjackPayout

	return g
}

type Game struct {
	nDecks          int
	nHands          int
	blackjackPayout float64

	deck  []deck.Card
	state state

	player    []deck.Card // this is a slice behind the scenes
	playerBet int
	balance   int

	dealer   []deck.Card
	dealerAI AI
}

func (g *Game) currentHand() *[]deck.Card {
	switch g.state {
	case statePlayerTurn:
		return &g.player
	case stateDealerTurn:
		return &g.dealer
	default:
		panic("it is not currently any player's turn")
	}
}

func bet(g *Game, ai AI, shuffled bool) {
	bet := ai.Bet(shuffled)
	g.playerBet = bet
}

type hand struct {
	cards []deck.Card
	bet   int
}

func deal(g *Game) {
	g.player = make([]deck.Card, 0, 5)
	g.dealer = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, g.deck = draw(g.deck)
		g.player = append(g.player, card)
		card, g.deck = draw(g.deck)
		g.dealer = append(g.dealer, card)

	}
	g.state = statePlayerTurn
}

func (g *Game) Play(ai AI) int {
	g.deck = nil
	min := 52 * g.nDecks / 3 //3 is an arbitary number
	// that indicates when should I reshuffle the rest of the deck
	for i := 0; i < g.nHands; i++ {
		shuffled := false
		if len(g.deck) < min {
			shuffled = true
			g.deck = deck.New(deck.Deck(g.nDecks), deck.Shuffle)

		}
		bet(g, ai, shuffled)
		deal(g)

		// if the dealer gets a BlackJack (Ace and Jack) dealt
		// the current game is over
		if Blackjack(g.dealer...) {
			endHand(g, ai)
			continue
		}
		for g.state == statePlayerTurn {
			hand := make([]deck.Card, len(g.player))
			copy(hand, g.player)
			move := ai.Play(hand, g.dealer[0])
			err := move(g)
			switch err {
			case errBust:
				MoveStand(g)
			case nil:
				// no op
			default:
				panic(err)
			}
		}

		// // If dealer score <= 16, we hit
		// // If dealer score has a soft 17, then we hit (with Aces inside)
		for g.state == stateDealerTurn {
			hand := make([]deck.Card, len(g.dealer))
			copy(hand, g.dealer)
			move := g.dealerAI.Play(hand, g.dealer[0])
			move(g)
		}

		endHand(g, ai)
	}
	return g.balance
}

var errBust = errors.New("hand score exceeded 21")

type Move func(*Game) error

func MoveHit(g *Game) error {
	hand := g.currentHand()
	var card deck.Card
	card, g.deck = draw(g.deck)
	*hand = append(*hand, card)
	if Score(*hand...) >= 21 {
		return errBust
	}
	return nil
}

func MoveDouble(g *Game) error {
	if len(g.player) != 2 {
		return errors.New("Can only double on a hand with 2 cards")
	}
	g.playerBet *= 2
	MoveHit(g)
	return MoveStand(g)
}

func MoveStand(g *Game) error {
	g.state++
	return nil
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}
func Score(hand ...deck.Card) int {
	minScore := MinScore(hand...)

	// if we use our Aces as 11 in this case
	// we get burned
	if minScore > 11 {
		return minScore
	}
	for _, c := range hand {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

func Soft(hand ...deck.Card) bool {
	minScore := MinScore(hand...)
	score := Score(hand...)
	return minScore != score
}

func Blackjack(hand ...deck.Card) bool {
	return len(hand) == 2 && Score(hand...) == 21
}

func MinScore(hand ...deck.Card) int {
	score := 0
	for _, c := range hand {
		// in order to deal with (J,Q,K)
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func endHand(g *Game, ai AI) {
	pScore, dScore := Score(g.player...), Score(g.dealer...)
	pBlackjack, dBlackjack := Blackjack(g.player...), Blackjack(g.dealer...)
	winnings := g.playerBet
	// TODO figure out  winnings and add/substract them
	switch {
	case pBlackjack && dBlackjack:
		winnings = 0
	case dBlackjack:
		winnings *= -1
	case pBlackjack:
		winnings = int(float64(winnings) * g.blackjackPayout)
	case pScore > 21:
		winnings *= -1
	case dScore > 21:
		// no op
	case pScore > dScore:
		// no op
	case dScore > pScore:
		winnings *= -1
	case dScore == pScore:
		winnings = 0
	}
	g.balance += winnings
	fmt.Println(g.balance)
	ai.Results([][]deck.Card{g.player}, g.dealer)
	g.player = nil
	g.dealer = nil
}
