package errors

import (
	"runtime"
	"testing"
)

func TestNewStackItem(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	expectedErrorLine := 11
	stackItem := newStackItem(1)

	if stackItem.pc == 0 {
		t.Error("FAILED: pc should be greater than 0")
	}

	if stackItem.file != file {
		t.Errorf("FAILED: file should be '%s' got '%s'", file, stackItem.file)
	}

	if stackItem.line != expectedErrorLine {
		t.Errorf("FAILED: line should be '%d' got '%d'", expectedErrorLine, stackItem.line)
	}

	t.Log("SUCCESS!")
}
