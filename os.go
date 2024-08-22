package must

import "os"

// OS provides functions for working with the [os] package.
var OS _os //nolint:gochecknoglobals

type _os struct {
	//
}

// Open panics if opening a file fails. It uses [os.Open] internally.
func (_os) Open(path string) *os.File {
	return Succeed(os.Open(path)) //nolint:gosec
}

// Create panics if creating a file fails. It uses [os.Create] internally.
func (_os) Create(path string) *os.File {
	return Succeed(os.Create(path)) //nolint:gosec
}

// Remove panics if removing a file fails. It uses [os.Remove] internally.
func (_os) Remove(path string) {
	NotError(os.Remove(path))
}
