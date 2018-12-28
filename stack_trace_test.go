package goerrors

import (
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func TestNewStackTrace(t *testing.T) {
	st := NewStackTrace()

	if st == nil {
		t.Error("FAILED: StackTrace must not be nil.")
	}

	if st.Items == nil {
		t.Error("FAILED: StackTrace Items must not be an empty slice of items.")
	}

	t.Log("SUCCESS!")
}

func TestAddStackItem(t *testing.T) {
	st := NewStackTrace()
	st.AddStackItem()

	_, file, expectedErrorLine, _ := runtime.Caller(1)

	if st.Items[0].PC == 0 {
		t.Error("FAILED: pc should be greater than 0")
	}

	if st.Items[0].File != file {
		t.Errorf("FAILED: file should be '%s' got '%s'", file, st.Items[0].File)
	}

	if st.Items[0].Line != expectedErrorLine {
		t.Errorf("FAILED: line should be '%d' got '%d'", expectedErrorLine, st.Items[0].Line)
	}

	t.Log("SUCCESS!")
}

func TestStackTraceString(t *testing.T) {
	st := NewStackTrace()
	st.AddStackItem()

	stackTraceMessage := st.String()

	_, file, line, _ := runtime.Caller(1)

	if !strings.Contains(stackTraceMessage, file) {
		t.Errorf("FAILED: stack trace item file not found")
	}

	if !strings.Contains(stackTraceMessage, strconv.Itoa(line)) {
		t.Errorf("FAILED: stack trace item line not found")
	}

	t.Log("SUCCESS!")
}
