package command

import (
	"testing"
)

/*
 * Mocks
 */
type mockRequest struct{}

func (r mockRequest) BasicReply(a int, b int, c string)    {}
func (r mockRequest) DiceKeyboardReply(a int, b int)           {}
func (r mockRequest) KeyboardReply(int, int, map[string]string) {}
func (r mockRequest) EditKeyboardReply(a int, b int, c string) {}
func (r mockRequest) MarkdownReply(a int, b int, c string) {}

type mockSource struct {
	Text string
}

func (s mockSource) GetChatId() int     { return 0 }
func (s mockSource) GetCommand() string { return s.Text }
func (s mockSource) GetReplyId() int    { return 0 }

func mockInput(command string) (mock mockSource) {
	mock.Text = "/" + command
	return
}

func mockCommand(command string, argument string) error {
	cmd := getBaseCommand(command, argument)
	cmd.source = mockSource{}
	return cmd.Run(mockRequest{})
}

func mockValidation(t *testing.T, command string) {
	_, err := GetValidatedCommandOrError(mockInput(command))
	if err != nil {
		t.Errorf("ERROR :: command %s : %s", command, err.Error())
	}
}

/*
 * Command validation tests (only if the command is recognized as valid)
 */
func TestValidCommandOK(t *testing.T) {
	mockValidation(t, "tira")
	mockValidation(t, "t")
	mockValidation(t, "v")
	mockValidation(t, "dv")
	mockValidation(t, "agrupa")
	mockValidation(t, "ayuda")
	mockValidation(t, "repite")
	mockValidation(t, "pj")
}

func TestValidCommandKO(t *testing.T) {
	_, err := GetValidatedCommandOrError(mockInput("test"))
	if err == nil {
		t.Error("ERROR :: No error was raised")
	}
}

/*
 * Command execution tests
 */
func TestCommandAgrupaOK(t *testing.T) {
	err := mockCommand("agrupa", "1d8+1d6")
	if err != nil {
		t.Errorf("ERROR :: %s", err.Error())
	}
}

func TestCommandAgrupaKO(t *testing.T) {
	err := mockCommand("agrupa", "")
	if err == nil {
		t.Error("ERROR :: Failed to catch non valid input")
	}
}

func TestCommandRepiteOK(t *testing.T) {
	err := mockCommand("repite", "2 1d20")
	if err != nil {
		t.Errorf("ERROR :: %s", err.Error())
	}
}

func TestCommandRepiteKO(t *testing.T) {
	err := mockCommand("repite", "1d20")
	if err == nil {
		t.Error("ERROR :: Failed to catch non valid input")
	}
}

func TestCommandAyudaOK(t *testing.T) {
	err := mockCommand("ayuda", "tira")
	if err != nil {
		t.Errorf("ERROR :: %s", err.Error())
	}
}

func TestCommandAyudaKO(t *testing.T) {
	help := getHelp("test")
	if help != DEFAULT_HELP {
		t.Error("ERROR :: Test command validated")
	}
}
