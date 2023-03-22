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
func IsPrimeNumbers(numbers []any) ([]bool, error) {
	result := make([]bool, 0, len(numbers))

	for i, v := range numbers {
		switch v := v.(type) {
		case int:
			result = append(result, isPrimeNumber(v))
		case float64:
			if isInt(v) {
				result = append(result, isPrimeNumber(int(v)))
			} else {
				result = append(result, false)
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
