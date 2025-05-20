package main

import (
	"errors"
	"fmt"
)

var (
	error1 = errors.New("test Error")
	error2 = errors.New("test Error")
)

func main() {

	if errors.Is(error1, error2) {
		fmt.Println("is equal")
	} else {
		fmt.Println("isn't equal")

	}

}
