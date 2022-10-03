package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/rsdoiel/gameset/dice"
)

var (
	helpText = `% {app_name}(1) user manual
% R. S. Doiel
% 2022-10-01

# NAME

{app_name}

# SYSNOPSIS

{app_name} DICE_ROLL_EXPR [DICE_ROLL_EXPR ...]

# DESCRIPTION

{app_name} parses the dice roll expressions and
models the rolling of the dice described by the expressions.
Dice roll expressions are forms by an integer (number of dice)
the letter "d" followed by an integer (number of sides) and optionally
follow by a plus or minus and integer which is treated as an offset.

{app_name} is part of the gameset collection of game programs.

# EXAMPLE

Rolling one twenty sided dice.

` + "```" + `
	{app_name} 1d20
` + "```" + `

Rolling a three four sided dices plus two.

` + "```" + `
	{app_name} 3d4+2
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

	result, err := dice.RollDice(args, !showTotalOnly)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	if showTotalOnly {
		fmt.Printf("%d\n", result)
	}
}
