package dice

import (
	"fmt"
	"io"
)

var ShowTotalOnly bool

func DoRollDice(in io.Reader, out io.Writer, eout io.Writer, args []string) error {
	result, err := RollDice(out, args, !ShowTotalOnly)
	if err != nil {
		return fmt.Errorf("%s\n", err)
	}
	if ShowTotalOnly {
		fmt.Fprintf(out, "%d\n", result)
	}
	return nil
}
