package bot

import "testing"

type parseMessageTest struct {
	rawMsg  string
	isValid bool
	command Command
}

var parseCommandTestcases = []parseMessageTest{
	// valid
	{
		"!joke",
		true,
		Command{CmdJoke, []string{}, ""},
	},
	{
		"!JOKE",
		true,
		Command{CmdJoke, []string{}, ""},
	},
	{
		"!jOkE",
		true,
		Command{CmdJoke, []string{}, ""},
	},
	{
		"!joke bla",
		true,
		Command{CmdJoke, []string{"bla"}, ""},
	},
	{
		"!joke BLA",
		true,
		Command{CmdJoke, []string{"bla"}, ""},
	},
	{
		"!joke bla blub",
		true,
		Command{CmdJoke, []string{"bla", "blub"}, ""},
	},
	// invalid
	{
		"",
		false,
		Command{},
	},
	{
		"\n",
		false,
		Command{},
	},
	{
		"\t",
		false,
		Command{},
	},
	{
		"!!",
		false,
		Command{},
	},
}

func TestParseMessage(t *testing.T) {

	for _, testCase := range parseCommandTestcases {
		result := parseMessage(testCase.rawMsg)
		if !compareParseMessageResult(result, testCase) {
			t.Errorf("parsing %s failed", testCase.rawMsg)
		}
	}
}

func FuzzTestParseMessage(f *testing.F) {
	var seed = []string{"asd", "   ", "!§!)=$", "'ädasö", "\n", "\t"}
	for _, s := range seed {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, a string) {
		result := parseMessage(a)
		if result.IsValidCommand {
			t.Errorf("Fuzzed valid command: %s", a)
		}

	})
}

func compareParseMessageResult(result commandParseResult, expected parseMessageTest) bool {

	if result.IsValidCommand != expected.isValid {
		return false
	}

	if result.Command.Command != expected.command.Command {
		return false
	}

	if len(result.Command.Params) != len(expected.command.Params) {
		return false
	}

	if len(expected.command.Params) == 0 {
		return true
	}

	for i, expectedParam := range expected.command.Params {
		if expectedParam != result.Command.Params[i] {
			return false
		}
	}

	return true
}
