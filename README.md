# jsonarray

[![Go Reference](https://pkg.go.dev/badge/github.com/brokeyourbike/jsonarray.svg)](https://pkg.go.dev/github.com/brokeyourbike/jsonarray)
[![Go Report Card](https://goreportcard.com/badge/github.com/brokeyourbike/jsonarray)](https://goreportcard.com/report/github.com/brokeyourbike/jsonarray)
[![Maintainability](https://api.codeclimate.com/v1/badges/2e5b535a4edce1a5f803/maintainability)](https://codeclimate.com/github/brokeyourbike/jsonarray/maintainability)
[![codecov](https://codecov.io/gh/brokeyourbike/jsonarray/branch/main/graph/badge.svg?token=20MwbA3eAn)](https://codecov.io/gh/brokeyourbike/jsonarray)

GORM JSON Array Types

## Installation

```bash
go get github.com/brokeyourbike/jsonarray
```

## Usage

```go
package main

import (
    "github.com/brokeyourbike/jsonarray"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Tags jsonarray.Slice[string]
}
```

## Authors
- [Ivan Stasiuk](https://github.com/brokeyourbike) | [Twitter](https://twitter.com/brokeyourbike) | [LinkedIn](https://www.linkedin.com/in/brokeyourbike) | [stasi.uk](https://stasi.uk)

## Thanks

- https://github.com/go-gorm/datatypes
- https://go.dev/blog/go1.18

## License
[MIT License](https://github.com/glocurrency/jsonarray/blob/main/LICENSE)