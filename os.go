package must

import "os"

// Open panics if opening a file fails. It uses [os.Open] internally.
func Open(path string) *os.File {
	return Succeed(os.Open(path)) //nolint:gosec
}

// Create panics if creating a file fails. It uses [os.Create] internally.
func Create(path string) *os.File {
	return Succeed(os.Create(path)) //nolint:gosec
}

// Remove panics if removing a file fails. It uses [os.Remove] internally.
func Remove(path string) {
	NotError(os.Remove(path))
}
