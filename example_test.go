package must_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/johnfrankmorgan/must"
)

func Example() {
	base := filepath.Join(os.TempDir(), "must")
	must.OS.RemoveAll(base)

	must.OS.MkdirAll(base, 0o755)
	defer must.OS.RemoveAll(base)

	input := must.OS.Create(filepath.Join(base, "input.txt"))
	must.IO.WriteString(input, "Hello, world!")
	must.IO.Close(input)

	input = must.OS.Open(filepath.Join(base, "input.txt"))
	defer must.IO.Close(input)

	contents := must.IO.ReadAll(input)

	fmt.Println(string(contents)) // Output: Hello, world!
}
