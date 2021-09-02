package repl

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

// REPL struct.
type REPL struct {
	commands map[string]func(string, io.Writer) error
	help     map[string]string
}

// Constructs an empty REPL.
func NewRepl() *REPL {
	panic("function not yet implemented");
}

// Combines a slice of REPLs.
func CombineRepls(repls []*REPL) (*REPL, error) {
	panic("function not yet implemented");
}

// Get commands
func (r *REPL) GetCommands() map[string]func(string, io.Writer) error {
	return r.commands
}

// Get help
func (r *REPL) GetHelp() map[string]string {
	return r.help
}

// Add a command to the set of commands.
func (r *REPL) AddCommand(trigger string, action func(string, io.Writer) error, help string) {
	panic("function not yet implemented");
}

// Return all REPL usage information as a string.
func (r *REPL) HelpString() string {
	panic("function not yet implemented");
}

// Run the REPL.
func (r *REPL) Run(c net.Conn, prompt string) {
	// Get reader and writer; stdin and stdout if no conn.
	var reader io.Reader
	var writer io.Writer
	if c == nil {
		reader = os.Stdin
		writer = os.Stdout
	} else {
		reader = c
		writer = c
	}
	scanner := bufio.NewScanner((reader))
	// Begin the repl loop
	panic("function not yet implemented");
}

// cleanInput preprocesses input to the db repl
func cleanInput(text string) string {
	panic("function not yet implemented");
}
