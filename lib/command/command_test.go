package command

import (
	"testing"
)

/*
 * Mocks
 */
type mockRequest struct {}
func (r mockRequest) BasicReply (a int, b int, c string) {}
func (r mockRequest) KeyboardReply (a int, b int) {}
func (r mockRequest) MarkdownReply (a int, b int, c string) {}

type mockSource struct {
	Text string
}
func (s mockSource) GetChatId () int {return 0}
func (s mockSource) GetCommand () string {return s.Text}
func (s mockSource) GetReplyId () int {return 0}

func mockInput(command string) (mock mockSource) {
	mock.Text = "/" + command
	return
}

/*
 * Command validation tests (only if the command is recognized as valid)
 */
func validateTest(t *testing.T, command string) {
	_, err := GetValidatedCommandOrError(mockInput(command))
	if err != nil {
		t.Errorf("ERROR :: %s", err.Error())
	}
}

func TestValidateTiraCommandOK(t *testing.T) {
	validateTest(t, "tira")
}

func TestValidateTCommandOK(t *testing.T) {
	validateTest(t, "t")
}

func TestValidateVCommandOK(t *testing.T) {
	validateTest(t, "v")
}

func TestValidateDvCommandOK(t *testing.T) {
	validateTest(t, "dv")
}

func TestValidateAgrupaCommandOK(t *testing.T) {
	validateTest(t, "agrupa")
}

func TestValidateAyudaCommandOK(t *testing.T) {
	validateTest(t, "ayuda")
}

func TestValidateRepiteCommandOK(t *testing.T) {
	validateTest(t, "repite")
}

func TestUnknownCommandOK(t *testing.T) {
	_, err := GetValidatedCommandOrError(mockInput("test"))
	if err == nil {
		t.Error("ERROR :: No error was raised")
	}
}

/*
 * Command execution tests
 */
func TestCommandAgrupaOK(t *testing.T) {
	command := getBaseCommand("agrupa", "1d8+1d6")
	command.source = mockSource{}
	err := command.Run(mockRequest{})
	if err != nil {
		t.Errorf("ERROR :: %s", err.Error())
	}
}

func TestCommandAgrupaKO(t *testing.T) {
	command := getBaseCommand("agrupa", "")
	command.source = mockSource{}
	err := command.Run(mockRequest{})
	if err == nil {
		t.Error("ERROR :: Failed to catch non valid input")
	}
}
