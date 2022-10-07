package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/rsdoiel/gameset"
	"github.com/rsdoiel/gameset/dice"
)

func usage(appName string) string {
	return gameset.FmtText(dice.HelpText, appName, nil)
}

func main() {
	var (
		showHelp bool
	)
	appName := path.Base(os.Args[0])
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&dice.ShowTotalOnly, "t", false, "display only roll total")
	flag.Parse()
	args := flag.Args()
	if showHelp || len(args) == 0 {
		fmt.Printf("%s\n", usage(appName))
		os.Exit(0)
	}
	if len(args) == 0 {
		fmt.Printf("%s\n", usage(appName))
		os.Exit(1)
	}

	in := os.Stdin
	out := os.Stdout
	eout := os.Stderr

	if err := dice.DoRollDice(in, out, eout, args); err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		os.Exit(1)
	}
}
