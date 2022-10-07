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

Notes about the `go mod` command:

* Supports operations on modules.
* `go mod init` creates a **go.mod** file in the current directory.
* The presence of a **go.mod** file in a directory tells Go that the directory
  is the root of a Go module.
* The only argument to `go mod init` is the path of the module to be created, e.g.
  `tutorial/myproj`. The path argument is optional, and `go mod` tries to
  infer the path from existing resources, if you omit it.

Notes about modules:

* Module: collection of packages.
* Go modules are a way to distribute code.
* Modules let you bundle packages into release versions.

TODO: START HERE:
Read this for more info: https://golang.org/ref/mod

## Create a library file

## Create a test

## Create a main file

## Generate documentation

## Learn more

* [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

Review the following:

* https://golang.org/ref/mod#go-mod-init
* https://golang.org/doc/tutorial/create-module
