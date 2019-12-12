package command

import (
	"testing"
)


func TestValidateCommandOK(t *testing.T) {
	_, err := getValidatedCommandOrError("/tira")
	if err != nil {
		t.Errorf("ERROR :: %s", err.Error())
	}
}

func TestValidateCommandKO(t *testing.T) {
	_, err := getValidatedCommandOrError("error")
	if err == nil {
		t.Error("ERROR :: No error was raised")
	}
}
