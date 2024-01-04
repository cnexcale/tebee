package bot

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	twitch "tam/src/twitch"
)

type Bot struct {
	Client *twitch.Client
}

type Command struct {
	Command       string
	Params        []string
	UserReference string
}

type commandParseResult struct {
	IsValidCommand bool
	Command        Command
}

const CommandToken = "!"

var KnownCommands = map[string]int{
	"joke": 0,
}

var CommandRegex = regexp.MustCompile(`^[a-z]($|(\w[a-z0-9]+)+$)`)

func (b Bot) Run(config twitch.ClientConfig) {
	b.Client = twitch.Init(config)

	for true {
		message := b.Client.ReceiveMessage()

		parseResult := parseMessage(message.Content)

		if !parseResult.IsValidCommand {
			fmt.Println("Not a valid command")
			continue
		}

		b.handleCommand(parseResult.Command)
	}

}

func (b Bot) handleCommand(cmd Command) {
	fmt.Println("Would handle command %s with params %#v", cmd.Command, cmd.Params)
}

func parseMessage(message string) commandParseResult {

	if !CommandRegex.MatchString(message) {
		return commandParseResult{false, Command{}}
	}

	cmd := parseCommand(message)

	if !slices.Contains(KnownCommands, cmd.Command) {
		return commandParseResult{false, Command{}}
	}

	return commandParseResult{true, cmd}
}

func parseCommand(msg string) Command {
	normalizedMsg := normalizeMessage(msg)
	matches := CommandRegex.FindStringSubmatch(normalizedMsg)
	cmd := Command{
		Command: matches[0],
	}

	for _, matchedParam := range matches[1:] {
		cmd.Params = append(cmd.Params, matchedParam)
	}

	return cmd
}

func normalizeMessage(msg string) string {
	msg = strings.TrimSpace(msg)
	msg = strings.ToLower(msg)
	return msg
}
