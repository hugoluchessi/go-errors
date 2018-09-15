package errors

import (
	"errors"
	"strings"
	"testing"
)

func TestNewError(t *testing.T) {
	rootError := errors.New("some error here")

	state := map[string]interface{}{
		"some":      "value",
		"int_value": 1234,
	}

	ew := NewError(rootError, state).(*errorWrapper)

	if ew.rootError != rootError {
		t.Errorf("FAILED: Wrong Root error, expected '%s' got '%s'", rootError.Error(), ew.Error())
	}

	if len(ew.errorStack) != 1 {
		t.Errorf("FAILED: stack size should be 1 got '%d'", len(ew.errorStack))
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

	if len(castEw.errorStack) != 1 {
		t.Errorf("FAILED: stack size should be 1 got '%d'", len(castEw.errorStack))
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

	if len(ew.errorStack) != 3 {
		t.Errorf("FAILED: stack size should be 3 got '%d'", len(ew.errorStack))
	}

	if !strings.HasSuffix(ew.errorStack[0].file, fileNameFn1) {
		t.Errorf("FAILED: first stack item file name should be '%s' got '%s'", fileNameFn1, ew.errorStack[0].file)
	}

	if ew.errorStack[0].line != lineFn1 {
		t.Errorf("FAILED: first stack item line should be '%d' got '%d'", lineFn1, ew.errorStack[0].line)
	}

	if !strings.HasSuffix(ew.errorStack[1].file, fileNameFn2) {
		t.Errorf("FAILED: second stack item file name should be '%s' got '%s'", fileNameFn2, ew.errorStack[1].file)
	}

	if ew.errorStack[1].line != lineFn2 {
		t.Errorf("FAILED: second stack item line should be '%d' got '%d'", lineFn2, ew.errorStack[1].line)
	}

	if !strings.HasSuffix(ew.errorStack[2].file, fileNameFn3) {
		t.Errorf("FAILED: third stack item file name should be '%s' got '%s'", fileNameFn3, ew.errorStack[2].file)
	}

	if ew.errorStack[2].line != lineFn3 {
		t.Errorf("FAILED: third stack item line should be '%d' got '%d'", lineFn3, ew.errorStack[2].line)
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

	state := map[string]interface{}{}

	ew := NewError(rootError, state)

	if ew.Error() != rootError.Error() {
		t.Errorf("FAILED: Wrong Root error, expected '%s' got '%s'", rootError.Error(), ew.Error())
	}

	t.Log("SUCCESS!")
}

func TestRootError(t *testing.T) {
	rootError := errors.New("some error here")

	state := map[string]interface{}{}

	ew := NewError(rootError, state).(*errorWrapper)

	if ew.RootError() != rootError {
		t.Errorf("FAILED: Wrong Root error, expected '%s' got '%s'", rootError.Error(), ew.Error())
	}

	t.Log("SUCCESS!")
}

func TestStackTrace(t *testing.T) {
	rootError := errors.New("some error here")
	state := map[string]interface{}{
		"some":      "value",
		"int_value": 1234,
	}

	ew := f3(rootError, state).(*errorWrapper)

	traceMessage := ew.StackTrace()

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
