// Package commands with 'exit' command implementation
package commands

import "os"

// CmdExit defined arguments and related methods
type CmdExit struct {
	Command
}

// NewCmdExit new exit command instance
func NewCmdExit() *CmdExit {
	var cmd = new(CmdExit)
	cmd.Cmd = "exit"
	return cmd
}

// Help to print exit get of permission command
func (cp *CmdExit) Help() string {
	return `exit
	Terminates the process and returns control to the shell.`
}

// Parse to parse arguments
func (cp *CmdExit) Parse(argString string) error {
	cp.Command.Parse(argString)
	return nil
}

// Verify to check the provided parameters are valid or not
func (cp *CmdExit) Verify() error {
	return nil
}

// Run to execute the command and provide result
func (cp *CmdExit) Run() (string, error) {
	os.Exit(0)
	return "", nil
}
