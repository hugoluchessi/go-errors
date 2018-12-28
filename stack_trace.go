package goerrors

import (
	"fmt"
	"strings"
)

type StackTrace struct {
	Items []StackTraceItem
}

const (
	stackDepth       = 3
	stackTraceFormat = "\t%d - %s:%d\n"
)

func NewStackTrace() *StackTrace {
	return &StackTrace{
		Items: []StackTraceItem{},
	}
}

func (st *StackTrace) AddStackItem() {
	st.Items = append(st.Items, NewStackTraceItem(stackDepth))
}

func (st *StackTrace) String() string {
	var sb strings.Builder

	sb.WriteString(ErrorStackTraceHeader)
	for i, stackItem := range st.Items {
		sb.WriteString(
			fmt.Sprintf(
				stackTraceFormat,
				i,
				stackItem.File,
				stackItem.Line,
			),
		)
	}

	return sb.String()
}
