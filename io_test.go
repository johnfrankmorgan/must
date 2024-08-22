package must_test

import (
	"bytes"
	"errors"
	"testing"
	"testing/iotest"

	"github.com/johnfrankmorgan/must"
	"github.com/stretchr/testify/assert"
)

type closer struct {
	err    error
	called bool
}

func (c *closer) Close() error {
	c.called = true

	return c.err
}

func TestIOClose_Error_Panics(t *testing.T) {
	t.Parallel()

	c := &closer{err: errors.New("Close")}

	assert.PanicsWithError(t, "must: unexpected error: Close", func() {
		must.IO.Close(c)
	})

	assert.True(t, c.called)
}

func TestIOClose_NoError_DoesNotPanic(t *testing.T) {
	t.Parallel()

	c := &closer{}

	assert.NotPanics(t, func() {
		must.IO.Close(c)
	})

	assert.True(t, c.called)
}

func TestIOCopy_Error_Panics(t *testing.T) {
	t.Parallel()

	src := iotest.ErrReader(errors.New("Copy"))

	assert.PanicsWithError(t, "must: unexpected error: Copy", func() {
		must.IO.Copy(&bytes.Buffer{}, src)
	})
}

func TestIOCopy_NoError_CopiesBytes(t *testing.T) {
	t.Parallel()

	src := bytes.NewBufferString("Hello, World!")
	dst := bytes.NewBuffer(nil)

	n := must.IO.Copy(dst, src)

	assert.Equal(t, "Hello, World!", dst.String())
	assert.EqualValues(t, 13, n)
}

func TestIOReadAll_Error_Panics(t *testing.T) {
	t.Parallel()

	src := iotest.ErrReader(errors.New("ReadAll"))

	assert.PanicsWithError(t, "must: unexpected error: ReadAll", func() {
		must.IO.ReadAll(src)
	})
}

func TestIOReadAll_NoError_ReadsData(t *testing.T) {
	t.Parallel()

	src := bytes.NewBufferString("Hello, World!")

	res := must.IO.ReadAll(src)

	assert.Equal(t, "Hello, World!", string(res))
}
