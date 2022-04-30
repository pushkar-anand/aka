# aka

Easily create aliases for linux shells. No need to edit your shell startup scripts every time to add an alias.
aka does that for you in just one line.

## Table of Contents

- [Usage](#usage)
- [Supported Shells](#supported-shells)
- [Installation](#installation)

## Usage

```shell
aka [--rcfile file] <alias> <command>

Options:
	-f, --rcfile		Shell configuration file to modify. Defaults to the shell's default configuration file.
```

### Supported Shells

- BASH
- ZSH
- TCSH
- KSH

## Installation

### Download pre-built binary (Linux or macOS)

[Release Downloads](https://github.com/pushkar-anand/aka/releases)

### With Go toolchain

```shell
$ # Go 1.15 and below
$ go get -u github.com/pushkar-anand/aka
$ # Go 1.16+
$ go install github.com/pushkar-anand/aka@latest
```