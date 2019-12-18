package command

import (
	"testing"
)

func TestResolveCommandOK(t *testing.T) {
	_, err := GetValidatedCommandOrError("/tira")
	if err != nil {
		t.Errorf("ERROR :: %s", err.Error())
	}
}

func TestResolveCommandKO(t *testing.T) {
	_, err := GetValidatedCommandOrError("error")
	if err == nil {
		t.Error("ERROR :: No error was raised")
	}
}
