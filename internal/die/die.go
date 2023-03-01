package die

import (
	"errors"
	"fmt"
	"os"
)

// Dief exits program execution with the given message
func Dief(x string, xs ...any) {
	fmt.Printf(x, xs)
	os.Exit(1)
}

// Must dies if error is not nil
func Must(err error) {
	if err != nil {
		Dief("fatal: %v\n", errors.Unwrap(err))
	}
}

// AllMust dies if any error in the group
func AllMust(errs ...error) {
	for _, err := range errs {
		Must(err)
	}
}
