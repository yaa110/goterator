# Goterator

[![Test and Build](https://github.com/yaa110/goterator/workflows/Test%20and%20Build/badge.svg)](https://github.com/yaa110/goterator/actions?query=workflow%3A"Test+and+Build") [![PkgGoDev](https://pkg.go.dev/badge/github.com/yaa110/goterator)](https://pkg.go.dev/github.com/yaa110/goterator) [![GoDoc](https://img.shields.io/badge/godoc-goterator-blue)](https://godoc.org/github.com/yaa110/goterator) [![Go Report](https://goreportcard.com/badge/github.com/yaa110/goterator)](https://goreportcard.com/report/github.com/yaa110/goterator) [![Coverage](https://gocover.io/_badge/github.com/yaa110/goterator)](https://gocover.io/github.com/yaa110/goterator)

Iterator implementation for Golang to provide map and reduce functionalities.

## Package

```go
import (
    "github.com/yaa110/goterator"
    "github.com/yaa110/goterator/generator"
)
```

## Getting Started

- Create a generator from slices

```go
words := []interface{}{"an", "example", "of", "goterator"}

gen := generator.NewSlice(words)
```

- Create a generator from channels

```go
chn := make(chan interface{}, 4)
chn <- "an"
chn <- "example"
chn <- "of"
chn <- "goterator"
close(chn)

gen := generator.NewChannel(chn)
```

- Create custom generators

```go
type TestGenerator struct {
    words []string
    i     int
    value string
}

func (g *TestGenerator) Next() bool {
    if g.i == len(g.words) {
        return false
    }
    g.value = g.words[g.i]
    g.i++
    return true
}

func (g *TestGenerator) Value() interface{} {
    return g.value
}

gen := &TestGenerator{
    words: []string{"an", "example", "of", "goterator"},
    i: 0,
    value: "",
}
```

- Iterate over generators

```go
lengths := goterator.New(gen).Map(func(word interface{}) interface{} {
    return len(word.(string))
}).Collect()

assert.Equal([]interface{}{2, 7, 2, 9}, lengths)
```

Please for more information about mappers and reducers (consumers) check the [documentation](https://godoc.org/github.com/yaa110/goterator).
