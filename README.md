# bashAliasCreator

Easily create aliases for bash/zsh shells. No need to edit your .bashrc/.zshrc every time to add an alias.
bashAliasCreator does that for you in just one line.

## Installation

1. Make sure go is installed. If not [install go](https://golang.org/doc/install) first.
    
    `$ go  version`

2. Get the source
    
    `$ go get github.com/pushkar-anand/bashAliasCreator`
    
3. Install 

    `$ go install github.com/pushkar-anand/bashAliasCreator`
    
## Usage

`$ bashAliasCreator alias command`

**Example:** To create a alias for `ls -la` as `la` run `bashAliasCreator la "ls -la"`

Start using your alias.
