package flib

import (
	"os"
	"fmt"
)

func abortIf(test func() bool, alsoDo func()) {
	if !test() {
		if alsoDo != nil {
			alsoDo()
		}

		os.Exit(1)
	}
}

func abortIfErr(test func() error, msg string, alsoDo func()) {
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
	fmt.Printf("Something went horribly wrong:\n%s\n%s\n", msg, err.Error())
}
