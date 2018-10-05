package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
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

	args := os.Args[1:]

	rcFilePath, err := getRcFilePath()

	if err != nil {
		panic(err)
	}

	if len(args) >= 2 {

		alias := args[0]
		cmd := args[1]

		ln := "alias " + alias + "='" + cmd + "' "
		println(ln)

		home := os.Getenv("HOME")

		rcFileLocation := home + rcFilePath

		bashRcFile, err := os.OpenFile(rcFileLocation, os.O_APPEND|os.O_WRONLY, 0644)
		check(err)

		defer bashRcFile.Close()

		writeToBashRc := "#ADDED BY bashAliasCreator \n" + ln + "\n"
		fmt.Println(writeToBashRc)

		if _, err = bashRcFile.WriteString(writeToBashRc); err != nil {
			panic(err)
		}
		exec.Command("/bin/bash", "-c", ln)

	} else {
		println("Incorrect Arguments. Use as bashAliasCreator alias \"command\"  ")
	}

}
