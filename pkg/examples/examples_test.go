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

func TestMultiplier(t *testing.T) {
	double := Multiplier(2)
	want := 20.0
	got := double(10)

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}
