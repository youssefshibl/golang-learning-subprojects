package main

import (
	"taskTrackes/src"
)

func main() {

	rootCmd := src.Root()
	rootCmd.Execute()
}
