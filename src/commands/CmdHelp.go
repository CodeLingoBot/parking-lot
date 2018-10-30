// Package commands with 'help' command implementation
package commands

import (
	"fmt"
)

// CmdHelp defined arguments and related methods
type CmdHelp struct {
	Command
}

// NewCmdHelp new help command instance
func NewCmdHelp() *CmdHelp {
	var cmd = new(CmdHelp)
	cmd.Cmd = "help"
	return cmd
}

// Help to print help get of permission command
func (cp *CmdHelp) Help() string {
	return `help
	Shows command's help`
}

// Parse to parse arguments
func (cp *CmdHelp) Parse(argString string) error {
	cp.Command.Parse(argString)
	return nil
}

// Verify to check the provided parameters are valid or not
func (cp *CmdHelp) Verify() error {
	return nil
}

// Run to execute the command and provide result
func (cp *CmdHelp) Run() (string, error) {
	for _, cmd := range mgrCmd.Commands {
		fmt.Println(cmd.Help())
	}
	return "", nil
}
