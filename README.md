# Goterator

[![Test and Build](https://github.com/yaa110/goterator/workflows/Test%20and%20Build/badge.svg)](https://github.com/yaa110/goterator/actions?query=workflow%3A"Test+and+Build") [![Documentation](https://img.shields.io/badge/Documentation-goterator-blue)](https://pkg.go.dev/github.com/yaa110/goterator?tab=doc) [![Go Report](https://goreportcard.com/badge/github.com/yaa110/goterator)](https://goreportcard.com/report/github.com/yaa110/goterator)

Iterator implementation for Golang to provide map and reduce functionalities.

## Package

```go
import "github.com/yaa110/goterator"
```

## Getting Started

- Create a generator from slices

```go
words := []interface{}{"an", "example", "of", "goterator"}

generator := goterator.NewSliceGenerator(words)
```

- Create a generator from channels

```go
chn := make(chan interface{}, 4)
chn <- "an"
chn <- "example"
chn <- "of"
chn <- "goterator"
close(chn)

generator := goterator.NewChannelGenerator(chn)
```

- Create custom generators

```go
type TestGenerator struct {
    words []string
    i     int
}

func (g *TestGenerator) Next() (interface{}, error) {
    if g.i == len(g.words) {
        return nil, goterator.End()
    }
    word := g.words[g.i]
    g.i++
    return word, nil
}

generator := &TestGenerator{
    words: []string{"an", "example", "of", "goterator"},
    i: 0,
}
```

- Iterate over generators

```go
lengths := goterator.New(generator).Map(func(word interface{}) interface{} {
    return len(word.(string))
}).Collect()

assert.Equal([]interface{}{2, 7, 2, 9}, lengths)
```

Please for more information about mappers and reducers (consumers) check the [documentation](https://pkg.go.dev/github.com/yaa110/goterator?tab=doc).
