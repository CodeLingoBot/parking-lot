package commands

import (
	"perror"
	"strings"
)

// ICommand for base command's required behaviour
type ICommand interface {
	Help() string
	GetName() string
	Parse(string) error
	Verify() error
	Run() (string, error)
	Clear()
}

// Command object structure
type Command struct {
	Cmd,
	InputArgs string
	OutPut string
	Args   []string
}

// NewCommand to create command instance
func NewCommand() *Command {
	var cmd = new(Command)
	return cmd
}

// Help to show usage
func (c *Command) Help() string {
	return "No help found"
}

// GetName to get the command name
func (c *Command) GetName() string {
	return c.Cmd
}

// Parse to help command to parse arguments from input string
func (c *Command) Parse(argString string) error {
	c.InputArgs = argString
	c.Args = strings.Split(argString, perror.Space)
	return nil
}

// Verify the provided Arguments
func (c *Command) Verify() error {
	return nil
}

// Run the command with arguments
func (c *Command) Run() (string, error) {
	return c.OutPut, nil
}
