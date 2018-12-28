package goerrors

import (
	"fmt"
	"strings"
)

const (
	ErrorHeaderFormat     = "Error: '%s'\n"
	ErrorStateHeader      = "State:\n"
	ErrorStackTraceHeader = "Stack Trace:\n"
	ErrorStateFormat      = "\t%s: %+v\n"
)

type errorWrapper struct {
	rootError       error
	errorStackTrace *StackTrace
	errorState      map[string]interface{}
}

func (ew errorWrapper) Error() string {
	return ew.rootError.Error()
}

func (ew errorWrapper) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(ErrorHeaderFormat, ew.Error()))

	sb.WriteString(ErrorStateHeader)
	for stateKey, stateValue := range ew.errorState {
		sb.WriteString(
			fmt.Sprintf(
				ErrorStateFormat,
				stateKey,
				stateValue,
			),
		)
	}

	sb.WriteString(ErrorStackTraceHeader)
	sb.WriteString(ew.errorStackTrace.String())
	return sb.String()
}

func NewError(rootError error) error {
	st := NewStackTrace()
	st.AddStackItem()

	return &errorWrapper{
		rootError:       rootError,
		errorStackTrace: st,
		errorState:      make(map[string]interface{}),
	}
}

func NewErrorWithState(rootError error, state map[string]interface{}) error {
	st := NewStackTrace()
	st.AddStackItem()

	return &errorWrapper{
		rootError:       rootError,
		errorStackTrace: st,
		errorState:      state,
	}
}

func WrapError(err error) error {
	ew, ok := err.(*errorWrapper)

	if !ok {
		return NewError(err)
	}

	ew.errorStackTrace.AddStackItem()
	return ew
}

func BuildStackTrace(err error) string {
	ew, ok := err.(*errorWrapper)

	if !ok {
		return err.Error()
	}

	return ew.String()
}

func RootError(err error) error {
	ew, ok := err.(*errorWrapper)

	if !ok {
		return err
	}

	return ew.rootError
}
