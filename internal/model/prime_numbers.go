package model

import (
	"errors"
	"fmt"
	"math"
)

var ErrParse = errors.New("the given input is invalid")

func isPrimeNumber(n int) bool {
	if n < 2 {
		return false
	}
	sq_root := int(math.Sqrt(float64(n)))
	for i := 2; i <= sq_root; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// IsPrimeNumbers takes a slice of integers and returns a slice of booleans indicating
// whether each integer in the input slice is prime or not.
func IsPrimeNumbers(numbers []int) []bool {
	result := make([]bool, len(numbers))

	for i, v := range numbers {
		result[i] = isPrimeNumber(v)
	}

	return result
}

// ParseNumbers parses a slice of values and returns a slice of integers.
// Each value in the input slice must be either an integer or a float64 that represents an integer.
//
// If a value is not a number, ParseNumbers returns an error with index of first.
//
// If a value is a float, 0 adds to slice, because float can't be a PrimeNumber
func ParseNumbers(data []any) ([]int, error) {
	var result []int = make([]int, 0, len(data))

	for i, v := range data {
		switch v := v.(type) {
		case int:
			result = append(result, int(v))
		case float64:
			if isInt(v) {
				result = append(result, int(v))
			} else {
				result = append(result, 0)
			}
		default:
			return nil, fmt.Errorf("%w. Element on index %d is not a number", ErrParse, i+1)
		}
	}

	return result, nil
}

func isInt(val float64) bool {
	return val == float64(int(val))
}
