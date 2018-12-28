package goerrors

func f1(err error, state map[string]interface{}) error {
	return NewErrorWithState(err, state)
}
