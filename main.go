package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pushkar-anand/aka/shell"
)

var rcFilePath = ""

func init() {
	flag.Usage = usage

	flag.StringVar(&rcFilePath, "rcfile", "", "Path to the rc file")
	flag.StringVar(&rcFilePath, "f", "", "Path to the rc file")
	flag.Parse()
}

func main() {
	alias := flag.Arg(0)
	cmd := flag.Arg(1)

	if alias == "" || cmd == "" {
		fmt.Println("Usage: aka [--rcfile file] <alias> <command>")
		os.Exit(1)
	}

	var err error

	if rcFilePath == "" {
		rcFilePath, err = shell.DefaultRC()
		exitOnError(err)
	}

	err = shell.AddAlias(rcFilePath, alias, cmd)
	exitOnError(err)
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Print(`
aka is a tool for managing aliases in the shell.

Usage:
  aka [--rcfile file] <alias> <command>

Options:
	-f, --rcfile		Shell configuration file to modify. Defaults to the shell's default configuration file.
`)
}
