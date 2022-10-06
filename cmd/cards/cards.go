package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

var (
	helpText = `% {app_name}(1) user manual
% R. S. Doiel
% 2022-10-01

# NAME

{app_name}

# SYSNOPSIS

{app_name} ACTION DECK [MODIFIERS]

# DESCRIPTION

{app_name} uses the command line an a data file presenting the
deck to play card games. The form of the command uses a verb
and modifier structure similar to programs a like the Go compiler
and Git source versioning system.  {app_name} doesn't apply or
maintain game ruleset it only supports managing the state of the
data file holding the representation of the deck.

{app_name} is part of the gameset collection of game programs.

## Supported actions

### Deck operations

new
: creates a new card deck data file

game
: sets the name of the card game in the data file

players
: sets the names of players in a card deck

shuffle
: shuffle the whole deck of playing cards setting

deal
: deal the cards to the players. Players list must be set previously.

deal-faceup
: deals cards to the players in a visible state

show
: shows the contents of the deck

### player's hand operations

show-hand
: shows the all contents of a player's hand (both held and visible)

reveal-card
: moves a card in player's hand from held to visible

hide-card
: moves a card in a player's hand from visible to held


# EXAMPLE

Creating a new deck called "my_deck.json".

` + "```" + `
	{app_name} new my_deck.json
` + "```" + `

Setting the name for the game using "guess-cards.json".

` + "```" + `
	{app_name} game my_deck.json "guess a my card"
` + "```" + `

Add two players.

` + "```" + `
	{app_name} players my_deck.jsonn "jane" "joe"
` + "```" + `

Shuffle cards

` + "```" + `
	{app_name} suffle my_deck.json
` + "```" + `

Deal cards

` + "```" + `
	{app_name} deal my_deck.json
` + "```" + `


`
)

func usage(appName string) string {
	return strings.ReplaceAll(helpText, "{app_name}", appName)
}

func main() {
	var (
		showHelp      bool
		showTotalOnly bool
		result        int
	)
	appName := path.Base(os.Args[0])
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showTotalOnly, "t", false, "display total only")
	flag.Parse()
	args := flag.Args()
	if showHelp || len(args) == 0 {
		fmt.Printf("%s\n", usage(appName))
		os.Exit(0)
	}

	fmt.Fprintf(os.Stderr, "%s not implemented", appName)
	os.Exit(1)
}
