package deck

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rsdoiel/gameset"
)

func doNew(in io.Reader, out io.Writer, eout io.Writer, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("missing filename for deck")
	}
	deckType := "standard"
	for _, deckName := range args[1:] {
		if strings.Contains(deckName, "=") {
			parts := strings.SplitN(deckName, "=", 2)
			deckType, deckName = strings.ToLower(strings.TrimSpace(parts[0])), strings.TrimSpace(parts[1])
		}
		if _, err := os.Stat(deckName); err == nil {
			return fmt.Errorf("%s already exists", deckName)
		}
		var deck *Deck
		switch deckType {
		case "standard":
			deck = NewStandardDeck()
		case "hwatu":
			deck = NewHwatuDeck()
		case "sakura":
			deck = NewSakuraDeck()
		default:
			return fmt.Errorf("unsupported deck type %q", deckType)
		}
		fmt.Fprintf(out, "%s\n", deck.String())
		deck.Save(deckName)
	}
	return nil
}

func doReset(in io.Reader, out io.Writer, eout io.Writer, args []string) error {
	return fmt.Errorf("doReset() not implemented")
}

func SetupActions() map[string]gameset.Runner {
	actions := map[string]gameset.Runner{}
	actions["new"] = doNew
	actions["reset"] = doReset
	return actions
}
