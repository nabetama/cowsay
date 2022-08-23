package main

import (
	"os"
)

const name = "fortune"

const version = "0.0.1"

var revision = "HEAD"

func main() {
	cow := &Cow{
		outStream: os.Stdout,
		errStream: os.Stdout,
	}

	os.Exit(cow.Say(os.Args[1]))
}