package main

import (
	"bytes"
	"os"
	"testing"
)

func TestOutputSingleLineSerif(t *testing.T) {
	arg := "hello"
	expect := ` _______
< hello >
 -------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||`

	buf := &bytes.Buffer{}

	cow := &Cow{
		outStream: buf,
		errStream: os.Stderr,
	}

	_ = cow.Say(arg)

	if buf.String() != expect {
		t.Errorf("expect: %s, but: %s", expect, buf.String())
	}
}

func TestOutputMultipleLineSerif(t *testing.T) {
	arg := "hello\nworld"
	expect := ` _______
/ hello \
\ world /
 -------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||`

	buf := &bytes.Buffer{}

	cow := &Cow{
		outStream: buf,
		errStream: os.Stderr,
	}

	_ = cow.Say(arg)

	if buf.String() != expect {
		t.Errorf("expect: %s, but: %s", expect, buf.String())
	}
}

func TestOutputMultipleLineSerif2(t *testing.T) {
	arg := "hello\nworld\nfoobarbaz"
	expect := ` ___________
/ hello     \
| world     |
\ foobarbaz /
 -----------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||`

	buf := &bytes.Buffer{}

	cow := &Cow{
		outStream: buf,
		errStream: os.Stderr,
	}

	_ = cow.Say(arg)

	if buf.String() != expect {
		t.Errorf("expect: %s, but: %s", expect, buf.String())
	}
}

func TestOutputMultipleLineSerif3(t *testing.T) {
	arg := "hello\nworld\nfoobarbaz\n"
	expect := ` ___________
/ hello     \
| world     |
| foobarbaz |
\          /
 -----------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||`

	buf := &bytes.Buffer{}

	cow := &Cow{
		outStream: buf,
		errStream: os.Stderr,
	}

	_ = cow.Say(arg)

	if buf.String() != expect {
		t.Errorf("expect: %s, but: %s", expect, buf.String())
	}
}
