# Go cheatsheet

## Slices

See:
* [A Tour of Go: Slices](https://go.dev/tour/moretypes/7)
* [Go by Example: Slices](https://gobyexample.com/slices)
* [yourbasic: Slices/arrays explained: create, index, slice, iterate](https://yourbasic.org/golang/slices-explained/)

## Testing

### Create a table test

An example table test:

```go
func TestMyFunc(t *testing.T) {
	testCases := []struct {
		name     string
		x        int
		y        int
		expected int
	}{
		{"t1", 11, -5, 16},
		{"t2", -11, -5, 6},
		{"t3", 11, 5, 6},
		{"t4", -11, 5, 16},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := tc.expected
			actual := MyFunc(tc.x, tc.y)
			if expected != actual {
				t.Errorf("expected %v; got %v", expected, actual)
			}
		})
	}
}
```

See also: [yourbasic: Table-driven unit tests](https://yourbasic.org/golang/table-driven-unit-test/).

### Create an example test:

An example test:

```go
func ExampleStack() {
	s := Stack{}
	nums := []int{1, 1, 2, 3, 5, 8, 13}
	for _, n := range nums {
		s.Push(n)
	}
	next := s.Iterator()
	// iterate over the items in the stack
	for {
		n, ok := next()
		if !ok {
			break
		}
		fmt.Println(n)
	}
	// Output:
	// 13
	// 8
	// 5
	// 3
	// 2
	// 1
	// 1
}
```

See also: [The Go Blog: Testable Examples in Go](https://go.dev/blog/examples)
