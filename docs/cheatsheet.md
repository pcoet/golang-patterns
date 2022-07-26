# Go cheatsheet

## Maps

```go
func DemoMap() {
	// create a map literal a with string keys and string values
	cm := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"blue":  "#0000ff",
	}
	fmt.Println(cm) // map[blue:#0000ff green:#008000 red:#ff0000]

	// create an empty map with make
	cm = make(map[string]string)
	fmt.Println(cm) // map[]

	// add elements to the map
	cm["red"] = "#ff0000"
	cm["green"] = "#008000"
	cm["blue"] = "#0000ff"
	fmt.Println(cm) // map[blue:#0000ff green:#008000 red:#ff0000]

	// get the length of the map
	fmt.Println(len(cm)) // 3

	// get a value from the map
	fmt.Println(cm["blue"]) // #0000ff

	// delete an element
	delete(cm, "blue")
	fmt.Println(len(cm)) // 2

	// try to get a deleted element
	_, ok := cm["blue"]
	fmt.Println(ok) // false

	// iterate over a sorted map (iteration order is not guaranteed for maps)
	keys := make([]string, 0)
	for k := range cm {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, cm[k])
		// green #008000
		// red #ff0000
	}
}
```

See also:
* [A Tour of Go: Maps](https://go.dev/tour/moretypes/19)
* [Go by Example: Maps](https://gobyexample.com/maps)
* [yourbasic: Maps explained: create, add, get, delete](https://yourbasic.org/golang/maps-explained/)

## Slices

```go
func DemoSlice() {
	// create an array
	a := [7]string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg"}

	// get a slice of an array
	s := a[0:5]
	fmt.Println(s) // [aaa bbb ccc ddd eee]

	// change an element of the slice
	s[4] = "EEE"

	// see that the slice is changed
	fmt.Println(s) // [aaa bbb ccc ddd EEE]

	// see that the underlying array is also changed
	fmt.Println(a) // [aaa bbb ccc ddd EEE fff ggg]

	// create a slice literal
	fibs := []int{1, 1, 2, 3, 5, 8, 13}

	// create a copy of the slice
	fibsCopy := make([]int, len(fibs))
	copy(fibsCopy, fibs)

	// append an item to the slice
	fibs = append(fibs, 21)

	// get the length of the slice
	fmt.Println(len(fibs)) // 8

	// see that the capacity of the slice has doubled
	fmt.Println(cap(fibs)) // 14

	// see that the length and capacity of the copy have not been changed
	fmt.Println(len(fibsCopy)) // 7
	fmt.Println(cap(fibsCopy)) // 7

	// create a dynamically sized slice
	n := 2
	birds := make([]string, n)

	// assign elements to the slice
	birds[0] = "crow"
	birds[1] = "duck"

	// append elements to the slice
	birds = append(birds, "seagull", "pigeon")

	// iterate over the slice
	for i, bird := range birds {
		fmt.Printf("%v) %v\n", i, bird)
		// 0) crow
		// 1) duck
		// 2) seagull
		// 3) pigeon
	}
}
```

See also:
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
