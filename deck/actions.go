package deck

import (
	"fmt"
	"io"
	"os"

	"github.com/rsdoiel/gameset"
)

func doNew(in io.Reader, out io.Writer, eout io.Writer, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("missing filename for deck")
	}
	for _, deckName := range args[1:] {
		if _, err := os.Stat(deckName); err == nil {
			return fmt.Errorf("%s already exists", deckName)
		}
		deck := MakeStandardDeck()
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
