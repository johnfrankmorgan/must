package must

import "io"

// Close can be used to panic if closing an [io.Closer] fails.
//
//	defer must.Close(file)
func Close(closer io.Closer) {
	NotError(closer.Close())
}
