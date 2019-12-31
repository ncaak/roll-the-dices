package command

import (
	"testing"
	"github.com/ncaak/roll-the-dices/lib/request/structs"
)

func mockResult(command string) (mock structs.Result) {
	mock.Message.Text = "/" + command
	return
}

/*
 * Command validation tests (only if the command is recognized as valid)
 */
func validateTest(t *testing.T, command string) {
	_, err := GetValidatedCommandOrError(mockResult(command))
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
	_, err := GetValidatedCommandOrError(mockResult("test"))
	if err == nil {
		t.Error("ERROR :: No error was raised")
	}
}

/*
 * Command execution tests
 */

