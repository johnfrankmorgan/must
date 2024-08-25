package must_test

import (
	"testing"

	"github.com/johnfrankmorgan/must"
	"github.com/stretchr/testify/assert"
)

func TestOpen_Error_Panics(t *testing.T) {
	t.Parallel()

	assert.PanicsWithError(t, "must: unexpected error: open does-not-exist: no such file or directory", func() {
		must.OS.Open("does-not-exist")
	})
}

func TestOpen_NoError_ReturnsFile(t *testing.T) {
	t.Parallel()

	f := must.OS.Open("os.go")
	defer must.IO.Close(f)

	assert.NotNil(t, f)
	assert.Equal(t, "os.go", f.Name())
}

func TestCreate_Error_Panics(t *testing.T) {
	t.Parallel()

	assert.PanicsWithError(t, "must: unexpected error: open /not-writable: permission denied", func() {
		must.OS.Create("/not-writable")
	})
}

func TestCreate_NoError_ReturnsFile(t *testing.T) {
	t.Parallel()

	f := must.OS.Create("tmp/Create.txt")
	defer must.IO.Close(f)
	defer must.OS.Remove(f.Name())

	assert.NotNil(t, f)
	assert.Equal(t, "tmp/Create.txt", f.Name())
}

func TestStat_Error_Panics(t *testing.T) {
	t.Parallel()

	assert.PanicsWithError(t, "must: unexpected error: stat does-not-exist: no such file or directory", func() {
		must.OS.Stat("does-not-exist")
	})
}

func TestStat_NoError_ReturnsFileInfo(t *testing.T) {
	t.Parallel()

	info := must.OS.Stat("os.go")

	assert.Positive(t, info.Size())
}

func TestMkdir_Error_Panics(t *testing.T) {
	t.Parallel()

	assert.PanicsWithError(t, "must: unexpected error: mkdir /not-writable: permission denied", func() {
		must.OS.Mkdir("/not-writable", 0o755)
	})
}

func TestMkdir_NoError_CreatesDirectory(t *testing.T) {
	t.Parallel()

	must.OS.Mkdir("tmp/Mkdir", 0o755)
	defer must.OS.RemoveAll("tmp/Mkdir")

	assert.True(t, must.OS.Exists("tmp/Mkdir"))
}

func TestMkdirAll_Error_Panics(t *testing.T) {
	t.Parallel()

	assert.PanicsWithError(t, "must: unexpected error: mkdir /not-writable: permission denied", func() {
		must.OS.MkdirAll("/not-writable", 0o755)
	})
}

func TestMkdirAll_NoError_CreatesDirectories(t *testing.T) {
	t.Parallel()

	must.OS.MkdirAll("tmp/MkdirAll/One/Two", 0o755)

	assert.True(t, must.OS.Exists("tmp/MkdirAll"))
	assert.True(t, must.OS.Exists("tmp/MkdirAll/One"))
	assert.True(t, must.OS.Exists("tmp/MkdirAll/One/Two"))
}

func TestRemove_Error_Panics(t *testing.T) {
	t.Parallel()

	assert.PanicsWithError(t, "must: unexpected error: remove does-not-exist: no such file or directory", func() {
		must.OS.Remove("does-not-exist")
	})
}

func TestRemove_NoError_RemovesFile(t *testing.T) {
	t.Parallel()

	f := must.OS.Create("tmp/Remove.txt")
	must.IO.Close(f)

	must.OS.Remove("tmp/Remove.txt")

	assert.False(t, must.OS.Exists("tmp/Remove.txt"))
}

func TestRemoveAll_NoError_RemovesDirectory(t *testing.T) {
	t.Parallel()

	must.OS.Mkdir("tmp/RemoveAll", 0o755)

	assert.NotPanics(t, func() {
		must.OS.RemoveAll("tmp/RemoveAll")
	})

	assert.False(t, must.OS.Exists("tmp/RemoveAll"))
}
