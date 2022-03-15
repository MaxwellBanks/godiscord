package godiscordutil

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

func TestIsCommand(t *testing.T) {
	realFlag := "!"
	fakeFlag := "&"
	command := "!test"
	if !IsCommand(command, realFlag) {
		t.Errorf(
			"Command validation failed, %s is not recognized as proper flag for %s",
			realFlag,
			command,
		)
	}
	if IsCommand(command, fakeFlag) {
		t.Errorf(
			"Command validation failed, %s was incorrectly recognized as proper flag for %s",
			fakeFlag,
			command,
		)
	}
}

func TestParseMessage(t *testing.T) {
	message := "!info"
	argMessage := "!info arg1 arg2"
	flag := "!"
	command, args := ParseMessage(message, flag)
	if command != "info" || len(args) != 0 {
		t.Errorf(
			"Message parse failed, expected command 'info' with no arguments."+
				"Received command %s with %d arguments",
			command,
			len(args),
		)
	}
	command, args = ParseMessage(argMessage, flag)
	if command != "info" || len(args) != 2 {
		t.Errorf(
			"Message parse failed, expected command 'info' with two arguments."+
				"Received command %s with %d arguments",
			command,
			len(args),
		)
	}
}

var dbMock *sql.DB

var CommandMap = map[string]BotFunc{
	"example": exampleFunc,
}

func exampleFunc(db *sql.DB, args []string) string {
	return fmt.Sprintf("exampleFunc %s", strings.Join(args, "|"))
}

func TestCommandToFunc(t *testing.T) {
	command := "example"
	args := []string{"arg1", "arg2"}
	result := CommandToFunc(command, args, CommandMap, dbMock)
	if result != "exampleFunc arg1|arg2" {
		t.Errorf(
			"Mapping command to function failed. Expected 'exampleFunc arg1|arg2'."+
				"Received command %s instead.", result,
		)
	}
}
