package bot

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	twitch "tam/src/twitch"
	utils "tam/src/utils"
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
const CmdJoke = CommandToken + "joke"

var KnownCommands = []string{
	CmdJoke,
}

// var CommandRegex = regexp.MustCompile(`^` + CommandToken + `[a-z]($|(\w[a-z0-9]+)+$)`)
// var CommandRegex = regexp.MustCompile(`(^` + CommandToken + `[a-z]($|(\w[a-z0-9]+)+$)`)

// var CommandRegex = regexp.MustCompile(`^(?P<CMD>` + CommandToken + `[a-z]+)(?P<PARAMS>\w+[a-z0-9]+)*$`)

// TODO - fix, doesnt caputure inner parameters
var CommandRegex = regexp.MustCompile(`^(` + CommandToken + `[a-z]+)(\s+[a-z0-9]+)*$`)

func (b Bot) Run(config twitch.ClientConfig) {
	b.Client = twitch.Init(config)

	for {
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
	fmt.Printf("Would handle command %s with params %#v\n", cmd.Command, cmd.Params)
}

func parseMessage(message string) commandParseResult {
	normalizedMsg := normalizeMessage(message)

	if !CommandRegex.MatchString(normalizedMsg) {
		return commandParseResult{false, Command{}}
	}

	cmd := parseCommand(normalizedMsg)

	if !slices.Contains(KnownCommands, cmd.Command) {
		return commandParseResult{false, Command{}}
	}

	return commandParseResult{true, cmd}
}

func parseCommand(message string) Command {
	matches := CommandRegex.FindStringSubmatch(message)

	cmd := Command{
		Command: matches[1],
	}

	filteredParams := utils.Filter[string](matches[2:], utils.IsStringNotEmpty)

	if len(filteredParams) == 0 {
		return cmd
	}

	for _, param := range filteredParams {
		cmd.Params = append(cmd.Params, strings.TrimSpace(param))
	}

	return cmd
}

func normalizeMessage(msg string) string {
	msg = strings.TrimSpace(msg)
	msg = strings.ToLower(msg)
	return msg
}
