package flib

import (
	"fmt"
	"os"
)

// AbortIf will abort execution if test returns false.  If alsoDo is
// not nil, will execute that first.
func AbortIf(test func() bool, alsoDo func()) {
	if !test() {
		if alsoDo != nil {
			alsoDo()
		}

		os.Exit(1)
	}
}

// AbortIfErr will abort execution if test returns an error, after
// dumping the error to stderr.  If alsoDo is not nil, will execute
// that before aborting.
func AbortIfErr(test func() error, msg string, alsoDo func()) {
	err := test()

	if err != nil {
		dumpErr(msg, err)

		if alsoDo != nil {
			alsoDo()
		}

		os.Exit(1)
	}
}

func dumpErr(msg string, err error) {
	fmt.Fprintf(os.Stderr, "Something went horribly wrong:\n%s\n%s\n", msg, err.Error())
}
