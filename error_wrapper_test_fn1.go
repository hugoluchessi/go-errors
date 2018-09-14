package errors

func f1(err error, state map[string]interface{}) error {
	return NewError(err, state)
}
