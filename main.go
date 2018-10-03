package main

import (
	"fmt"
	"os"
	"os/exec"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	args := os.Args[1:]

	if len(args) >= 2 {

		alias := args[0]
		cmd := args[1]

		ln := "alias " + alias + "='" + cmd + "' "
		println("append: " + ln)

		home := os.Getenv("HOME")
		println("HOME: " + home)

		bashRcFileLocation := home + "/.bashrc"

		bashRcFile, err := os.OpenFile(bashRcFileLocation, os.O_APPEND|os.O_WRONLY, 0644)
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
