# godiscordutil

Package `MaxwellBanks/godiscordutil` provides a set of utility functions that may be helpful for Discord bots. 

![Unit Testing](https:/github.com/MaxwellBanks/godiscordutil/actions/workflows/push.yml/badge.svg)

## Datatypes
### Argparser
```
type BotFunc func(*sql.DB, []string) string
```
Abstract type for custom bot functions mapped to in `CommandToFunc`.
### Tablegen
```
type RawTable [][]string
```
Two dimensional list containing raw data to be loaded into ASCII table.

## Exposed Functions
### Argparser 
```
IsCommand(message string, flag string) bool
```
Checks for the activation flag in front of the message.

```
IsOwnMessage(s *discordgo.Session, m *discordgo.MessageCreate) bool
```
Prevents recursion by blacklisting self-responses.

```
ParseMessage(message string, flag string) (command string, args []string)
```
Scrubs flag from message and splits the message into a command and a list of arguments.

```
CommandToFunc(command string, arguments[string], cmdMap map[string]BotFunc, db *sql.DB) string
```
Takes in command and arguments and returns the value of the function mapped to the command.
### Tablegen
```
GenTable(data RawTable) string
```
Converts two dimensional list of strings to ASCII table suitable for Discord chat displays.
