package goerrors

import "runtime"

type StackTraceItem struct {
	PC   uintptr
	File string
	Line int
}

func NewStackTraceItem(depth int) StackTraceItem {
	pc, file, line, ok := runtime.Caller(depth)

	if !ok {
		return StackTraceItem{}
	}

	return StackTraceItem{
		PC:   pc,
		File: file,
		Line: line,
	}
}
