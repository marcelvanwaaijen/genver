# Description

This utility can be used to automatically generate a `version.go` file with the current git tag version. To get the current version, it executes the following command:
`git describe --tags`. This will then use the full string for `gitcommit` and the string up to the first "-" as the `version`, which will then be merged in `version.go`.

## Install
`go install github.com/marcelvanwaaijen/genver@latest`

## Usage
The Makefile in this repo shows an example of its usage. Just run `genver` before you build your project.

or you can use `go:generate`:
```go
package main

import "fmt"

//go:generate genver.exe

func main() {
	fmt.Println(version)
      fmt.Println(gitcommit)
}

```

### Extra options:
```sh
$ genver -h

Usage of genver:
  -o string
        output file name (default "version.go")
  -p string
        package name (default "main")
  -version
        show version
```
