package must

import "io"

// IO provides functions for working with the [io] package.
var IO _io //nolint:gochecknoglobals

type _io struct {
	//
}

// Close can be used to panic if closing an [io.Closer] fails.
//
//	defer must.IO.Close(file)
func (_io) Close(closer io.Closer) {
	NotError(closer.Close())
}

// Copy can be used to panic if a call to [io.Copy] fails.
func (_io) Copy(dst io.Writer, src io.Reader) int64 {
	return Succeed(io.Copy(dst, src))
}

// ReadAll can be used to panic if a call to [io.ReadAll] fails.
func (_io) ReadAll(r io.Reader) []byte {
	return Succeed(io.ReadAll(r))
}
