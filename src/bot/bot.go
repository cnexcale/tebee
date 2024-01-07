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

// TODO - fix, doesnt caputure inner parameters for > 1 param
// var CommandRegex = regexp.MustCompile(`^(` + CommandToken + `[a-z]+)(\s+[a-z0-9]+)*$`)

var CommandRegex = regexp.MustCompile(`^(` + CommandToken + `[a-z]+)($|\s+[a-z0-9]+)`)

func (b Bot) Run(config twitch.ClientConfig) {
	b.Client = twitch.Init(config)

	for {
		message := b.Client.ReceiveMessage()

		// use concurrency once bot is fully implemented
		// go b.handleMessage(message)

		b.handleMessage(message)
	}
}

func (b Bot) handleMessage(message twitch.ChatMessage) {
	parseResult := parseMessage(message.Content)

	if !parseResult.IsValidCommand {
		fmt.Println("Not a valid command")
		return
	}

	b.HandleCommand(parseResult.Command)
}

func (b Bot) HandleCommand(cmd Command) {
	// fmt.Printf("Would handle command %s with params %#v\n", cmd.Command, cmd.Params)
	if cmd.Command == CmdJoke {
		handleJokeCommand(*b.Client, cmd)
	}
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

	// only consider first param, discard the rest
	cmd.Params = append(cmd.Params, strings.TrimSpace(filteredParams[0]))

	return cmd
}

func normalizeMessage(msg string) string {
	msg = strings.TrimSpace(msg)
	msg = strings.ToLower(msg)
	return msg
}
