package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getRcFilePath() (string, error) {
	shell := os.Getenv("SHELL")

	switch {
	case strings.Contains(shell, "bash"):
		return "/.bashrc", nil
	case strings.Contains(shell, "zsh"):
		return "/.zshrc", nil
	default:
		var errMsg = fmt.Sprintf("Unsupported shell %s", shell)
		return "", errors.New(errMsg)
	}
}

func main() {

	rcFilePath, err := getRcFilePath()
	check(err)

	args := os.Args[1:]
	if len(args) >= 2 {

		alias := args[0]
		cmd := args[1]

		ln := "alias " + alias + "='" + cmd + "' "
		println(ln)

		home := os.Getenv("HOME")

		rcFileLocation := home + rcFilePath

		rcFile, err := os.OpenFile(rcFileLocation, os.O_APPEND|os.O_WRONLY, 0644)
		check(err)
		defer rcFile.Close()

		writeToRc := "#ADDED BY bashAliasCreator \n" + ln + "\n"

		if _, err = rcFile.WriteString(writeToRc); err != nil {
			panic(err)
		}

	} else {
		println("Incorrect Arguments. Use as bashAliasCreator alias \"command\"  ")
	}

}
