package errors

func f3(err error, state map[string]interface{}) error {
	e := f2(err, state)

	return WrapError(e)
}
