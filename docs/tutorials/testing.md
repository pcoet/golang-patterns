# An Introduction to testing with Go

This tutorial provides an introduction to automated testing in Go. You'll create
a Go project, a function, a simple test that verifies a single output for the function, and a
table-driven test that verifies multiple outputs for the function. The tutorial
assumes that you have Go installed and that you have some basic familiarity with
Go programming and with software testing.

For help with installing Go, see
[Download and install](https://go.dev/doc/install).

The Go standard library provides a [testing](https://pkg.go.dev/testing) package
to help you test your code. You can use `testing` to create test functions and
then run the test functions with the `go test` command. The functions must have
a signature of the form `func TestXxx(*testing.T)`, and they must be in files
with names ending in **_test.go**. For example, if you have an `Add` function in
an **arithmetic.go** file, and you want to test it, you can create an
**arithmetic_test.go** file and add a function with the signature
`func TestAdd(*testing.t)`. Then you can use the `test` command to run the test.

## Create a Go project

1. Create a directory called **tutorial** and change into the directory:
   `mkdir tutorial && cd tutorial`
2. Create a module: `go mod init tutorial/calculator`

You'll see a new file, **go.mod**. If you inspect the contents of the file,
you should see something like this:

```go
module tutorial/calculator

go 1.19
```

The Go version depends on your environment. The module is called
`tutorial/calculator` because you're going to create an example function that
performs simple arithmetic.

## Create an example function

Here's the function that you'll test in this tutorial:

```go
package calculator

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

1. In the **tutorial** directory, create a file called **calculate.go**:
   `touch calculate.go`
2. Copy the code above and paste it into **calculate.go**.
3. Save the file.

`Calculate` is a simple calculator function that performs binary arithmetic
operations. It takes an input string, splits the string on white space,
validates the substrings, performs the appropriate operation, and returns the
result. The terms and symbol in the input must be separated by white space.
For example, "2 + 2" is valid input, but "2+2" is not. If you were building a
production calculator application that accepted arbitrary input strings, you
would probably want to handle cases like "2+2" or "10/5", because users are likely to enter such input. But for the
sake of simplicity, the example function doesn't do that kind of string parsing.

Notice that `Calculate` returns a float and an
error. If the function finishes successfully, it returns the result of the arithmetic operation as a
float and `nil` for the error. If the function doesn't finish successfully, it
returns the zero value for the float, which is 0, and an error. This means that you
always have to check for an error before using the result. Otherwise you can't
tell the difference between an expected 0 (from "2 - 2", for example)
and a zero value returned from an error.

## A single test

First you'll write a simple test to verify a single output of the function. The
test will verify that `Calculate` can do addition. Here's the test function:

```go
package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	want := 4.0
	got, _ := Calculate("2 + 2")

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}
```

Add the test to your project and run it:

1. In the **tutorial** directory, create a file called **calculate_test.go**:
   `touch calculate_test.go`
2. Copy the code above and paste it into **calculate_test.go**.
3. Save the file.
4. Run `go test`. This runs all the tests in the current directory.

You should see output similar to the following:

```
PASS
ok  	tutorial/calculator	0.556s
```

This means that your test ran successfully.

`TestAdd` has a signature of the form `func TestXxx(*testing.T)` and is in a file
with a name ending in **_test.go**, so the `test` command runs it. But because
of the special file name, the test code won't be compiled as part of a build.
There are other ways to invoke the `test` tool. For example, to test all of
the packages in a project, you can run `go test ./...` from the top directory.
Similarly, if you had a **pkg** directory, you could run all the tests beneath
it using `go test ./pkg/...`. To learn more about the `test` command, see
[Test packages](https://pkg.go.dev/cmd/go#hdr-Test_packages).

`TestAdd` is slightly verbose. It could be rewritten without using the
`want` variable, like this:

```go
func TestAdd(t *testing.T) {
	got, _ := Calculate("2 + 2")

	if got != 4.0 {
    t.Errorf("got %v; want 4.0", got)
	}
}
```

But the `want` and `got` variables are somewhat conventional in Go testing, with
`want` naming the expected value and `got` naming the actual value. In this case,
using the `want` variable also makes the test more maintainable. If you need to
change the expected result, you only need to update the code in one place, rather
than two (the constant `4.0` in `if got != 4.0` and the string "4.0" in theerror
message `"expected 4.0; got %v"`).

The test only logs information in the failure case. If the actual and expected
results don't match (`want != got`), [Errorf](https://pkg.go.dev/testing#T.Errorf)
logs an error message and marks the test as failed. If you want to stop
execution of the failed function immediately, you can use
[Fatalf](https://pkg.go.dev/testing#T.Fatalf). If you don't need formatted
output, you can use the [Error](https://pkg.go.dev/testing#T.Error) or
[Fatal](https://pkg.go.dev/testing#T.Fatal) methods.

## A table-driven test

<!-- TODO: start here -->

You could write multiple functions like `TestAdd` that verify a single output,
but that would get repetitive. To reduce boilerplate, you can create a
table-driven test:

```go
func TestCalculate(t *testing.T) {
	cases := []struct {
		name  string
		in    string
		want  float64
		isErr bool
	}{
		{"too few fields", "2 +", 0, true},
		{"too many fields", "2 + 2 +", 0, true},
		{"bad first term", "n + 2", 0, true},
		{"bad second term", "2 + n", 0, true},
		{"add", "2 + 2", 4, false},
		{"subtract", "2 - 2", 0, false},
		{"multiply", "2 * 2", 4, false},
		{"divide", "2 / 2", 1, false},
		{"unknown op", "2 # 2", 0, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := Calculate(c.in)
			isErr := err != nil
			if (got != c.want) || (isErr != c.isErr) {
				t.Errorf("got %v, %v; want %v, %v", got, isErr, c.want, c.isErr)
			}
		})
	}
}
```

Add the test to your project and run it:

1. Copy the code above into **calculate_test.go**.
2. Save the file.
3. Run `go test`.

You should see output like before, indicating that the test has passed.

This table-driven test, `TestCalculate`, uses a slice of structs, `cases` to define the inputs and
outputs of tests to be run. This is the test table. The `name` field identifies
the test case. The `in` field holds the input string. The `want` field
specifies the expected result. And the `isErr` field indicates if an error is
expected. It's important to test both the result value and the error status,
because both are needed to verify a result. Because the zero value for a float is 0,
the inputs `0 + 0` and `0 + foo` both cause `Calculate` to return `0` as a first value. But only `0 + foo`
returns a non-nil error.

The test table has a test for each error condition and each operation in
`Calculate`. You could add more tests, but you'd be testing the Go language
more than the `Calculate` function.

Like `TestAdd`, `TestCalculate` checks an actual result against an expected
result and invokes `Errorf` if the results aren't equal. There's also a new
method, [t.Run](https://pkg.go.dev/testing#T.Run), that's invoked for each
element in the test table. `Run` has two parameters: a `name` string that
identifies the test and an anonymous function to manage the tests. Behind the scenes,
`Run` creates a goroutine for each test. In this way, you can use `Run` to create subtests.

You could also create a table-driven test without using `Run`, but the support for named subtests is
useful, as you'll see in the next section.

### Test failures

It's useful to know what to expect when tests fail. In `TestCalculate`, change
the expected output of the addition and multiplication tests, so that instead
of `4` each test expects `5`.

Before:

```go
{"add", "2 + 2", 4, false},
...
{"multiply", "2 * 2", 4, false},
```

After:

```go
{"add", "2 + 2", 5, false},
...
{"multiply", "2 * 2", 5, false},
```

Now, when you run `go test`, these subtests fail with output like this:

```
--- FAIL: TestCalculate (0.00s)
    --- FAIL: TestCalculate/add (0.00s)
        calculate_test.go:37: got 4, false; want 5, false
    --- FAIL: TestCalculate/multiply (0.00s)
        calculate_test.go:37: got 4, false; want 5, false
FAIL
exit status 1
FAIL	tutorial/calculator	0.219s
```

The log formatting highlights the parent/child relationship of `TestCalculate`
to `TestCalculate/add` and `TestCalculate/multiply`. You can see how `Run` uses
the `name` arguments to identify the subtests. If you simply removed `Run` from
your test function and tested each element in the table without it, you'd still
know when a test failed. But the output would look something like this:

```
--- FAIL: TestCalculate (0.00s)
    calculate_test.go:37: got 4, false; want 5, false
    calculate_test.go:37: got 4, false; want 5, false
FAIL
exit status 1
FAIL	tutorial/calculator	0.602s
```

Here you can see that the test returned two errors, but you have to use the
error message to try to identify which test cases have failed. `Run` helps take
the guesswork out of it.

Change the expected outputs back to `4` before going on to the next section.

## Logging

In some cases, you might want to log output from a test run &ndash; for example, if
you're debugging the code under test. The `test` command provides the `-v` flag
for such scenarios.

In `TestAdd`, insert the log statement `t.Logf("got: %v", got)`.

Before:

```go
got, _ := Calculate("2 + 2")

if got != want {
```

After:

```go
got, _ := Calculate("2 + 2")

t.Logf("Logging the actual result: %v", got)

if got != want {
```

Save your changes and run the test with the `-v` flag: `go test -v`. You'll see
verbose output showing the tests that have run and passed. You should also see
the output from the log statement:

```
=== RUN   TestAdd
    calculate_test.go:10: Logging the actual result: 4
--- PASS: TestAdd (0.00s)
```

## Learn more

This tutorial covered the basics of automated testing with Go. To learn more,
see the following resources:

* [Go by Example: Testing](https://gobyexample.com/testing)
* [Go docs: How to Write Go Code: Testing](https://go.dev/doc/code#Testing)
* [testing](https://pkg.go.dev/testing)
* [yourbasic: Table-driven unit tests](https://yourbasic.org/golang/table-driven-unit-test/)
