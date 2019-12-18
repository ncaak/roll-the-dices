package command

import (
	"testing"
)

func TestResolveCommandOK(t *testing.T) {
	_, err := ResolveOrError("/tira")
	if err != nil {
		t.Errorf("ERROR :: %s", err.Error())
	}
}

func TestResolveCommandKO(t *testing.T) {
	_, err := ResolveOrError("error")
	if err == nil {
		t.Error("ERROR :: No error was raised")
	}
}
