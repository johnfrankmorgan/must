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
