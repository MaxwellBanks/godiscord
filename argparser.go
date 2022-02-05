package godiscordutil

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type BotFunc func([]string) string

// Verifies that message is meant for this bot
func IsCommand(message string, flag string) bool {
	return strings.HasPrefix(message, flag)
}

// Verifies that message is not from this bot
func IsOwnMessage(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) bool {
	return m.Author.ID == s.State.User.ID
}

// Splits message into command (scrubbed of bot flag)
// And list of arguments
func ParseMessage(
	message string,
	flag string,
) (command string, args []string) {
	msg := strings.Split(message, " ")
	command, args = strings.TrimPrefix(msg[0], flag), msg[1:]
	return
}

// Takes command, arguments for command, command map
// Passes command and arguments through given map
// That maps command to its requisite function
func CommandToFunc(
	command string,
	arguments []string,
	cmdMap map[string]BotFunc,
) string {
	return cmdMap[command](arguments)
}
