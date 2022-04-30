package shell

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

var shellRCMap = map[string]string{
	"bash": ".bashrc",
	"zsh":  ".zshrc",
	"ksh":  ".kshrc",
	"tcsh": ".tcshrc",
	"csh":  ".cshrc",
}

//DefaultRC returns the default shell rc file for the current shell
func DefaultRC() (string, error) {
	shellName := os.Getenv("SHELL")
	if shellName == "" {
		return "", fmt.Errorf("failed to identify shell. SHELL environment variable not set")
	}

	shellName = cleanShellName(shellName)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to identify home directory. %v", err)
	}

	rcFile, ok := shellRCMap[shellName]
	if !ok {
		return "", fmt.Errorf("failed to identify rc file for shell %s", shellName)
	}

	return homeDir + "/" + rcFile, nil
}

//AddAlias adds an alias to the current shell rc file
func AddAlias(rcFilePath, alias, cmd string) error {
	file, err := os.OpenFile(rcFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open rc file: %w", err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Printf("failed to close rc file: %v", err)
		}
	}(file)

	return writer(file, alias, cmd)
}

func writer(w io.Writer, alias, cmd string) error {
	cmdLine := aliasCmdLine(alias, cmd)

	_, err := fmt.Fprintln(w, cmdLine)
	if err != nil {
		return fmt.Errorf("failed to write to rc file: %w", err)
	}

	return nil
}

func aliasCmdLine(alias, cmd string) string {
	return fmt.Sprintf("alias %s='%s' # Added by aka", alias, cmd)
}

var cleanupRegex = regexp.MustCompile(`/[A-Za-z/]*/`)

func cleanShellName(name string) string {
	return cleanupRegex.ReplaceAllString(name, "")
}
