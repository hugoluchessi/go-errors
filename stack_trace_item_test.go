package goerrors

import (
	"runtime"
	"testing"
)

func TestNewStackTraceItem(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	expectedErrorLine := 11
	stackTraceItem := NewStackTraceItem(1)

	if stackTraceItem.PC == 0 {
		t.Error("FAILED: pc should be greater than 0")
	}

	if stackTraceItem.File != file {
		t.Errorf("FAILED: file should be '%s' got '%s'", file, stackTraceItem.File)
	}

	if stackTraceItem.Line != expectedErrorLine {
		t.Errorf("FAILED: line should be '%d' got '%d'", expectedErrorLine, stackTraceItem.Line)
	}

	t.Log("SUCCESS!")
}
