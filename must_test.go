package must_test

import (
	"errors"
	"testing"

	"github.com/johnfrankmorgan/must"
	"github.com/stretchr/testify/assert"
)

func TestNotError_Error_Panics(t *testing.T) {
	t.Parallel()

	assert.PanicsWithError(t, "must: unexpected error: NotError", func() {
		must.NotError(errors.New("NotError"))
	})
}

func TestNotError_NoError_DoesNotPanic(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		must.NotError(nil)
	})
}

func TestSucceed_Error_ReturnsValue(t *testing.T) {
	t.Parallel()

	assert.PanicsWithError(t, "must: unexpected error: Succeed", func() {
		must.Succeed(0, errors.New("Succeed"))
	})
}

func TestSucceed_NoError_ReturnsValue(t *testing.T) {
	t.Parallel()

	got := must.Succeed(100, nil)

	assert.Equal(t, 100, got)
}
