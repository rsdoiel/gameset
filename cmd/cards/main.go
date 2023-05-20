package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/rsdoiel/gameset"
	"github.com/rsdoiel/gameset/deck"
)
var (
	helpText = deck.HelpText

	showHelp bool
	showLicense bool
	showVersion bool
	inputFName string
	outputFName string
)


func main() {
	appName := path.Base(os.Args[0])
	// NOTE: the following are set when version.go is generated
	version := gameset.Version
	releaseDate := gameset.ReleaseDate
	releaseHash := gameset.ReleaseHash
	fmtHelp := gameset.FmtHelp

	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		inputFName = args[0]
	}
	if len(args) > 1 {
		outputFName = args[1]
	}

	var err error

	in := os.Stdin
	out := os.Stdout
	eout := os.Stderr


	if showHelp {
		fmt.Fprintf(out, "%s\n", fmtHelp(helpText, appName, version, releaseDate, releaseHash))
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintf(out, "%s\n", gameset.LicenseText)
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintf(out, "%s %s %s\n", appName, version, releaseHash)
		os.Exit(0)
	}

	if len(args) == 0 {
		fmt.Printf("%s\n", fmtHelp(helpText, appName, version, releaseDate, releaseHash))
		os.Exit(1)
	}

	if inputFName  != "" {
		in, err = os.Open(inputFName)
		if err != nil {
			fmt.Fprintf(eout, "%s\n", err)
			os.Exit(1)
		}
		defer in.Close()
	}

	if outputFName != "" {
		out, err = os.Create(outputFName)
		if err != nil {
			fmt.Fprintf(eout, "%s\n", err)
			os.Exit(1)
		}
		defer out.Close()
	}

	actions := deck.SetupActions()
	if fn, ok := actions[args[0]]; ok {
		err = fn(in, out, eout, args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	fmt.Fprintf(eout, "%q is unsupported by %s", args[0], appName)
	os.Exit(1)
}
