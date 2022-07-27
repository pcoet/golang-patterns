// Copyright 2022 Google LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

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
