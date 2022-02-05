package godiscordutil

import (
	"strings"
)

func IsCommand(message string, flag string) bool {
	return strings.HasPrefix(message, flag)
}

func CommandToFunc(
	command string,
	arguments []string,
	flag string,
	cmdMap map[string]func([]string) string,
) string {
	if IsCommand(command, flag) {
		return cmdMap[command](arguments)
	}
	return ""
}
