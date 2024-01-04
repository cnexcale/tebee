package bot

import "testing"

type parseCommandTest struct {
	rawMsg   string
	expected Command
}

var parseCommandTestcases = []parseCommandTest{
	{
		"!joke",
		Command{"joke", []string{}, ""},
	},
	{
		"!JOKE",
		Command{"joke", []string{}, ""},
	},
	{
		"!jOkE",
		Command{"joke", []string{}, ""},
	},
	{
		"!joke bla",
		Command{"joke", []string{"bla"}, ""},
	},
	{
		"!joke BLA",
		Command{"joke", []string{"bla"}, ""},
	},
	{
		"!joke bla blub",
		Command{"joke", []string{"bla", "blub"}, ""},
	},
}

func TestParseCommand(t *testing.T) {

	for _, testCase := range parseCommandTestcases {
		cmd := parseCommand(testCase.rawMsg)
		if compareCommand(cmd, testCase.expected) {
			t.Errorf("parsing %s failed", testCase.rawMsg)
		}
	}
}

func compareCommand(cmd, expected Command) bool {
	if cmd.Command != expected.Command {
		return false
	}

	if len(cmd.Params) != len(expected.Params) {
		return false
	}

	for i, expectedParam := range expected.Params {
		if expectedParam != cmd.Params[i] {
			return false
		}
	}

	return true
}
