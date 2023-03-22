package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPrimeNumber(t *testing.T) {
	tests := []struct {
		a        int
		expected bool
	}{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{23, true},
		{1000, false},
		{-10, false},
	}

	for _, test := range tests {
		test := test
		t.Run(strconv.Itoa(test.a), func(t *testing.T) {
			t.Parallel()

			result := isPrimeNumber(test.a)

			assert.Equal(t, test.expected, result)
		})
	}
}

func Test_IsPrimeNumbers(t *testing.T) {
	tests := []struct {
		a        []any
		expected []bool
	}{
		{[]any{0, 1, -10, -1000, 100}, []bool{false, false, false, false, false}},
		{[]any{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}, []bool{true, true, true, true, true, true, true, true, true, true}},
		{[]any{997.23, 983.17, 977.01, 971.0, 967.1}, []bool{false, false, false, true, false}},
		{[]any{8191, 131071, 524287, 6700417, 2147483647}, []bool{true, true, true, true, true}},
		{[]any{999999000001, 67280421310721}, []bool{true, true}},
		{[]any{922337203687, 922337217019, 922337217021}, []bool{true, true, false}},
	}

	for _, test := range tests {
		test := test
		testName := strings.Trim(strings.Join(strings.Split(fmt.Sprint(test.a), " "), ","), "[]")
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			result, err := IsPrimeNumbers(test.a)

			if assert.NoError(t, err) {
				assert.Equal(t, test.expected, result)
			}
		})
	}
}

func Test_IsPrimeNumbersError(t *testing.T) {
	t.Parallel()

	t.Run("Check ErrParse error", func(t *testing.T) {
		_, err := IsPrimeNumbers([]any{1, "2", 3})

		if !errors.Is(err, ErrParse) {
			t.Errorf("wrong error: %v", err)
		}
	})
}
