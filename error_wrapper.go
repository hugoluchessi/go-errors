package errors

import (
	"fmt"
	"strings"
)

const (
	ErrorHeaderFormat     = "Error: '%s'\n"
	ErrorStateHeader      = "State:\n"
	ErrorStackTraceHeader = "Stack Trace:\n"

	ErrorStateFormat      = "\t%s: %+v\n"
	ErrorStackTraceFormat = "\t%d - %s:%d\n"
)

type errorWrapper struct {
	rootError  error
	errorStack []stackItem
	errorState map[string]interface{}
}

func WrapError(err error) error {
	ew, ok := err.(*errorWrapper)

	if !ok {
		return NewError(err, make(map[string]interface{}))
	}

	ew.errorStack = append(ew.errorStack, newStackItem(2))
	return ew
}

func (ew errorWrapper) Error() string {
	return ew.rootError.Error()
}

func (ew errorWrapper) StackTrace() string {
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
	for i, stackItem := range ew.errorStack {
		sb.WriteString(
			fmt.Sprintf(
				ErrorStackTraceFormat,
				i,
				stackItem.file,
				stackItem.line,
			),
		)
	}

	return sb.String()
}

func (ew errorWrapper) RootError() error {
	return ew.rootError
}

func NewError(rootError error, state map[string]interface{}) error {
	return &errorWrapper{
		rootError:  rootError,
		errorStack: []stackItem{newStackItem(2)},
		errorState: state,
	}
}
