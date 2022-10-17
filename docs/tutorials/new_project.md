# Create a Go project

This tutorial is more than a quickstart; it aims to help you get up and running with a new Go project while
also providing a deeper look into Go modules and the `go` commands you'll need
to work with a Go project effectively.

## Prerequisites

* Go installed

## Create a Go module

1. Make a directory named **tutorial** and change into it:
   `mkdir tutorial && cd tutorial`
2. Initialize a Go module: `go mod init tutorial/myproj`

If you're planning to host a Go project on GitHub, you can instead
initialize the module with the name of the GitHub repo. For example, if you have
(or will have) a GitHub username of `myuser` and a repo of `myproj`, you'd run
`go mod init github.com/myuser/myproj`.

* A module is a collection of packages that can be released and distributed as
  a unit.
* Every module has a module path that identifies it. The module path is declared
  in the **go.mod** file and prefixes package import paths from the module.
  `github.com/pcoet/golang-patterns` is a module path.
* When you run `go mod init`, you create a **go.mod** file that defines the
  module path and Go version for the module. If you add dependencies, these will
  be defined in the **go.mod** file too. The root directory of a module is, by
  definition, the directory containing the **go.mod** file.
* The main module is the module where you run the `go` command. The main module
  can be in a directory with a **go.mod** file, or it can be in a child of the
  directory with the **go.mod** file.
* A package is a set of source files compiled together from the same directory.
  Every package is uniquely identified by an import path - a string formed by
  prepending the module path to the directory of the package.
* You can find the official definitions of module entities in the
  [glossary of the Go module reference](https://go.dev/ref/mod#glossary).

Notes about the `go mod` command:

* Supports operations on modules.
* `go mod init` creates a **go.mod** file in the current directory.
* The presence of a **go.mod** file in a directory tells Go that the directory
  is the root of a Go module.
* The only argument to `go mod init` is the path of the module to be created, e.g.
  `tutorial/myproj`. The path argument is optional, and `go mod` tries to
  infer the path from existing resources, if you omit it.

## Create a package

Create an example package:

1. In the **tutorial** directory, create a **pkg/multiplier.go** file:
   `mkdir pkg && touch pkg/multiplier.go`
2. Copy the following function into **multiplier.go** and save the file.
   ```go
   func Multiplier(m float64) func(float64) float64 {
     return func(n float64) float64 {
       return m * n
     }
   }
   ```

`Multiplier` takes a float, `m`, and returns a function that takes float `n` and
returns the result of `m * n`. `Multiplier` implements a
[function closure](https://go.dev/tour/moretypes/25) and demonstrates Go support
for
[higher-order functions](https://en.wikipedia.org/wiki/Higher-order_function).
Although function closures and higher-order functions are not directly related
to this tutorial, they're neat features and worth knowing about.

## Test a package

Create and run test:

1. In the **tutorial** directory, create a **pkg/multiplier_test.go** file:
   `touch pkg/multiplier_test.go`
2. Copy the following function into **multiplier_test.go** and save the file.
   ```go
   func TestMultiplier(t *testing.T) {
	   double := Multiplier(2)
	   want := 20.0
	   got := double(10)

	   if got != want {
		   t.Errorf("got %v; want %v", got, want)
	   }
   }
   ```
3. Run `TestMultiplier`:
   ```
   go test ./pkg/...
   ```
   If there were other tests in the **pkg** directory, this command would run them too.

## Create a main file

<!-- TODO: start here; use as a reference: https://github.com/pcoet/golang-patterns/tree/main/pkg/examples -->

## Learn more

* [Go docs: Go Modules Reference](https://go.dev/ref/mod)
* [Go docs: How to Write Go Code](https://go.dev/doc/code)
* [Go docs: How to Write Go Code (with GOPATH)](https://go.dev/doc/gopath_code)
* [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

Review the following:

* https://golang.org/ref/mod#go-mod-init
* https://golang.org/doc/tutorial/create-module
