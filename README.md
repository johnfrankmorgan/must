# must

Package `must` provides error handling utilities for Go.

[![CI](https://github.com/johnfrankmorgan/must/actions/workflows/ci.yaml/badge.svg)](https://github.com/johnfrankmorgan/must/actions/workflows/ci.yaml)
[![Coverage](https://codecov.io/gh/johnfrankmorgan/must/graph/badge.svg?token=FA6ZZd1UVM)](https://codecov.io/gh/johnfrankmorgan/must)
[![Go Report Card](https://goreportcard.com/badge/github.com/johnfrankmorgan/must)](https://goreportcard.com/report/github.com/johnfrankmorgan/must)
[![Go Reference](https://pkg.go.dev/badge/github.com/johnfrankmorgan/must.svg)](https://pkg.go.dev/github.com/johnfrankmorgan/must)

```go
file := must.Create("file.json")
defer must.Close(file)
must.NotError(json.NewEncoder(file).Encode(data))
```

See the [documentation](https://pkg.go.dev/github.com/johnfrankmorgan/must)
for more information.
