package goerrors

import (
	"errors"
	"strings"
	"testing"
)

func TestNewError(t *testing.T) {
	rootError := errors.New("some error here")

	ew := NewError(rootError).(*errorWrapper)

	if ew.rootError != rootError {
		t.Errorf("FAILED: Wrong Root error, expected '%s' got '%s'", rootError.Error(), ew.Error())
	}

	if len(ew.errorStackTrace.Items) != 1 {
		t.Errorf("FAILED: stack size should be 1 got '%d'", len(ew.errorStackTrace.Items))
	}

	t.Log("SUCCESS!")
}
func TestNewErrorWithState(t *testing.T) {
	rootError := errors.New("some error here")

	state := map[string]interface{}{
		"some":      "value",
		"int_value": 1234,
	}

	ew := NewErrorWithState(rootError, state).(*errorWrapper)

	if ew.rootError != rootError {
		t.Errorf("FAILED: Wrong Root error, expected '%s' got '%s'", rootError.Error(), ew.Error())
	}

	if len(ew.errorStackTrace.Items) != 1 {
		t.Errorf("FAILED: stack size should be 1 got '%d'", len(ew.errorStackTrace.Items))
	}

	for key, value := range state {
		v, _ := ew.errorState[key]

		if v != value {
			t.Errorf("FAILED: Key '%s', hash incorrect value, expected '%s' got '%s'.", key, value, v)
		}
	}

	t.Log("SUCCESS!")
}

func TestNewErrorWithInvalidError(t *testing.T) {
	rootError := errors.New("some error here")

	ew := WrapError(rootError)

	castEw, ok := ew.(*errorWrapper)

	if !ok {
		t.Error("FAILED: Unable to aseert ew to *errorWrapper")
	}

	if castEw.rootError != rootError {
		t.Errorf("FAILED: Wrong Root error, expected '%s' got '%s'", rootError.Error(), castEw.Error())
	}

	if len(castEw.errorStackTrace.Items) != 1 {
		t.Errorf("FAILED: stack size should be 1 got '%d'", len(castEw.errorStackTrace.Items))
	}

	t.Log("SUCCESS!")
}

func TestWrapError(t *testing.T) {
	rootError := errors.New("some error here")
	fileNameFn1 := "error_wrapper_test_fn1.go"
	lineFn1 := 4

	fileNameFn2 := "error_wrapper_test_fn2.go"
	lineFn2 := 5

	fileNameFn3 := "error_wrapper_test_fn3.go"
	lineFn3 := 6

	state := map[string]interface{}{
		"some":      "value",
		"int_value": 1234,
	}

	ew := f3(rootError, state).(*errorWrapper)

	if ew.rootError != rootError {
		t.Errorf("FAILED: Wrong Root error, expected '%s' got '%s'", rootError.Error(), ew.Error())
	}

	if len(ew.errorStackTrace.Items) != 3 {
		t.Errorf("FAILED: stack size should be 3 got '%d'", len(ew.errorStackTrace.Items))
	}

	if !strings.HasSuffix(ew.errorStackTrace.Items[0].File, fileNameFn1) {
		t.Errorf("FAILED: first stack item file name should be '%s' got '%s'", fileNameFn1, ew.errorStackTrace.Items[0].File)
	}

	if ew.errorStackTrace.Items[0].Line != lineFn1 {
		t.Errorf("FAILED: first stack item line should be '%d' got '%d'", lineFn1, ew.errorStackTrace.Items[0].Line)
	}

	if !strings.HasSuffix(ew.errorStackTrace.Items[1].File, fileNameFn2) {
		t.Errorf("FAILED: second stack item file name should be '%s' got '%s'", fileNameFn2, ew.errorStackTrace.Items[1].File)
	}

	if ew.errorStackTrace.Items[1].Line != lineFn2 {
		t.Errorf("FAILED: second stack item line should be '%d' got '%d'", lineFn2, ew.errorStackTrace.Items[1].Line)
	}

	if !strings.HasSuffix(ew.errorStackTrace.Items[2].File, fileNameFn3) {
		t.Errorf("FAILED: third stack item file name should be '%s' got '%s'", fileNameFn3, ew.errorStackTrace.Items[2].File)
	}

	if ew.errorStackTrace.Items[2].Line != lineFn3 {
		t.Errorf("FAILED: third stack item line should be '%d' got '%d'", lineFn3, ew.errorStackTrace.Items[2].Line)
	}

	for key, value := range state {
		v, _ := ew.errorState[key]

		if v != value {
			t.Errorf("FAILED: Key '%s', hash incorrect value, expected '%s' got '%s'.", key, value, v)
		}
	}

	t.Log("SUCCESS!")
}

func TestError(t *testing.T) {
	rootError := errors.New("some error here")

	ew := NewError(rootError)

	if ew.Error() != rootError.Error() {
		t.Errorf("FAILED: Wrong Root error, expected '%s' got '%s'", rootError.Error(), ew.Error())
	}

	t.Log("SUCCESS!")
}

func TestErrorString(t *testing.T) {
	rootError := errors.New("some error here")
	state := map[string]interface{}{
		"some":      "value",
		"int_value": 1234,
	}

	ew := f3(rootError, state).(*errorWrapper)

	traceMessage := ew.String()

	expectedErrorHeader := "Error: 'some error here'"

	expectedStateHeader := "State:"
	expectedStateItem1Value := "some: value"
	expectedStateItem2Value := "int_value: 1234"

	expectedStackHeader := "Stack Trace:"
	expectedFileNameAndLine1 := "error_wrapper_test_fn1.go:4"
	expectedFileNameAndLine2 := "error_wrapper_test_fn2.go:5"
	expectedFileNameAndLine3 := "error_wrapper_test_fn3.go:6"

	if !strings.Contains(traceMessage, expectedErrorHeader) {
		t.Errorf("FAILED: Error header not found")
	}

	if !strings.Contains(traceMessage, expectedStateHeader) {
		t.Errorf("FAILED: state header not found")
	}

	if !strings.Contains(traceMessage, expectedStateItem1Value) {
		t.Errorf("FAILED: state item 1 not found")
	}

	if !strings.Contains(traceMessage, expectedStateItem2Value) {
		t.Errorf("FAILED: state item 2 not found")
	}

	if !strings.Contains(traceMessage, expectedStackHeader) {
		t.Errorf("FAILED: stack trace header not found")
	}

	if !strings.Contains(traceMessage, expectedFileNameAndLine1) {
		t.Errorf("FAILED: stack item 1 not found")
	}

	if !strings.Contains(traceMessage, expectedFileNameAndLine2) {
		t.Errorf("FAILED: stack item 2 not found")
	}

	if !strings.Contains(traceMessage, expectedFileNameAndLine3) {
		t.Errorf("FAILED: stack item 3 not found")
	}

	t.Log("SUCCESS!")
}

func TestRootWithErrorWrapper(t *testing.T) {
	err := errors.New("some error here")

	ew := NewError(err).(*errorWrapper)

	rerr := RootError(ew)

	if rerr != err {
		t.Errorf("FAILED: Wrong Root error, expected '%s' got '%s'", err.Error(), rerr.Error())
	}

	t.Log("SUCCESS!")
}

func TestRootWithError(t *testing.T) {
	err := errors.New("some error here")

	rerr := RootError(err)

	if rerr != err {
		t.Errorf("FAILED: Wrong Root error, expected '%s' got '%s'", err.Error(), rerr.Error())
	}

	t.Log("SUCCESS!")
}

func TestBuildStacktraceWithErrorWrapper(t *testing.T) {
	err := errors.New("some error here")

	ew := NewError(err).(*errorWrapper)

	msg := BuildStackTrace(ew)

	expectedErrorHeader := "Error: 'some error here'"
	expectedStateHeader := "State:"
	expectedStackHeader := "Stack Trace:"

	if !strings.Contains(msg, expectedErrorHeader) {
		t.Errorf("FAILED: Error header not found")
	}

	if !strings.Contains(msg, expectedStateHeader) {
		t.Errorf("FAILED: state header not found")
	}

	if !strings.Contains(msg, expectedStackHeader) {
		t.Errorf("FAILED: stack trace header not found")
	}

	t.Log("SUCCESS!")
}

func TestBuildStacktraceWithError(t *testing.T) {
	err := errors.New("some error here")

	msg := BuildStackTrace(err)

	if err.Error() != msg {
		t.Errorf("FAILED: Wrong StackTrace message error, expected '%s' got '%s'", err.Error(), msg)
	}

	t.Log("SUCCESS!")
}
