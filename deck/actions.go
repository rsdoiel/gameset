package deck

import (
	"fmt"
	"io"

	"github.com/rsdoiel/gameset"
)

func doNew(in io.Reader, out io.Writer, eout io.Writer, args []string) error {
	return fmt.Errorf("doNew() not implemented")
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
