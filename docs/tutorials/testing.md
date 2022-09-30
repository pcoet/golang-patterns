# Testing with Go

This tutorial provides an introduction to automated testing in Go. The tutorial
assumes that you have Go installed and that you have some basic familiarity with
Go programming and with software testing.

For help with installing Go, see
[Download and install](https://go.dev/doc/install).

The Go standard library provides a [testing](https://pkg.go.dev/testing) package
to help you test your code. You can use `testing` to create test functions that
run at the `go test` command. The functions must have a signature of the form
`func TestXxx(*testing.T)`, and they must be in files with names ending in
**_test.go**. For example, if you had an `Add` function in an **examples.go**
file, and you wanted to test it, you could create an **examples_test.go** file
with a function `func TestAdd(*testing.t)`.

## Create a Go project

To create a Go project, follow these steps:

1. Create a directory called **testtut**: `mkdir testtut`
2. Change into the directory: `cd testtut`
3. Create a module: `go mod init testtut/example`

## Create an example function

Here's the function that you'll test in this tutorial:

```go
package examples

import (
	"fmt"
	"strconv"
	"strings"
)

func Calculate(input string) (float64, error) {
	var result float64
	strs := strings.Fields(input)
	if len(strs) != 3 {
		return result, fmt.Errorf("expected 3 elements; received %v", len(strs))
	}
	n1, err := strconv.ParseFloat(strs[0], 64)
	if err != nil {
		return result, fmt.Errorf("error converting %v to float", strs[0])
	}
	n2, err := strconv.ParseFloat(strs[2], 64)
	if err != nil {
		return result, fmt.Errorf("error converting %v to float", strs[2])
	}
	switch strs[1] {
	case "+":
		result = n1 + n2
	case "-":
		result = n1 - n2
	case "*":
		result = n1 * n2
	case "/":
		result = n1 / n2
	default:
		return result, fmt.Errorf("unknown operation: %v", strs[1])
	}

	return result, nil
}
```

Add the function to your project:

1. In the **testtut** directory, create a file called **calculate.go**:
   `touch calculate.go`
2. Copy the code above and paste it into **calculate.go**.
3. Save the file.

`Calculate` is a simple calculator function that performs binary arithmetic
operations. It takes an input string, splits the string on white space,
validates the substrings, performs the appropriate operation, and returns the
result. The terms and symbol in the input must be separated by white space.
For example, "2 + 2" is valid input, but "2+2" is not. If you were building a
production calculator application, you might want to handle cases
like "2+2" or "10/5", because users are likely to enter such input. But for the
sake of simplicity, the example function doesn't do that kind of string parsing.

Notice that `Calculate` returns a float and an
error. If the function finishes successfully, it returns the result of the arithmetic operation as a
float and `nil` for the error. If the function doesn't finish successfully, it
returns the zero value for the float, which is 0, and an error. This means that you
always have to check for an error before using the result. Otherwise you can't
tell the difference between an expected 0 (from "2 - 2", for example)
and a zero value.

## A single test

First you'll write a simple test to verify a single output of the function. The
test will verify that `Calculate` can do addition. Here's the test function:

```go
package examples

import "testing"

func TestCalculate(t *testing.T) {
	want := 4.0
	got, _ := Calculate("2 + 2")

	if want != got {
		t.Errorf("expected %v; got %v", want, got)
	}
}
```

Add the test to your project:

1. In the **testtut** directory, create a file called **calculate_test.go**:
   `touch calculate_test.go`
2. Copy the code above and paste it into **calculate_test.go**.
3. Save the file.

Now run the test:

1. Change into the top-level directory of your project: `cd ..`
2. Run `go test ./testtut/examples`. This runs all the tests in the **examples**
   directory.

You should see output similar to `ok      testtut/examples   0.702s`. This means
that your test ran successfully, and your function added two plus two.

The test has a signature of the form `func TestXxx(*testing.T)` and is in a file
with a name ending in **_test.go**, so the `test` command runs it. But because
of the special file name, the test code won't be compiled as part of a build. When you
run the command `go test ./testtut/examples`, `test` runs all the tests in the
`examples` package. You can test multiple packages using `...`. To test all of
the packages in a project, you'd run `go test ./...` from the top directory.
Similarly, if you had a **pkg** directory, you could run all the tests beneath
it using `go test ./pkg/...`.

`TestCalculate` is slightly verbose. It could be rewritten without using the
`want` variable, like this:

```go
func TestCalculate(t *testing.T) {
	got, _ := Calculate("2 + 2")

	if got != 4.0 {
		t.Errorf("expected 4.0; got %v", got)
	}
}
```

But the `want` and `got` variables highlight a convention of Go testing: using
those names to define the expected and actual results of the test. Arguably,
using the `want` variable also makes this test more readable, and it makes the test more maintainable,
in that you can change the expected result in one place rather than
two (the condition `if got != 4.0` and the error message `"expected 4.0; got %v"`).

The test only logs information in the failure case. If the actual and expected
results don't match (`want != got`), `Errorf` logs an error message and marks
the test as failed. If you want to stop execution of the failed function
immediately, you can use [Fatalf](https://pkg.go.dev/testing#T.Fatalf).

## A table test

## An example

## Useful flags

## Additional resources

* [Go Blog: Testable Examples in Go](https://go.dev/blog/examples)
* [yourbasic: Table-driven unit tests](https://yourbasic.org/golang/table-driven-unit-test/)

## Notes
