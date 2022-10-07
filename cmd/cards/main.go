package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/rsdoiel/gameset"
	"github.com/rsdoiel/gameset/deck"
)

func usage(appName string) string {
	return gameset.FmtText(deck.HelpText, appName, nil)
}

func main() {
	var (
		showHelp bool
	)
	appName := path.Base(os.Args[0])
	flag.BoolVar(&showHelp, "help", false, "display help")
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
	oerr := os.Stderr

	actions := deck.SetupActions()
	if fn, ok := actions[args[0]]; ok {
		err := fn(in, out, oerr, args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	fmt.Fprintf(os.Stderr, "%q is unsupported by %s", args[0], appName)
	os.Exit(1)
}
