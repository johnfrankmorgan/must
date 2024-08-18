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

func TestClose_Error_Panics(t *testing.T) {
	t.Parallel()

	c := &closer{err: errors.New("Close")}

	assert.PanicsWithError(t, "must: unexpected error: Close", func() {
		must.Close(c)
	})

	assert.True(t, c.called)
}

func TestClose_NoError_DoesNotPanic(t *testing.T) {
	t.Parallel()

	c := &closer{}

	assert.NotPanics(t, func() {
		must.Close(c)
	})

	assert.True(t, c.called)
}
