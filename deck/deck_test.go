package deck

import "testing"

func TestDeck(t *testing.T) {
	deck := MakeStandardDeck()
	if deck == nil {
		t.Errorf("expected a standard deck to be made")
		t.FailNow()
	}
	if len(deck.Cards) != 52 {
		t.Errorf("expected 52 cards in deck.Cards")
	}
}

func TestMethods(t *testing.T) {
	t.Errorf("not implemented, need to write test for various Deck methods")
}
