package must_test

import (
	"testing"

	"github.com/johnfrankmorgan/must"
	"github.com/stretchr/testify/assert"
)

func TestOpen_Error_Panics(t *testing.T) {
	t.Parallel()

	assert.PanicsWithError(t, "must: unexpected error: open does-not-exist: no such file or directory", func() {
		must.Open("does-not-exist")
	})
}

func TestOpen_NoError_ReturnsFile(t *testing.T) {
	t.Parallel()

	f := must.Open("os.go")
	defer must.Close(f)

	assert.NotNil(t, f)
	assert.Equal(t, "os.go", f.Name())
}

func TestCreate_Error_Panics(t *testing.T) {
	t.Parallel()

	assert.PanicsWithError(t, "must: unexpected error: open /not-writable: permission denied", func() {
		must.Create("/not-writable")
	})
}

func TestCreate_NoError_ReturnsFile(t *testing.T) {
	t.Parallel()

	f := must.Create("tmp/Create.txt")
	defer must.Close(f)
	defer must.Remove(f.Name())

	assert.NotNil(t, f)
	assert.Equal(t, "tmp/Create.txt", f.Name())
}

func TestRemove_Error_Panics(t *testing.T) {
	t.Parallel()

	assert.PanicsWithError(t, "must: unexpected error: remove does-not-exist: no such file or directory", func() {
		must.Remove("does-not-exist")
	})
}

func TestRemove_NoError_DoesNotPanic(t *testing.T) {
	t.Parallel()

	f := must.Create("tmp/Remove.txt")
	must.Close(f)

	assert.NotPanics(t, func() {
		must.Remove("tmp/Remove.txt")
	})
}
