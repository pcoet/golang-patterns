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
	"fmt"
	"strconv"
	"strings"
)

// Calculate is a simple calculator function that computes binary arithmetic
// operations. It takes input as a string, splits the string on white space,
// validates the substrings, performs the appropriate operation, and returns
// the result. The terms and symbol in the input must be separated by white
// space. For example, "2 + 2" is valid input, but "2+2" is not.
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

// Multiplier takes a float m and returns a function that takes a float n and
// returns m * n.
func Multiplier(m float64) func(float64) float64 {
	return func(n float64) float64 {
		return m * n
	}
}
