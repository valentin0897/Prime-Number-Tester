package model

import (
	"math"
	"math/big"
)

type PrimeChecker interface {
	isPrimeNumber(num int) bool
}

type TrialDivision struct {
}

func (*TrialDivision) isPrimeNumber(n int) bool {
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

type MillerRabin struct {
}

func (*MillerRabin) isPrimeNumber(n int) bool {
	num := big.NewInt(int64(n))
	return num.ProbablyPrime(15)
}

// IsPrimeNumbers takes a slice of integers and returns a slice of booleans indicating
// whether each integer in the input slice is prime or not.
func IsPrimeNumbers(alg PrimeChecker, numbers []int) []bool {
	result := make([]bool, len(numbers))

	for i, v := range numbers {
		result[i] = alg.isPrimeNumber(v)
	}

	return result
}
