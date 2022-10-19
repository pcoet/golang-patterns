# How to create a Go project

**Table of contents**

* [Before you begin](#before-you-begin)
  * [Assumptions](#assumptions)
  * [Prerequisites](#prerequisites)
* [Create a module](#create-a-module)

This tutorial walks you through the process of creating a Go project. By
working through the tutorial, you'll accomplish the following objectives:

1. Set up a project structure that you can use for Go development
2. Understand the layout and packaging of a typical Go project
3. Learn about the most important tools for managing Go source code

<a id="before-you-begin"></a>
## Before you begin

Before starting the tutorial, make sure you're familiar with the assumptions and
prerequisites.

<a id="assumptions"></a>
### Assumptions

This tutorial makes the following assumptions:

* **You're new to Go.** The tutorial assumes that you're still learning Go, but
  it doesn't discuss how to write Go code. If you're looking for an overview of
  the Go language, [A Tour of Go](https://go.dev/tour/welcome/1) is a great
  place to start.
* **You have a Unix-like OS.** If you're not working on macOS or Linux,
you might need to adapt some of the commands for your environment.

<a id="prerequisites"></a>
### Prerequisites

To complete this tutorial, you need to have Go installed. For help installing
Go, see [Download and install](https://go.dev/doc/install).

<a id="create-a-module"></a>
## Create a module

A Go [module](https://go.dev/ref/mod#glos-module) is a collection of packages
that can be released, distributed, and versioned together.

To get started with a new project, create a Go module:

1. Make a directory named **myproj** and change into it:
   `mkdir myproj && cd myproj`
2. Initialize a Go module: `go mod init myproj`

When you run `go mod init`, you create a **go.mod** file in the current
directory. The directory containing the
[go.mod file](https://go.dev/ref/mod#glos-go-mod-file) is, by definition, the
root directory of the module.

The contents of your new **go.mod** file should look something like this:

```
module myproj

go 1.19
```

The first line defines the
[module path](https://go.dev/ref/mod#glos-module-path), which uniquely
identifies the module and prefixes import paths for packages contained in the
module.

The module path can also include a
[repository root path](https://go.dev/ref/mod#glos-repository-root-path),
which is a path segment that specifies the root directory of a version control
repository. For example, if you're planning to distribute a Go module from
GitHub under the user `myuser`, and you have a repository named `myproj`, you'd
initialize your module with `go mod init github.com/myuser/myproj`, where
`github.com/myuser/myproj` is the repository root path.

## Create a package

<!-- TODO: start here -->

There's more than one way to structure a Go project. This tutorial relies on the
[Standard Go Project Layout](https://github.com/golang-standards/project-layout).

<!-- TODO: paraphrase glossary and cite -->
* A package is a set of source files compiled together from the same directory.
  Every package is uniquely identified by an import path - a string formed by
  prepending the module path to the directory of the package.

Create an example package:

1. In the **tutorial** directory, create a **pkg/mypack/multiplier.go** file:
   `mkdir -p pkg/mypack && touch pkg/mypack/multiplier.go`
2. Copy the following code into **multiplier.go** and save the file.
   ```go
   package mypack

   func Multiplier(m float64) func(float64) float64 {
     return func(n float64) float64 {
       return m * n
     }
   }
   ```
3. (Optional) Format the file: `gofmt pkg/mypack/multiplier.go`. If you're working in
   and IDE with Go tools installed, the file might be automatically formatted on
   save.

`Multiplier` takes a float, `m`, and returns a function that takes float `n` and
returns the result of `m * n`. `Multiplier` implements a
[function closure](https://go.dev/tour/moretypes/25) and demonstrates Go support
for
[higher-order functions](https://en.wikipedia.org/wiki/Higher-order_function).
Although function closures and higher-order functions are not directly related
to this tutorial, they're neat features and worth knowing about.

<!-- TODO: explain the pkg directory (see SGPL) -->

## Test a package

Create and run a test:

1. In the **tutorial** directory, create a **pkg/multiplier_test.go** file:
   `touch pkg/mypack/multiplier_test.go`
2. Copy the following code into **multiplier_test.go** and save the file:

   ```go
   package mypack

   import (
     "testing"
   )

   func TestMultiplier(t *testing.T) {
     double := Multiplier(2)
     want := 20.0
     got := double(10)

     if got != want {
       t.Errorf("got %v; want %v", got, want)
     }
   }
   ```
3. (Optional) Format the file: `gofmt pkg/mypack/multiplier_test.go`
4. Run `TestMultiplier`:
   ```
   go test ./pkg/...
   ```
   If there were other tests in the **pkg** directory, this command would run them too.

   <!-- TODO: explain go test -->

## Create a main file

If you're just creating a library for other applications to consume, you already
have the basic structure of your package. But if want to run your code, you
need a main file. Your project could even have multiple main files, each for
a different client.

Create and run a main file:

1. In the **tutorial** directory, create a **cmd/app/main.go** file:
   `mkdir -p cmd/myapp && touch cmd/myapp/main.go`
2. Copy the following code into **main.go** and save the file:

   ```go
   package main

   import (
     "fmt"

     "myproj/pkg/mypack"
   )

   func main() {
     double := mypack.Multiplier(2)
     fmt.Println(double(2))
   }
   ```
3. (Optional) Format the file: `gofmt cmd/myapp/main.go`
4. Run the application: `go run cmd/myapp/main.go`

You should see output similar to ...

<!-- TODO: explain the imports -->
<!-- TODO: explain go run -->
<!-- TODO: explain the cmd directory (see SGPL) -->

## Install the application

Install and run your application:

1. In the root directory (**tutorial**), run
   `go install myproj/cmd/myapp`.
2. Run the app: `~/go/bin/myapp`

<!-- TODO: explain go install -->

## Learn more

* [Go docs: Go Modules Reference](https://go.dev/ref/mod)
* [Go docs: How to Write Go Code](https://go.dev/doc/code)
* [Go docs: How to Write Go Code (with GOPATH)](https://go.dev/doc/gopath_code)
* [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

Review the following:

* https://golang.org/ref/mod#go-mod-init
* https://golang.org/doc/tutorial/create-module

## Notes

1. [Go Modules Reference: Glossary](https://go.dev/ref/mod#glossary)
