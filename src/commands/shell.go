// Package commands base command features, shell, file processor
package commands

import (
	"bufio"
	"fmt"
	"os"
	"perror"
	"strings"
)

// Shell stores current shell instance information
type Shell struct {
	PS1     string
	cmdList []string
}

// NewShell create a shell
func NewShell() *Shell {
	return &Shell{
		PS1:     ">",
		cmdList: make([]string, 0),
	}
}

// Process method to handle commands
func (sh *Shell) Process() error {
	reader := bufio.NewReader(os.Stdin)
	cmdMgr := NewManager()
	sh.prompt()
	for {
		cmdInput, _ := reader.ReadString('\n')
		cmdInput = strings.TrimRight(cmdInput, perror.NewLine)
		if perror.Empty != cmdInput {
			out, err := cmdMgr.Run(cmdInput)
			processOutput(out, err)
		} else {
			cmdMgr.Commands["help"].Run()
		}
	}
}

// propmt display command propmt PS1
func (sh *Shell) prompt() {
	fmt.Print(sh.PS1)
}
