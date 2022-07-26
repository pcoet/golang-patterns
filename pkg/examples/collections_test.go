package examples

func ExampleDemoSlice() {
	DemoSlice()
	// Output:
	// [aaa bbb ccc ddd eee]
	// [aaa bbb ccc ddd EEE]
	// [aaa bbb ccc ddd EEE fff ggg]
	// 8
	// 14
	// 7
	// 7
	// 0) crow
	// 1) duck
	// 2) seagull
	// 3) pigeon
}

func ExampleDemoMap() {
	DemoMap()
	// Output:
	// map[blue:#0000ff green:#008000 red:#ff0000]
	// map[]
	// map[blue:#0000ff green:#008000 red:#ff0000]
	// 3
	// #0000ff
	// 2
	// false
	// green #008000
	// red #ff0000
}
