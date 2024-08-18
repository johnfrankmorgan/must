// Package must provides a set of helper functions to make error handling more
// concise. The functions in this package are intended to be used in situations
// where an error should never occur.
//
//	file := must.Create("file.txt")
//	defer must.Close(file)
package must

import (
	"errors"
	"fmt"
)

// ErrUnexpected is used when any panics occur.
var ErrUnexpected = errors.New("must: unexpected error")

// NotError panics if the provided error is not nil.
//
//	must.NotError(json.NewDecoder(r).Decode(&value))
func NotError(err error) {
	if err != nil {
		panic(fmt.Errorf("%w: %w", ErrUnexpected, err))
	}
}

// Succeed can be used to convert a (T, error) into a T. It panics if the
// provided error is not nil.
//
//	file := must.Succeed(os.Create("file.txt"))
func Succeed[T any](value T, err error) T { //nolint:ireturn
	NotError(err)

	return value
}
