# An Introduction to testing with Go

This tutorial provides an introduction to automated testing in Go. As you work
through the tutorial, you'll create the following resources:

* A Go project
* A calculator function that performs simple arithmetic operations
* A test that verifies a single output for the function
* A table-driven test that verifies multiple outputs for the function

The tutorial assumes that you have Go installed and that you have a basic
familiarity with Go programming and with software testing.

For help installing Go, see [Download and install](https://go.dev/doc/install).

The Go standard library provides a [testing](https://pkg.go.dev/testing) package
to help you test your code. You can use `testing` to create test functions and
then run the test functions with the `go test` command. The functions must have
a signature of the form `func TestXxx(*testing.T)`, and they must be in files
with names ending in **_test.go**. For example, if you had an `Add` function in
an **arithmetic.go** file, and you wanted to test it, you could create an
**arithmetic_test.go** file and add a function with the signature
`func TestAdd(*testing.t)`. Then you could use the `go test` command to run the
test. You'll set up a similar test harness in the following sections.

## Create a project

First, create a Go project:

1. Make a directory called **tutorial** and change into the directory:
   `mkdir tutorial && cd tutorial`
2. Initialize a module: `go mod init tutorial/calculator`

You should see a new file, **go.mod**. If you inspect the contents of the file,
you should see something like this:

```go
module tutorial/calculator

go 1.19
```

The Go version depends on your environment.

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
2. Copy the code above, including the package and import statements, and paste
   it into **calculate.go**.
3. Save the file.

`Calculate` performs an arithmetic operation (addition, subtraction,
multiplication, or division) on two numbers. It takes an input string, splits
the string on white space, validates the substrings, performs the
specified operation, and returns the result. The
terms and symbol in the input must be separated by white space. For example,
"2 + 2" is valid input, but "2+2" is not. If you were building a production
calculator application that accepted arbitrary input strings, you would probably
want to handle inputs like "2 + 2*3" or "10/5", because users are likely to
enter them. But for the sake of simplicity, the example function doesn't parse
such strings.

Notice that `Calculate` returns a float and an error. If the function finishes
successfully, it returns the result of the arithmetic operation as a
float and `nil` for the error. If the function doesn't finish successfully, it
returns the zero value for the float, which is 0, and an error. This means that
you should check for an error before using a returned value of `0`. Otherwise
you can't tell the difference between an expected 0 (from "2 - 2", for example)
and a zero value returned from an error.

> Note: To learn more about verifying a zero value using idiomatic Go, see
the discussion of the "comma ok" pattern in
[Effective Go](https://go.dev/doc/effective_go#maps).

## Create a simple test

First you'll write a simple test to verify a single output of the function. The
test will verify that `Calculate` can do addition. To keep things simple, it
won't inspect the error value.

Here's the test function:

```go
package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	want := 4.0
	got, _ := Calculate("2 + 2") // note that the error is discarded

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}
```

Add the test to your project and run it:

1. In the **tutorial** directory, create a file called **calculate_test.go**:
   `touch calculate_test.go`
2. Copy the code above, including the package and import statements, and paste
   it into **calculate_test.go**.
3. Save the file.
4. Run `go test`. This runs all the tests in the current directory.

You should see output similar to the following:

```
PASS
ok  	tutorial/calculator	0.445s
```

This means that your test ran successfully.

`TestAdd` has a signature of the form `func TestXxx(*testing.T)` and is in a
file with a name ending in **_test.go**, so `go test` knows to run it. Because
of the special file name, the test code won't be compiled as part of a build.

You can also specify which packages you want to test with `go test`. If you had
a **pkg** directory, you could run all the tests beneath it using
`go test ./pkg/...`. To test all of the packages in a project, you can run
`go test ./...` from the top directory. To learn more about the `go test`
command, see [Test packages](https://pkg.go.dev/cmd/go#hdr-Test_packages).

`TestAdd` is slightly verbose. It could be rewritten without using the `want`
variable, like this:

```go
func TestAdd(t *testing.T) {
	got, _ := Calculate("2 + 2")

	if got != 4.0 {
    t.Errorf("got %v; want 4.0", got)
	}
}
```

But in this case, using `want` is a better choice. For one thing, it makes the
test more readable to other Go programmers. The `want` and `got` names are
common in Go testing, with the expected value assigned to `want` and the
actual value assigned to `got`. Also, using the `want` variable makes the test
more maintainable. If you need to change the expected result, you only need to
update the code in one place (`want`), rather than two (the constant `4.0` and
the string `"4.0"` in the error message).

If the actual and expected results don't match (`if got != want`),
[Errorf](https://pkg.go.dev/testing#T.Errorf) logs an error message and marks
the test as failed. If you want to stop execution of a failed function
immediately, you can use [Fatalf](https://pkg.go.dev/testing#T.Fatalf). If you
don't need formatted output, you can use the
[Error](https://pkg.go.dev/testing#T.Error) or
[Fatal](https://pkg.go.dev/testing#T.Fatal) methods.

## Create a table-driven test

You could write multiple functions like `TestAdd` that verify a single output,
but that would get repetitive. To reduce boilerplate, you can create a
table-driven test:

```go
func TestCalculate(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want float64
		ok   bool
	}{
		{"too few fields", "2 +", 0, false},
		{"too many fields", "2 + 2 +", 0, false},
		{"bad first term", "n + 2", 0, false},
		{"bad second term", "2 + n", 0, false},
		{"add", "2 + 2", 4, true},
		{"subtract", "2 - 2", 0, true},
		{"multiply", "2 * 2", 4, true},
		{"divide", "2 / 2", 1, true},
		{"unknown op", "2 # 2", 0, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := Calculate(c.in)
			ok := err == nil
			if (got != c.want) || (ok != c.ok) {
				t.Errorf("got %v, %v; want %v, %v", got, ok, c.want, c.ok)
			}
		})
	}
}
```

Add the test to your project and run it:

1. Copy the code above into **calculate_test.go**.
2. Save the file.
3. Run `go test`.

You should see output like before, indicating that the tests have passed.

`TestCalculate` uses a slice of structs, `cases` to define the inputs and
outputs of tests to be run. This is the test table. The `name` field identifies
the test case. The `in` field holds the input string. The `want` field
specifies the expected result. And the `ok` field indicates if `Calculate`
is expected to return without an error. Unlike `TestAdd`, `TestCalculate`
verifies both the float and the error returned by `Calculate`:

```go
			got, err := Calculate(c.in)
			ok := err == nil // no error means the result is ok
			if (got != c.want) || (ok != c.ok) { // check both `got` and `ok`
				t.Errorf("got %v, %v; want %v, %v", got, ok, c.want, c.ok)
			}
```

For some cases, you do need to check both `got` and `ok` to verify the result.
For example, when `TestCalculate` tests the input `"n + 2"`, it expects a `0`
and an `ok` value of `false`. But when it tests `"2 - 2"`, it expects a `0` and
an `ok` value of `true`.

The test table has a test for each error condition and each operation in
`Calculate`. You could add more tests, but you'd soon be testing the Go language
more than the logic of the `Calculate` function.

If the actual and expected results don't match, `TestCalculate` invokes `Errorf`.

The tests happen in the body of the [t.Run](https://pkg.go.dev/testing#T.Run)
method, which is invoked for each element in the test table. `Run` has two
parameters: a `name` string that identifies the test and an anonymous function
to manage the test. Behind the scenes, `Run` creates a goroutine for each test.
In this way, you can use `Run` to create subtests.

You could also create a table-driven test without using `Run`, but the support
for named subtests is useful, as you'll see in the next section.

## Examine test failures

It's useful to know what to expect when tests fail. In `TestCalculate`, change
the expected output of the addition and multiplication tests, so that instead
of `4` each test expects `5`.

Before:

```go
		{"add", "2 + 2", 4, true},
		...
		{"multiply", "2 * 2", 4, true},
```

After:

```go
		{"add", "2 + 2", 5, true}, // 2 + 2 = 5; test will fail
		...
		{"multiply", "2 * 2", 5, true}, // 2 + 2 = 5; test will fail
```

Now, when you run `go test`, these cases fail with output like this:

```
--- FAIL: TestCalculate (0.00s)
    --- FAIL: TestCalculate/add (0.00s)
        calculate_test.go:38: got 4, true; want 5, true
    --- FAIL: TestCalculate/multiply (0.00s)
        calculate_test.go:38: got 4, true; want 5, true
FAIL
exit status 1
FAIL	tutorial/calculator	4.003s
```

You can see how `Run` uses the `name` arguments to identify the subtests:
`TestCalculate/add` and `TestCalculate/multiply`. If you removed `Run` and ran
the tests again, you'd still know when a test failed. But the output would be
less helpful. It would look something like this:

```
--- FAIL: TestCalculate (0.00s)
    calculate_test.go:38: got 4, true; want 5, true
    calculate_test.go:38: got 4, true; want 5, true
FAIL
exit status 1
FAIL	tutorial/calculator	0.603s
```

Here you can see that the test returned two errors, but you have to use the
error message to try to identify which test cases have failed. `Run` removes
the guesswork.

Before going on to the next section, change the expected outputs back to `4`.

## Examine log output from a passing test

In some cases, you might want to log output from a passing test &ndash; for
example, if you're debugging the code under test. The `go test` command provides
the `-v` flag for such scenarios.

In `TestAdd`, insert the log statement
`t.Logf("Logging the actual result: %v", got)`.

```go
func TestAdd(t *testing.T) {
	want := 4.0
	got, _ := Calculate("2 + 2") // note that the error is discarded: `_`

	// NEW CODE: write to the error log
	t.Logf("Logging the actual result: %v", got)

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}
```

Save your changes and run the test with the `-v` flag: `go test -v`. You'll see
verbose output showing the tests that have run and passed. You should also see
the output from the log statement:

```
=== RUN   TestAdd
    calculate_test.go:12: Logging the actual result: 4
--- PASS: TestAdd (0.00s)
```

## Learn more

This tutorial has covered the basics of automated testing with Go. To learn
more, see the following resources:

* [Go by Example: Testing](https://gobyexample.com/testing)
* [Go docs: How to Write Go Code: Testing](https://go.dev/doc/code#Testing)
* [testing](https://pkg.go.dev/testing)
* [yourbasic: Table-driven unit tests](https://yourbasic.org/golang/table-driven-unit-test/)
