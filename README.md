# goerrors
Simple way to wrap errors and have a backtrace with state.

## Why use goerrors?
Handling errors in go can prove to be a challenge in more complex situations, and not having a callstack to trace what went wrong can be even more challenging.

This package makes easier to wrap errors and trace your callstack and add some state on where the error was generated.

## How to use?
When you make a external (or internal) call, and an error is returned, simply wrap the error and return it.

``` go
// Main code which will start the call stack
func MainFoo(strFoo string, intFoo int64) {
    // When this function is called, an wrapped error will return
    value, err := Foo1(strFoo, intFoo)

    // To find out what was the original error
    rootError := goerrors.RootError(err)

    // Now you can type assert, or use the error to make the correct handling
    switch err.(type) {
        case baz:
            // Do something cool here
        default:
            // For unhandled errors you can log the stack trace with:
            msg := goerrors.BuildStackTrace(err)
            log.Error(msg)

            // This will try to build a string format of a stack trace, or return the 
            // .Error() of the given error            
    }
}

// Intermediate code will only wrap the error
func Foo1(strFoo string, intFoo int64) (string, error) {
    value, err := Foo2(strFoo, intFoo)

    if err != nil {
        return "", goerrors.WrapError(err)
    }

    return value, nil
}

// Code in which the error is generated (by own code or third party packages)
func Foo2(strFoo string, intFoo int64) (string, error) {
    // Creating the error with state
    err := bar.ErrorReturnningCall()

    if err != nil {
        return "", goerrors.NewErrorWithState(
            err, 
            map[string]interface{} {
                "strFoo": strFoo,
                "intFoo": intFoo,
            }
        )
    }

    return "finish!", nil
}
```

As the wrap error is an error migrate your code to use `goerrors` will be easy!


