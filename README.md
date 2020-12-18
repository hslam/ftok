# ftok
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hslam/ftok)](https://pkg.go.dev/github.com/hslam/ftok)
[![Build Status](https://github.com/hslam/ftok/workflows/build/badge.svg)](https://github.com/hslam/ftok/actions)
[![codecov](https://codecov.io/gh/hslam/ftok/branch/master/graph/badge.svg)](https://codecov.io/gh/hslam/ftok)
[![Go Report Card](https://goreportcard.com/badge/github.com/hslam/ftok)](https://goreportcard.com/report/github.com/hslam/ftok)
[![LICENSE](https://img.shields.io/github/license/hslam/ftok.svg?style=flat-square)](https://github.com/hslam/ftok/blob/master/LICENSE)

Package ftok provides a way to generate a System V IPC key, suitable for using with msgget, semget, or shmget.

## Get started

### Install
```
go get github.com/hslam/ftok
```
### Import
```
import "github.com/hslam/ftok"
```
### Usage
#### Example
```go
package main

import (
	"fmt"
	"github.com/hslam/ftok"
)

func main() {
	key, err := ftok.Ftok("/tmp", 0x22)
	if err != nil {
		return
	}
	fmt.Printf("%x", key)
}
```

### License
This package is licensed under a MIT license (Copyright (c) 2020 Meng Huang)


### Author
ftok was written by Meng Huang.


