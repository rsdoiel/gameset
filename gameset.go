package gameset

import (
	"fmt"
	"io"
	"strings"
)

// Runner is the function signature for running the various gameset
// tools. E.g.
//
// ```
//
//	func Help(in io.Reader, out io.Writer, eout io.Writer, args []string) error {
//	 // ... do stuff ...
//	 return nil
//	}
//
// // ... do stuff ...
//
// actions["help"] = helpFunction
//
// // ... do stuff ...
//
//	if fn, ok := actions["help"]; ok {
//	 if err  := fn(os.Stdin, os.Stdout, os.Stderr, os.Argv); err != nil {
//	     //  ... handle error
//	 }
//	}
//
// ```
type Runner func(io.Reader, io.Writer, io.Writer, []string) error

// FmtText takes a source string that is marked up with curly
// brackets with value name and replaces them returning revised string.
//
// ```
// src := `Hello {app_name}, my name is {my_name}`
// m := map[string]string{ "my_name": "Jane Doe" }
// message := FmtText(src, path.Base(os.Argv[0]), m)
// fmt.Printf("messge: %s\n", message)
// ```
func FmtText(src string, appName string, m map[string]string) string {
	docs := strings.ReplaceAll(src, "{app_name}", appName)
	for k, v := range m {
		tag := fmt.Sprintf("{%s}", k)
		docs = strings.ReplaceAll(docs, tag, v)
	}
	return docs
}
