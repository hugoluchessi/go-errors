package errors

import (
	"errors"
	"testing"
)

func TestNewError(t *testing.T) {
	rootError := errors.New("some error here")

	state := map[string]interface{}{
		"some":      "value",
		"int_value": 1234,
	}

	ew := f3(rootError, state).(*errorWrapper)

	if ew.rootError != rootError {
		t.Errorf("FAILED: Wrong Root error, expected '%s' got '%s'", rootError.Error(), ew.Error())
	}

	if len(ew.errorStack) != 3 {
		t.Errorf("FAILED: stack size should be 3 got '%d'", len(ew.errorStack))
	}

	for key, value := range state {
		v, _ := ew.errorState[key]

		if v != value {
			t.Errorf("FAILED: Key '%s', hash incorrect value, expected '%s' got '%s'.", key, value, v)
		}
	}

	t.Log("SUCCESS!")
}
