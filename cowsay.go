package main

import (
	"fmt"
	"io"
	"strings"
)

const (
	ExitCodeOK    = 0
	ExitCodeError = 1
	leftSpaceNum = 1
	rightSpaceNum = 1
)

type Cow struct {
	outStream, errStream io.Writer
}

func getMessageMaxLength(lines []string) int {
	max := 0
	for _, s := range lines {
		if max < len(s) {
			max = len(s)
		}
	}
	return max
}

func messageWithSpaces(msg string, msgMaxlen int) string {
	if len(msg) <= 0 {
		return msg + strings.Repeat(" ", msgMaxlen -1)
	}
	if len(msg) >= msgMaxlen {
		return msg
	}
	return msg + strings.Repeat(" ", len(msg) -1)
}

func createMessaage(message string) string {
	var newMessageLines []string
	messageLines := strings.Split(message, "\n")
	lastLineNum := len(messageLines) -1
	msgMaxLen := getMessageMaxLength(messageLines)

	for i, msg := range messageLines {
		if i == 0 {
			newMessageLines = append(newMessageLines, "/ " + messageWithSpaces(msg, msgMaxLen) + " \\\n")
		} else if i == lastLineNum {
			newMessageLines = append(newMessageLines, "\\ " + messageWithSpaces(msg, msgMaxLen) + " /")
		} else {
			newMessageLines = append(newMessageLines, "| " + messageWithSpaces(msg, msgMaxLen) + " |\n")
		}
	}

	return strings.Join(newMessageLines, "")
}

func createMessageBody(message string) string {
	newLineNum := strings.Count(message, "\n")
	if newLineNum < 1 {
		return "< " + message + " >"
	}
	return createMessaage(message)
}

func (c *Cow) serif(message string) string {
	msgLenMax := getMessageMaxLength(strings.Split(message, "\n"))

	msgHeader := " " + strings.Repeat("_", msgLenMax + leftSpaceNum + rightSpaceNum)
	msg := createMessageBody(message)
	msgFooter := " " + strings.Repeat("-", msgLenMax + leftSpaceNum + rightSpaceNum)

	return msgHeader + "\n" + msg + "\n" + msgFooter
}

func (c *Cow) me() string {
	return `
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||`
}

func (c *Cow) Say(message string) int {
	fmt.Fprintf(c.outStream, c.serif(message) + c.me())

	fmt.Println(c.serif(message) + c.me())

	return ExitCodeOK
}
