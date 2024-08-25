package must

import "os"

// OS provides functions for working with the [os] package.
//
//nolint:gochecknoglobals
var OS OSPkg

// OSPkg provides functions for working with the [os] package.
type OSPkg struct {
	//
}

// Open panics if opening a file fails. It uses [os.Open] internally.
func (OSPkg) Open(path string) *os.File {
	return Succeed(os.Open(path)) //nolint:gosec
}

// Create panics if creating a file fails. It uses [os.Create] internally.
func (OSPkg) Create(path string) *os.File {
	return Succeed(os.Create(path)) //nolint:gosec
}

// Stat panics if getting file information fails. It uses [os.Stat] internally.
func (OSPkg) Stat(path string) os.FileInfo {
	return Succeed(os.Stat(path))
}

// Exists returns true if the file at the specified path exists, and false if
// not. It panics if getting file information fails.
func (OSPkg) Exists(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	NotError(err)

	return false
}

// Mkdir panics if creating a directory fails. It uses [os.Mkdir] internally.
func (OSPkg) Mkdir(path string, perm os.FileMode) {
	NotError(os.Mkdir(path, perm))
}

// MkdirAll panics if creating a directory fails. It uses [os.MkdirAll]
// internally.
func (OSPkg) MkdirAll(path string, perm os.FileMode) {
	NotError(os.MkdirAll(path, perm))
}

// Remove panics if removing a file fails. It uses [os.Remove] internally.
func (OSPkg) Remove(path string) {
	NotError(os.Remove(path))
}

// RemoveAll panics if removing a directory fails. It uses [os.RemoveAll]
// internally.
func (OSPkg) RemoveAll(path string) {
	NotError(os.RemoveAll(path))
}
