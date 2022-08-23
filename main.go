package main

import (
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
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

	if terminal.IsTerminal(0) == false {
		b, _ := ioutil.ReadAll(os.Stdin)
		os.Exit(cow.Say(string(b)))
	}
}
