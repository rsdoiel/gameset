package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/rsdoiel/gameset"
	"github.com/rsdoiel/gameset/dice"
)

var (
	helpText = dice.HelpText

	showHelp bool
	showLicense bool
	showVersion bool
	inputFName string
	outputFName string
)

func main() {
	appName := path.Base(os.Args[0])
	// NOTE: the follow are set when version.go is generated
	version := gameset.Version
	releaseDate := gameset.ReleaseDate
	releaseHash := gameset.ReleaseHash
	fmtHelp := gameset.FmtHelp

	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.StringVar(&inputFName, "i", "", "read from filename")
	flag.StringVar(&outputFName, "o", "", "write to filename")
	flag.BoolVar(&dice.ShowTotalOnly, "t", false, "display only roll total")
	flag.Parse()
	args := flag.Args()

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

	if inputFName != "" {
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


	if err := dice.DoRollDice(in, out, eout, args); err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		os.Exit(1)
	}
}
