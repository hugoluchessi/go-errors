package errors

func f2(err error, state map[string]interface{}) error {
	e := f1(err, state)
	return WrapError(e)
}
