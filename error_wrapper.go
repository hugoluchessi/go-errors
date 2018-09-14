package errors

type errorWrapper struct {
	rootError  error
	errorStack []stackItem
	errorState map[string]interface{}
}

func WrapError(err error) *errorWrapper {
	ew, ok := err.(*errorWrapper)

	if !ok {
		ew = NewError(err, make(map[string]interface{})).(*errorWrapper)
	}

	ew.errorStack = append(ew.errorStack, newStackItem(2))
	return ew
}

func (ew errorWrapper) Error() string {
	return ew.rootError.Error()
}

func (ew errorWrapper) StackTrace() []string {
	// TODO: DO!
	return []string{}
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
