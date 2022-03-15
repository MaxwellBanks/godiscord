package godiscordutil

import (
	"database/sql"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// BotFunc is an abstract type for functions that serve a Discord bot
type BotFunc func(*sql.DB, []string) string

// IsCommand verifies that a given message is meant for this bot
func IsCommand(message string, flag string) bool {
	return strings.HasPrefix(message, flag)
}

// IsOwnMessage verifies that a given message is not from this bot
func IsOwnMessage(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) bool {
	return m.Author.ID == s.State.User.ID
}

// ParseMessage splits message into command (scrubbed of bot flag)
// And list of arguments
func ParseMessage(
	message string,
	flag string,
) (command string, args []string) {
	msg := strings.Split(message, " ")
	command, args = strings.TrimPrefix(msg[0], flag), msg[1:]
	return
}

// CommandToFunc passes command and argument parameters through
// Map parameter, returning the resulting function from the map
func CommandToFunc(
	command string,
	arguments []string,
	cmdMap map[string]BotFunc,
	db *sql.DB,
) string {
	result := ""
	if _, err := cmdMap[command]; err {
		result = cmdMap[command](db, arguments)
	}
	return result
}
