package deck

import (
	"fmt"
	"strings"
	"testing"
)

func TestDeck(t *testing.T) {
	deck := NewStandardDeck()
	if deck == nil {
		t.Errorf("expected a standard deck to be made")
		t.FailNow()
	}
	if len(deck.Cards) != 52 {
		t.Errorf("expected 52 cards in deck.Cards")
	}
}

func TestShuffle(t *testing.T) {
	deck := NewStandardDeck()
	if len(deck.Play) > 0 {
		t.Errorf("deck.Play should be empty before shuffle, %d cards %+v", len(deck.Play), deck.Play)
	}
	deck.SetupGame("TestShuffle", []string{"npc1"})
	deck.Shuffle()
	if len(deck.Play) != len(deck.Cards) {
		t.Errorf("expect len(deck.Cards) %d, got len(deck.Play) %d", len(deck.Cards), len(deck.Play))
	}
	elementsMatch := true
	for i := 0; i < len(deck.Play); i++ {
		if strings.Compare(deck.Cards[i], deck.Play[i]) != 0 {
			elementsMatch = false
		}
	}
	if elementsMatch {
		t.Errorf("failed to shuffle deck and populate play list, deck.Play is same order as deck.Cards")
	}
}

func TestDeal(t *testing.T) {
	deck := NewStandardDeck()
	players := []string{"npc1", "npc2"}
	deck.SetupGame("Test Deal", players)
	deck.Shuffle()
	if err := deck.Deal(3); err != nil {
		t.Errorf("deck.Deal(players, 3) should not return an error, %s", err)
	}
	hand1 := deck.Hand(players[0])
	if len(hand1) != 3 {
		t.Errorf("expected three cards in hand for %q, got %+v", players[0], hand1)
	}
	hand2 := deck.Hand(players[1])
	if len(hand2) != 3 {
		t.Errorf("expected three cards in hand for %q, got %+v", players[1], hand2)
	}
	// Confirmed cards have been removed from play
	playedCards := append(hand1, hand2...)
	for _, aCard := range playedCards {
		for _, card := range deck.Play {
			if strings.Compare(aCard, card) == 0 {
				t.Errorf("expect no match cound %q in %+v", aCard, deck.Play)
				break
			}
		}
	}
}

func TestTake(t *testing.T) {
	// Model the simple game GoFish
	deck := NewStandardDeck()
	players := []string{"npc1", "npc2"}
	deck.SetupGame("Test Go Fish", players)
	deck.Shuffle()
	deck.Deal(3)

	hand1 := deck.Hand(players[0])
	if len(hand1) != 3 {
		t.Errorf("hand1 should have 3 cards, got %d", len(hand1))
	}
	hand2 := deck.Hand(players[1])
	if len(hand2) != 3 {
		t.Errorf("hand2 should have 3 cards, got %d", len(hand2))
	}
	card := hand1[0]
	err := deck.TakeHeld(players[1], players[0], card)
	if err != nil {
		t.Errorf("expected %q, (hand1 %+v, hand2 %+v) got error %s", card, hand1, hand2, err)
	}
	hand1 = deck.Hand(players[0])
	if len(hand1) != 2 {
		t.Errorf("hand1 should have 2 cards, got %d", len(hand1))
	}
	hand2 = deck.Hand(players[1])
	if len(hand2) != 4 {
		t.Errorf("hand2 should have 4 cards, got %d", len(hand2))
	}

	card = hand1[0]
	err = deck.TakeHeld(players[1], players[0], card)
	if err != nil {
		t.Errorf("expected %q, (hand1 %+v, hand2 %+v) got error %s", card, hand1, hand2, err)
	}
	hand1 = deck.Hand(players[0])
	if len(hand1) != 1 {
		t.Errorf("hand1 should have 1 cards, got %d", len(hand1))
	}
	hand2 = deck.Hand(players[1])
	if len(hand2) != 5 {
		t.Errorf("hand2 should have 5 cards, got %d", len(hand2))
	}

	card = hand1[0]
	err = deck.TakeHeld(players[1], players[0], card)
	if err != nil {
		t.Errorf("expected %q, (hand1 %+v, hand2 %+v) got error %s", card, hand1, hand2, err)
	}
	hand1 = deck.Hand(players[0])
	if len(hand1) != 0 {
		t.Errorf("hand1 should have 0 cards, got %d", len(hand1))
	}
	hand2 = deck.Hand(players[1])
	if len(hand2) != 6 {
		t.Errorf("hand2 should have 6 cards, got %d", len(hand2))
	}
	if (len(hand2) + len(deck.Play)) != len(deck.Cards) {
		t.Errorf("expected %d cards, got total %d", len(deck.Cards), (len(hand2) + len(deck.Play)))
	}
}

func TestDraw(t *testing.T) {
	deck := NewStandardDeck()
	players := []string{"npc1"}
	deck.SetupGame("Tests Draw", players)
	cardCount := len(deck.Play)
	if cardCount != 52 {
		t.Errorf("expected 52 cards, got %d", cardCount)
	}
	for i := 0; i < 52; i += 1 {
		cards, err := deck.Draw(players[0], 1)
		if err != nil {
			t.Errorf("deck.Draw(%q, 1) returned (%d), %s", players[0], i, err)
			t.FailNow()
		}
		hand := deck.Hand(players[0])
		if !HasCard(hand, cards[0]) {
			t.Errorf("player %q should have card %s, it is missing", players[0], deck.Play[0])
		}
		if HasCard(deck.Play, cards[0]) {
			t.Errorf("deck should NOT have card %s", cards[0])
		}
	}
}

func TestGoFish(t *testing.T) {
	// Model the simple game GoFish
	deck := NewStandardDeck()
	players := []string{"npc1", "npc2"}
	deck.SetupGame("Test Go Fish", players)
	deck.Shuffle()
	deck.Deal(2)
	for i := 0; i < 1000; i += 1 {
		// Check state of current hand
		hand1 := deck.Hand(players[0])
		if len(hand1) == 0 {
			// player one looses
			fmt.Printf("%s lost", players[0])
			break
		}
		hand2 := deck.Hand(players[1])
		if len(hand2) == 0 {
			// player two lost
			fmt.Printf("%s lost", players[1])
			break
		}
		// Player one requests a card
		card1 := PickACard(deck.Cards, hand1)
		if card1 == "" {
			t.Errorf("player %q has no cards to pick", players[0])
			continue
		}
		// Player two requests a card
		card2 := PickACard(deck.Cards, hand2)
		if card2 == "" {
			t.Errorf("player %q has no cards to pick", players[1])
			continue
		}
		if err := deck.TakeHeld(players[0], players[1], card1); err != nil {
			// Draw card
			deck.Draw(players[0], 1)
		}
		if err := deck.TakeHeld(players[1], players[0], card2); err != nil {
			// Draw card
			deck.Draw(players[1], 1)
		}
	}
	hand1 := deck.Hand(players[0])
	hand2 := deck.Hand(players[1])
	//fmt.Printf("DEBUG hand1 (%d) %+v\nDEBUG hand2 (%d) %+v\n", len(hand1), hand1, len(hand2), hand2)
	if len(deck.Cards) != (len(hand1) + len(hand2)) {
		t.Errorf("oops, missing cards %s", deck.String())
	}
	//fmt.Printf("deck.String()\n%s\n", deck.String())
}
