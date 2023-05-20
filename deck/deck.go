// deck.go models a user defined deck of cards organized around suites and
// face card names.
package deck

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
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
	HandsHeld    map[string][]string `json:"hands_held"`
	HandsVisible map[string][]string `json:"handls_visible"`
}

// String displays the cards in a deck
func (deck *Deck) String() string {
	return strings.Join(deck.Cards, ", ")
}

// AddSuite takes takes the name of a suite and the names of
// the faces and adds them to a `*Deck.Cards`.
func (deck *Deck) AddSuite(suite string, cards ...string) {
	for _, card := range cards {
		deck.Cards = append(deck.Cards, fmt.Sprintf("%s of %s", card, suite))
	}
}

// AddCard takes a name (e.g. "joker", "rook") and adds it
// to the `*Deck.Cards`
func (deck *Deck) AddCard(name string) {
	deck.Cards = append(deck.Cards, name)
}

// NewDeck builds an set of cards based on suite (slice of string)
// and face (slice of string). It each suite will be associated with
// a set of faces forming a set of cards. E.g.
//
// ```
//
// suites := []string{ "Hart", "Club", "Diamond", "Spade" }
// faces := []string{ "Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack","Queen", "King" }
// deck := NewDeck("hi_card_wins", suites, faces)
// fmt.Printf("%s", deck.String())
//
// ```
func NewDeck(suites []string, faces []string) *Deck {
	deck := new(Deck)
	deck.Game = ""
	deck.Players = []string{}
	deck.Cards = []string{}
	deck.Play = []string{}
	deck.Discarded = []string{}
	deck.HandsHeld = map[string][]string{}
	deck.HandsVisible = map[string][]string{}
	for _, suite := range suites {
		deck.AddSuite(suite, faces...)
	}
	return deck
}

// NewStandardDeck creates a standard deck of playing cards.
func NewStandardDeck() *Deck {
	suites := []string{
		"♣", "♢", "♡", "♠",
	}
	faces := []string{
		"2", "3", "4", "5",
		"6", "7", "8", "9",
		"10", "Jack", "Queen", "King", "Ace",
	}
	return NewDeck(suites, faces)
}

// NewPinochleDeck creates a Pinochle desck of playing cards.
func NewPinochleDeck() *Deck {
	suites := []string{
		"♣", "♢", "♡", "♠",
	}
	faces := []string{
		"9", "9",
		"10", "10",
		"Jack", "Jack",
		"Queen", "Queen",
		"King", "King",
		"Ace", "Ace",
	}
	return NewDeck(suites, faces)
}

// NewSakuraDeck creates a Sakura using hanafuda card deck
func NewSakuraDeck() *Deck {
	deck := new(Deck)
	deck.Cards = []string{}
	deck.AddSuite("pine", "crane of sun (20)", "peotry paper (5)", "kasu (1)", "kasu (1)")
	deck.AddSuite("plum", "bush warbler (10)", "peotry paper (5)", "kasu (1)", "kasu (1)")
	deck.AddSuite("cherry", "curtain (20)", "peotry paper (5)", "kasu (1)", "kasu (1)")
	deck.AddSuite("wisteria", "cuckoo (10)", "plain paper (5)", "kasu (1)", "kasu (1)")
	deck.AddSuite("iris", "eight-plank bridge (10)", "plain paper (5)", "kasu (1)", "kasu (1)")
	deck.AddSuite("peony", "butterflies (10)", "blue paper (5)", "kasu (1)", "kasu (1)")
	deck.AddSuite("bush clover", "boar (10)", "plain paper (5)", "kasu (1)", "kasu (1)")
	deck.AddSuite("susuki grass", "full moon (20)", "geese (10)", "kasu (1)", "kasu (1)")
	deck.AddSuite("chrysanthemum", "sake cup (10)", "blue paper (5)", "kasu (1)", "kasu (1)")
	deck.AddSuite("maple", "deer (10)", "blue paper (5)", "kasu (1)", "kasu (1)")
	deck.AddSuite("willow", "ono no Michikaze (20)", "swallow (10)", "plain paper (5)", "lightning (1)")
	deck.AddSuite("paulownia", "Chinese pheonix (20)", "kasu (1)", "kasu (1)", "kasu (1)")
	return deck
}

// NewHwatuDeck creates a Korean style Sakura deck
func NewHwatuDeck() *Deck {
	deck := new(Deck)
	deck.AddSuite("January 송학", "Bright 광", "Ribbon 띠", "Junk 피", "Junk 피")
	deck.AddSuite("February 매화", "Animal 끗", "Ribbon 띠", "Junk 피", "Junk 피")
	deck.AddSuite("March 벚꽃", "Bright 광", "Ribbon 띠", "Junk 피", "Junk 피")
	deck.AddSuite("April 등꽃", "Animal 끗", "Ribbon 띠", "Junk 피", "Junk 피")
	deck.AddSuite("May 난초", "Ribbon 띠", "Junk 피", "Junk 피", "Double Junk 쌍피")
	deck.AddSuite("June 모란", "Animal 끗", "Ribbon 띠", "Junk 피", "Junk 피")
	deck.AddSuite("July 흑싸리", "Animal 끗", "Ribbon 띠", "Junk 피", "Junk 피")
	deck.AddSuite("August 공산", "Animal 끗", "Bright 광", "Junk 피", "Junk 피")
	deck.AddSuite("September 국화", "Ribbon 띠", "Junk 피", "Junk 피", "Double Junk 쌍피")
	deck.AddSuite("October 단풍", "Animal 끗", "Ribbon 띠", "Junk 피", "Junk 피")
	deck.AddSuite("November 오동", "Bright 광", "Junk 피", "Junk 피", "Double Junk 쌍피")
	deck.AddSuite("December 비", "Animal 끗", "Bright 광", "Junk 피", "Double Junk 쌍피")
	return deck
}

// NewRookDeck creates a Rook card deck
func NewRookDeck() *Deck {
	suites := []string{
		"red", "yellow", "green", "black",
	}
	faces := []string{
		"1", "2", "3", "4", "5", "6", "7",
		"8", "9", "10", "11", "12", "13", "14",
	}
	deck := NewDeck(suites, faces)
	deck.AddCard("rook")
	return deck
}

// Reset clears the deck's state to a "new" like state.
//
// ```
//
// deck := NewStandardDeck()
// deck.shuffle()
// // now reset the deck.
// deck.Reset()
// ```
func (deck *Deck) Reset() {
	deck.Game = ""
	deck.Players = []string{}
	deck.Play = []string{}
	deck.Discarded = []string{}
	deck.HandsHeld = map[string][]string{}
	deck.HandsVisible = map[string][]string{}
}

// SetupGame sets up an existing deck of cards for initial play.
func (deck *Deck) SetupGame(gameName string, players []string) {
	deck.Game = gameName
	deck.Players = append([]string{}, players...)
	// Reset the deck of cards
	deck.Play = append([]string{}, deck.Cards...)
	deck.Discarded = []string{}
	deck.HandsHeld = map[string][]string{}
	deck.HandsVisible = map[string][]string{}
}

// Shuffles randomizes order of the cards in Play
//
// ```
//
//	deck.Shuffle()
//
// ```
func (deck *Deck) Shuffle() {
	rand.Shuffle(len(deck.Play), func(i, j int) {
		deck.Play[i], deck.Play[j] = deck.Play[j], deck.Play[i]
	})
}

// DiscardToPlay takes the list of discarded cards and returns
// them to play.
func (deck *Deck) DiscardToPlay() {
	deck.Play = append(deck.Play, deck.Discarded...)
}

// Deal takes an integer for the number of cards to be dealt
// based on the current list of players associated with deck
// and deals cards to each player.
//
// ```
// // cards are dealt to players in order of slice
// deck.SetupGame("high card", []string{"player1","player2"})
// err := deck.Deal(3)
// ```
func (deck *Deck) Deal(cardCount int) error {
	var card string

	if len(deck.Players) < 1 {
		return fmt.Errorf("not enough players")
	}

	if len(deck.Play) < (len(deck.Players) * cardCount) {
		return fmt.Errorf("not enough cards to deal")
	}

	for i := 0; i < cardCount; i++ {
		for _, player := range deck.Players {
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
	if count <= 0 {
		return nil, fmt.Errorf("count must be a positive integer greater than zero")
	}
	if count > len(deck.Play) {
		return nil, fmt.Errorf("not enough cards to draw %d -> %d", count, len(deck.Play))
	}
	var drawn []string

	// draw the cards from deck
	drawn = append(drawn, deck.Play[0:count]...)
	deck.HandsHeld[player] = append(deck.HandsHeld[player], drawn...)

	// Update the cards in play
	if count < len(deck.Play) {
		deck.Play = append([]string{}, deck.Play[count:]...)
	} else {
		deck.Play = []string{}
	}
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
func (deck *Deck) TakeHeld(requesting string, responding string, card string) error {
	respondingHand := deck.Hand(responding)
	if len(respondingHand) == 0 {
		return fmt.Errorf("%s has no cards", responding)
	}
	pos := -1
	for i, aCard := range respondingHand {
		if strings.Compare(aCard, card) == 0 {
			pos = i
			break
		}
	}
	if pos == -1 {
		return fmt.Errorf("card not found in %s's hand", responding)
	}
	deck.HandsHeld[responding] = append(deck.HandsHeld[responding][0:pos], deck.HandsHeld[responding][pos+1:]...)
	deck.HandsHeld[requesting] = append(deck.HandsHeld[requesting], card)
	return nil
}

// TakeVisible takes a card by card name from a player and
// puts it into a player's held hand. Returns an error if card not
// available.
func (deck *Deck) TakeVisible(requesting string, responding string, card string) error {
	respondingHand := deck.HandVisible(responding)
	if len(respondingHand) == 0 {
		return fmt.Errorf("%s has no cards", responding)
	}
	pos := -1
	for i, aCard := range respondingHand {
		if strings.Compare(aCard, card) == 0 {
			pos = i
			break
		}
	}
	if pos == -1 {
		return fmt.Errorf("card not found in %s's hand", responding)
	}
	deck.HandsVisible[responding] = append(deck.HandsVisible[responding][0:pos], deck.HandsVisible[responding][pos+1:]...)
	deck.HandsVisible[requesting] = append(deck.HandsVisible[requesting], card)
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

// Report returns the state of cards in game
// in a human readable form as a string.
func (deck *Deck) Report() string {
	players := []string{}

	if len(deck.Players) == 0 {
		for player := range deck.HandsHeld {
			if !inList(player, players) {
				players = append(players, player)
			}
		}
		for player := range deck.HandsVisible {
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
			parts = append(parts, fmt.Sprintf("%s holds no more card", player))
		}
		if hand, ok := deck.HandsVisible[player]; ok {
			parts = append(parts, fmt.Sprintf("%s shows %s", player, strings.Join(hand, ", ")))
		}
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

// Init initializes the rand number generator used in Shuffle
func Init() {
	rand.Seed(time.Now().UnixNano())
}

func HasCard(cards []string, target string) bool {
	for _, card := range cards {
		if strings.Compare(card, target) == 0 {
			return true
		}
	}
	return false
}

// PickACard is an example of randomly picking a card from a list
// of cards respecting any excluded cards (e.g. cards in someone's hand)
func PickACard(cards []string, excluded []string) string {
	cardsRemaining := []string{}
	if len(excluded) == 0 {
		cardsRemaining = append(cardsRemaining, cards...)
	} else {
		for _, aCard := range cards {
			if !HasCard(excluded, aCard) {
				cardsRemaining = append(cardsRemaining, aCard)
			}
		}
		rand.Shuffle(len(cardsRemaining), func(i, j int) {
			cardsRemaining[i], cardsRemaining[j] = cardsRemaining[j], cardsRemaining[i]
		})
	}
	if len(cardsRemaining) > 0 {
		return cardsRemaining[0]
	}
	return ""
}

// copyFile does a simple copy
func copyFile(destName string, srcName string) error {
	src, err := os.Open(srcName)
	if err != nil {
		return err
	}
	defer src.Close()
	dest, err := os.Create(destName)
	if err != nil {
		return err
	}
	defer dest.Close()
	if _, err := io.Copy(dest, src); err != nil {
		return err
	}
	return nil
}

// Save writes the deck to a file, replacing an existing file
// if it already exists.
func (deck *Deck) Save(fName string) error {
	if deck == nil {
		return fmt.Errorf("deck cannot be saved, not initialized")
	}
	if _, err := os.Stat(fName); err == nil {
		// It exists so rename it to fName.bak
		backupName := fmt.Sprintf("%s.bak", fName)
		if err := copyFile(backupName, fName); err != nil {
			return fmt.Errorf("failed to save backup %s, %s", backupName, err)
		}
	}
	src, err := json.MarshalIndent(deck, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to encode %s, %s\n", fName, err)
	}
	if err := ioutil.WriteFile(fName, src, 0664); err != nil {
		return err
	}
	return nil
}
