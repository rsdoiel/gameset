// deck.go models a user defined deck of cards organized around suites and
// face card names.
package deck

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// Deck holds the state of cards in play. I represents generally
// a deck based on the standard cards organized around suites and
// faces.
type Deck struct {
	Game         string              `json:"game,omitempty"`
	Players      []string            `json:"players,omitempty"`
	Cards        []string            `json:"cards"`
	Play         []string            `json:"play"`
	Discarded    []string            `json:"discard"`
	HandsHeld    map[string][]string `json:"heldHands"`
	HandsVisible map[string][]string `json:"visibleHands"`
}

// MakeDeck builds an set of cards based on symbol (string)
// and values (int). E.g.
//
// ```
//
// suites := []string{ "Hart", "Club", "Diamond", "Spade" }
// faces := []string{ "Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack","Queen", "King" }
// deck := MakeDeck("hi_card_wins", suites, faces)
// ```
func MakeDeck(suites []string, faces []string) *Deck {
	deck := new(Deck)
	deck.Game = ""
	deck.Players = []string{}
	deck.Cards = []string{}
	deck.Play = []string{}
	deck.Discarded = []string{}
	deck.HandsHeld = map[string][]string{}
	deck.HandsVisible = map[string][]string{}
	for _, suit := range suites {
		for _, face := range faces {
			var card string
			if suit != "" {
				card = fmt.Sprintf("%s%s", face, suit)
			} else {
				card = face
			}
			deck.Cards = append(deck.Cards, card)
			deck.Play = append(deck.Play, card)
		}
	}
	return deck
}

// MakeStandardDeck creates a standard deck of playing cards.
func MakeStandardDeck() *Deck {
	suites := []string{
		"Harts", "Clubs", "Diamonds", "Spades",
	}
	faces := []string{
		"Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine",
		"Ten", "Jack", "Queen", "King", "Ace",
	}
	return MakeDeck(suites, faces)
}

func (deck *Deck) SetGame(name string) {
	deck.Game = name
}

func (deck *Deck) SetPlayers(names []string) {
	deck.Players = names[:]
}

// Shuffles randomizes order of the cards in Play
//
// ```
//
//	deck.Shuffle()
//
// ```
func (deck *Deck) Shuffle() {
	if len(deck.Play) == 0 {
		deck.Play = deck.Cards[:]
	}
	rand.Shuffle(len(deck.Play), func(i, j int) {
		deck.Play[i], deck.Play[j] = deck.Play[j], deck.Play[i]
	})
}

// DiscardToPlay takes the list of discarded cards and returns
// them to play.
func (deck *Deck) DiscardToPlay() {
	deck.Play = append(deck.Play, deck.Discarded...)
}

// Deal takes an ordered list of players and a count and
// forms hands for each player by randomly drawing from the
// deck in play.
//
// ```
// // cards are dealt to players in order of slice
// err := deck.Deal([]string{"player1", "player2"}, 3)
// ```
func (deck *Deck) Deal(players []string, cardCount int) error {
	var card string

	if len(players) < 1 {
		return fmt.Errorf("not enough players")
	}

	if len(deck.Play) < (len(players) * cardCount) {
		return fmt.Errorf("not enough cards to deal")
	}

	for i := 0; i < cardCount; i++ {
		for _, player := range players {
			card, deck.Play = deck.Play[0], deck.Play[1:]
			deck.HandsHeld[player] = append(deck.HandsHeld[player], card)
		}
	}
	return nil
}

// Hand takes a players name and returns the cards they are controlling.
// (both shown and not shown). If the player is not found then nil is
// returned.
//
// ```
// hand1 := deck.Hand("player1")
// hand2 := deck.Hand("player2")
// ```
func (deck *Deck) Hand(player string) []string {
	if hand, ok := deck.HandsHeld[player]; ok {
		return hand[:]
	}
	return nil
}

func (deck *Deck) HandVisible(player string) []string {
	if hand, ok := deck.HandsVisible[player]; ok {
		return hand[:]
	}
	return nil
}

// Show makes a card in the hand visible and returns the name of the card
// shown. If card is not found then it returns an empty string
//
// ```
// // Each player shows one card from their hand
// fmt.Printf("Player 1: %q, Player 2: %q\n", deck.Show("player1", hand1[1]), deck.Show("player2", handl2[2]))
// ```
func (deck *Deck) Show(player string, card string) string {
	cards := deck.Hand(player)
	pos := -1
	for i, aCard := range cards {
		if strings.Compare(aCard, card) == 0 {
			pos = i
			break
		}
	}
	if pos < 0 {
		return ""
	}
	if pos == 0 {
		deck.HandsHeld[player] = cards[1:]
	} else {
		deck.HandsHeld[player] = append(cards[0:pos], cards[pos+1:]...)
	}
	deck.HandsVisible[player] = append(deck.HandsVisible[player], card)
	return card
}

// Shown returns a list of cards for a player that are visible
func (deck *Deck) Shown(player string) []string {
	if hand, ok := deck.HandsVisible[player]; ok {
		return hand[:]
	}
	return nil
}

// Draw, similar to deal but selects a number of cards from the
// top of the deck putting it in the players' hand. Returns
// the new set of cards in the hand.
//
// ```
// hand1, err = deck.Draw("player1", 1)
// // ... handle errors ...
// hand2, err = deck.Draw("player2", 1)
// // ... handle errors ...
// ```
func (deck *Deck) Draw(player string, count int) ([]string, error) {
	if count >= len(deck.Play) {
		return nil, fmt.Errorf("not enough cards to draw")
	}
	drawn, remaining := deck.Play[0:count], deck.Play[count+1:]
	deck.HandsHeld[player] = append(deck.HandsHeld[player], drawn...)
	deck.Play = remaining[:]
	return drawn, nil
}

// Discard takes a player's name and discard's the card indicated
// by the card's name. Returns the remaining cards.
//
// ```
// hand1, err = deck.Discard("player1", hand1[2])
// // ... handle errors ...
// hand2, err = deck.Discard("player2", hand2[3])
// // ... handle errors ...
// // Now show all cards, player with highest card wins.
// fmt.Printf("Player 1 has %q, Player 2 has %q\n", strings.Join(hand1, ", "), strings.Join(hand2, ", "))
// ```
func (deck *Deck) Discard(player string, card string) ([]string, error) {
	var discarded string

	cards := deck.Hand(player)
	if len(cards) == 0 {
		return nil, fmt.Errorf("not cards held to discard")
	}
	pos := -1
	for i, aCard := range cards {
		if strings.Compare(aCard, card) == 0 {
			pos = i
			break
		}
	}
	if pos < 0 {
		return nil, fmt.Errorf("failed to find %q", card)
	}
	discarded = cards[pos]
	if pos == 0 {
		deck.HandsHeld[player] = cards[1:]
	} else {
		deck.HandsHeld[player] = append(cards[0:pos], cards[pos+1:]...)
	}
	deck.Discarded = append(deck.Discarded, discarded)
	return nil, fmt.Errorf("not implemented")
}

// Pickup takes the last N cards discarded and puts them into
// a players' hand.
func (deck *Deck) Pickup(player string, n int) error {
	if n >= len(deck.Discarded) {
		return fmt.Errorf("not enough discarded cards")
	}
	drawn, remaining := deck.Discarded[len(deck.Discarded)-n:], deck.Discarded[0:len(deck.Discarded)-n]
	deck.HandsHeld[player] = append(deck.HandsHeld[player], drawn...)
	deck.Discarded = remaining[:]
	return nil
}

// TakeHeld takes a card by card name from a player and
// and puts it into the hand of another player. Returns an error
// if not available. E.g. like when playing Go Fish
func (deck *Deck) TakeHeld(originName string, receivingName string, card string) error {
	originHand := deck.Hand(originName)
	if len(originHand) == 0 {
		return fmt.Errorf("%s has no cards", originName)
	}
	pos := -1
	for i, aCard := range originHand {
		if strings.Compare(aCard, card) == 0 {
			pos = i
			break
		}
	}
	if pos == -1 {
		return fmt.Errorf("card not found in %s's hand", originName)
	}
	if pos == 0 {
		deck.HandsHeld[originName] = deck.HandsHeld[originName][1:]
	} else {
		deck.HandsHeld[originName] = append(deck.HandsHeld[originName][0:pos], deck.HandsHeld[originName][pos+1:]...)
	}
	deck.HandsHeld[receivingName] = append(deck.HandsHeld[receivingName], card)
	return nil
}

// TakeVisible takes a card by card name from a player and
// puts it into a player's held hand. Returns an error if card not
// available.
func (deck *Deck) TakeVisible(originName string, receivingName string, card string) error {
	originHand := deck.HandVisible(originName)
	if len(originHand) == 0 {
		return fmt.Errorf("%s has no cards", originName)
	}
	pos := -1
	for i, aCard := range originHand {
		if strings.Compare(aCard, card) == 0 {
			pos = i
			break
		}
	}
	if pos == -1 {
		return fmt.Errorf("card not found in %s's hand", originName)
	}
	if pos == 0 {
		deck.HandsVisible[originName] = deck.HandsVisible[originName][1:]
	} else {
		deck.HandsVisible[originName] = append(deck.HandsVisible[originName][0:pos], deck.HandsVisible[originName][pos+1:]...)
	}
	deck.HandsHeld[receivingName] = append(deck.HandsHeld[receivingName], card)
	return nil
}

func inList(target string, list []string) bool {
	for _, val := range list {
		if strings.Compare(target, val) == 0 {
			return true
		}
	}
	return false
}

// String returns a report on the state of cards in game
// in a human readable form.
func (deck *Deck) String() string {
	players := []string{}

	if len(deck.Players) == 0 {
		for player, _ := range deck.HandsHeld {
			if !inList(player, players) {
				players = append(players, player)
			}
		}
		for player, _ := range deck.HandsVisible {
			if !inList(player, players) {
				players = append(players, player)
			}
		}
		sort.Strings(players)
	} else {
		players = deck.Players[:]
	}
	parts := []string{}
	for _, player := range players {
		if hand, ok := deck.HandsHeld[player]; ok {
			parts = append(parts, fmt.Sprintf("%s holds %s", player, strings.Join(hand, ", ")))
		} else {
			parts = append(parts, fmt.Sprintf("%s holds no more cards", player))
		}
		if hand, ok := deck.HandsVisible[player]; ok {
			parts = append(parts, fmt.Sprintf("%s shows %s", player, strings.Join(hand, ", ")))
		} else {
			parts = append(parts, fmt.Sprintf("%s show no cards", player))
		}
		if len(deck.Play) > 0 {
			parts = append(parts, fmt.Sprintf("%d cards left to play", len(deck.Play)))
		} else {
			parts = append(parts, "no cards left to play")
		}
		if len(deck.Discarded) > 0 {
			parts = append(parts, fmt.Sprintf("%d cards in discard pile", len(deck.Discarded)))
		} else {
			parts = append(parts, "no cards in discard pile")
		}
	}
	return strings.Join(parts, "\n")
}

// ToJSON returns a JSON representation of the state of the came.
func (deck *Deck) ToJSON() ([]byte, error) {
	return json.MarshalIndent(deck, "", "    ")
}

// Load reads a JSON source slice of bytes and returns
// a deck.
func FromJSON(src []byte, deck *Deck) error {
	return json.Unmarshal(src, &deck)
}
