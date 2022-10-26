# How to create a Go project

**Table of contents**

* [Before you begin](#before-you-begin)
  * [Assumptions](#assumptions)
  * [Prerequisites](#prerequisites)
* [Create a module](#create-a-module)
* [Create a package](#create-a-package)
* [Test a package](#test-a-package)
* [Create a main program](#create-a-main-program)

This tutorial walks you through the process of creating a Go project. By
working through the tutorial, you'll accomplish the following objectives:

1. Set up a project structure that you can use for Go development
2. Understand the layout and packaging of a typical Go project
3. Learn about some of the tools available for managing Go source code

<a id="before-you-begin"></a>
## Before you begin

Before you begin, review the assumptions and prerequisites.

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

To complete this tutorial, you need to have Go installed. To install Go, see
[Download and install](https://go.dev/doc/install).

<a id="create-a-module"></a>
## Create a module

A Go [module](https://go.dev/ref/mod#glos-module) is a collection of packages
that can be released and distributed together in versions.

To get started with a new project, create a Go module:

1. Make a directory named **myproj** and change into it:
   `mkdir myproj && cd myproj`
2. Initialize a Go module: `go mod init myproj`

When you run `go mod init`, you create a **go.mod** file in the current
directory. The directory containing the
[go.mod file](https://go.dev/ref/mod#glos-go-mod-file) is, by definition, the
root directory of the module.

The content of your new **go.mod** file should look something like this:

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
repository. For example, if you were planning to distribute a Go module from
GitHub under the user/repo `myuser/myproj`, you could use
`github.com/myuser/myproj` as both the module path and the repository root
path, and you'd initialize your module with the command
`go mod init github.com/myuser/myproj`.

<a id="create-a-package"></a>
## Create a package

A [package](https://go.dev/ref/mod#glos-package) is a collection of source files
that are located in the same directory and compiled together. For this tutorial,
you'll create a package under the **pkg** directory. This is a convention
from the
[Standard Go Project Layout](https://github.com/golang-standards/project-layout)
(SGPL), which is the model for the project layout in this tutorial.
[pkg](https://github.com/golang-standards/project-layout/tree/master/pkg)
should contain Go libraries intended for use by external consumers.

Create an example package:

1. In the root directory (**myproj**), create a **pkg/mypack/multiplier.go**
   file: `mkdir -p pkg/mypack && touch pkg/mypack/multiplier.go`
2. Copy the following code into **multiplier.go** and save the file.
   ```go
   package mypack

   func Multiplier(m float64) func(float64) float64 {
     return func(n float64) float64 {
       return m * n
     }
   }
   ```
3. (Optional) Format the file: `gofmt pkg/mypack/multiplier.go`.
   [gofmt](https://pkg.go.dev/cmd/gofmt) is a tool for formatting Go source code.
   Depending on the configuration of your editor or IDE, the file might be
   automatically formatted on save, in which case you don't need to run this
   command.

The function that you just created, `Multiplier`, takes a float, `m`, and
returns another function that takes a float, `n`, and returns the result of
`m * n`. `Multiplier` implements a
[function closure](https://go.dev/tour/moretypes/25) and demonstrates Go support
for
[higher-order functions](https://en.wikipedia.org/wiki/Higher-order_function).
Although function closures and higher-order functions are not directly related
to this tutorial, they're neat features and worth knowing about.

<a id="test-a-package"></a>
## Test a package

Go provides the [testing](https://pkg.go.dev/testing) package to help you write
unit tests and the `go test` command to run them.

Create and run a test:

1. In the root directory (**myproj**), create a **pkg/multiplier_test.go** file:
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
4. Run the tests in the **pkg** directory:
   ```
   go test ./pkg/...
   ```

You should see output similar to `ok  	myproj/pkg/mypack	0.626s`. This means
that your test ran successfully.

The `TestMultiplier` function uses `Multiplier` to create a `double` function
that multiplies an input float by 2 and returns the result. The test follows a
common Go pattern of naming the expected output `want` and the actual output
`got`. If these values are not equal, `T.Errorf` is invoked and the test fails.

The signature of `TestMultiplier` is important. The `go test` command runs
functions that have a signature of the form `func TestXxx(t *testing.T)` and
that are located in files with filenames ending in **_test.go**. To learn more
about Go testing, see
[An introduction to testing in Go](https://github.com/pcoet/golang-patterns/blob/main/docs/tutorials/testing.md).

<a id="create-a-main-program"></a>
## Create a main program

If you only want to create a module for other packages to consume, you already
have the basic structure for your project. But if want to create a client to
to consume your own package, you need a main file. Your project could even have
multiple main files, each for a different client.

Create and run a main file:

1. In the root directory (**myproj**), create a **cmd/myapp/main.go** file:
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

Your `main` function should output `4`.

This is a pretty simple program. Like the test you created earlier, the `main`
function creates a `double` function and uses that to double a number. The
result is printed to standard output using `fmt.Println`. But there are a couple
of things to note. First, there's the `import` statement, which imports
[fmt](https://pkg.go.dev/fmt) from the standard library and also imports your
`mypack` package. This is how you use package in a Go application.

The other thing to note is the structure of your **main.go** file, which
contains a `main` function in a `main` package. When you use
[go run](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program)
to complie the `main` package and run it, the `main` function is invoked.

The directory structure for the `main` package follows the SGPL convention of
using a [cmd](https://github.com/golang-standards/project-layout/tree/master/cmd)
directory for a project's main applications. In this case, you only have one
application, but in more complex projects you might have multiple client
applications.

## Install the application

<!-- TODO: work through previous material and then START HERE -->

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
* https://go.dev/doc/code

