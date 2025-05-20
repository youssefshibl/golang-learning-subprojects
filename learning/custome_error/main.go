package main

import (
	"errors"
	"fmt"
)

type ArgError struct {
	arg     int
	message string
}

func (e *ArgError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &ArgError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	_, err := f(42)

	var ae *ArgError
	if errors.As(err, &ae) {
		fmt.Println("Matched ArgError:", ae)
	} else {
		fmt.Println("err doesn't match ArgError")
	}
}
