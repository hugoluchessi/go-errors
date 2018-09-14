package errors

import "runtime"

type stackItem struct {
	pc   uintptr
	file string
	line int
}

func newStackItem(depth int) stackItem {
	pc, file, line, ok := runtime.Caller(depth)

	if !ok {
		return stackItem{}
	}

	return stackItem{pc, file, line}
}
