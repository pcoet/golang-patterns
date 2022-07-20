package examples

import "fmt"

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
