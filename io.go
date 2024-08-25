package must

import "io"

// IO provides functions for working with the [io] package.
//
//nolint:gochecknoglobals
var IO IOPkg

// IOPkg provides functions for working with the [io] package.
type IOPkg struct {
	//
}

// Close can be used to panic if closing an [io.Closer] fails.
//
//	defer must.IO.Close(file)
func (IOPkg) Close(closer io.Closer) {
	NotError(closer.Close())
}

// Copy can be used to panic if a call to [io.Copy] fails.
func (IOPkg) Copy(dst io.Writer, src io.Reader) int64 {
	return Succeed(io.Copy(dst, src))
}

// ReadAll can be used to panic if a call to [io.ReadAll] fails.
func (IOPkg) ReadAll(r io.Reader) []byte {
	return Succeed(io.ReadAll(r))
}

// WriteString can be used to panic if a call to [io.WriteString] fails.
func (IOPkg) WriteString(w io.Writer, s string) {
	Succeed(io.WriteString(w, s))
}
