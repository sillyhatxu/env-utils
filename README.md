# env-utils

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![Go-Version](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/sillyhatxu/env-utils)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/sillyhatxu/env-utils)](https://pkg.go.dev/github.com/sillyhatxu/env-utils)
[![Build and Test](https://github.com/sillyhatxu/env-utils/workflows/Build%20and%20Test/badge.svg?branch=master&event=push)](https://github.com/sillyhatxu/env-utils/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/sillyhatxu/env-utils)](https://goreportcard.com/report/github.com/sillyhatxu/env-utils)
[![codecov](https://codecov.io/gh/sillyhatxu/env-utils/branch/master/graph/badge.svg)](https://codecov.io/gh/sillyhatxu/env-utils)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://choosealicense.com/licenses/mit/)
[![Release](https://img.shields.io/github/release/sillyhatxu/env-utils.svg?style=flat-square)](https://github.com/sillyhatxu/env-utils/releases)

## example

[config_test.go](https://github.com/sillyhatxu/env-utils/blob/master/config_test.go)

load evn config to struct

```go
package main

import (
    "fmt"
    "github.com/sillyhatxu/env-utils"
	"os"
	"strings"
	"testing"
)

type EnvConfig struct {
	TestString    string  `env:"TEST.STRING"`
	TestInt       int     `env:"TEST.INT"`
	TestInt8      int8    `env:"TEST.INT8"`
	TestInt32     int32   `env:"TEST.INT32"`
	TestInt64     int64   `env:"TEST.INT64"`
	TestBoolean   bool    `env:"TEST.BOOLEAN"`
	TestFloat32   float32 `env:"TEST.FLOAT32"`
	TestFloat64   float64 `env:"TEST.FLOAT64"`
	TestURL       string  `env:"TEST.URL"`
	TestNilString string  `env:"TEST.NIL_STRING"`
}

func main() {
	var config EnvConfig
	fileName := ""
	if os.Getenv("APPLICATION_ENV") == "" {
		fileName = "local.env"
	} else {
		fileName = ".env"
	}
	err := envutils.ParseConfig(&config, envutils.Filenames([]string{fileName}))
    if err != nil{
        panic(err)
    }
    fmt.Println(config.TestString)
    fmt.Println(config.TestURL)
    fmt.Println(config.TestNilString)
    fmt.Println(config.TestInt)
    fmt.Println(config.TestInt8)
    fmt.Println(config.TestInt32)
    fmt.Println(config.TestInt64)
    fmt.Println(config.TestBoolean)
    fmt.Println(config.TestFloat64)
    fmt.Println(config.TestFloat64)
}
```