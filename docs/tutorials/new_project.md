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

The following definitions come from the
[Go Modules Reference Glossary](https://go.dev/ref/mod#glossary) (a few of the
entries have been lightly edited):

* **module:** A collection of packages that are released, versioned, and
  distributed together.
* **module path:** A path that identifies a module and acts as a prefix for
  package import paths within the module. For example, `"golang.org/x/net"`.
* **go.mod file:** The file that defines a module’s path, requirements, and
  other metadata. Appears in the module’s root directory.
* **module root directory:** The directory that contains the go.mod file that
  defines a module.
* **main module:** The module in which the go command is invoked. The main
  module is defined by a go.mod file in the current directory or a parent
  directory.
* **package:** A collection of source files in the same directory that are
  compiled together.
**import path:** A string used to import a package in a Go source file.
  Synonymous with package path.
* **package path:** The path that uniquely identifies a package. A package path
  is a module path joined with a subdirectory within the module. For example
  `"golang.org/x/net/html"` is the package path for the package in the module
  `"golang.org/x/net"` in the `"html"` subdirectory. Synonym of import path.
* **repository root path:** The portion of a module path that corresponds to a
  version control repository’s root directory.

Notes about the `go mod` command:

* Supports operations on modules.
* `go mod init` creates a **go.mod** file in the current directory.
* The presence of a **go.mod** file in a directory tells Go that the directory
  is the root of a Go module.
* The only argument to `go mod init` is the path of the module to be created, e.g.
  `tutorial/myproj`. The path argument is optional, and `go mod` tries to
  infer the path from existing resources, if you omit it.

## Create a library file

<!-- TODO: develop an example library function (what? something cool...) -->
<!-- TODO: start here; use as a reference: https://github.com/pcoet/golang-patterns/tree/main/pkg/examples -->

## Create a test

## Create a main file

## Generate documentation

## Learn more

* [Go docs: Go Modules Reference](https://go.dev/ref/mod)
* [Go docs: How to Write Go Code](https://go.dev/doc/code)
* [Go docs: How to Write Go Code (with GOPATH)](https://go.dev/doc/gopath_code)
* [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

Review the following:

* https://golang.org/ref/mod#go-mod-init
* https://golang.org/doc/tutorial/create-module
