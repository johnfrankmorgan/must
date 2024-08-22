package must_test

import (
	"errors"
	"testing"

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
