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

			td := &TrialDivision{}
			result := td.isPrimeNumber(test.a)

			assert.Equal(t, test.expected, result)
		})
	}
}

func Test_IsPrimeNumbers(t *testing.T) {
	tests := []struct {
		a        []int
		expected []bool
	}{
		{[]int{0, 1, -10, -1000, 100}, []bool{false, false, false, false, false}},
		{[]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}, []bool{true, true, true, true, true, true, true, true, true, true}},
		{[]int{997, 983, 977, 971, 967}, []bool{true, true, true, true, true}},
		{[]int{8191, 131071, 524287, 6700417, 2147483647}, []bool{true, true, true, true, true}},
		{[]int{999999000001, 67280421310721}, []bool{true, true}},
		{[]int{922337203687, 922337217019, 922337217021}, []bool{true, true, false}},
	}

	for _, test := range tests {
		test := test
		testName := strings.Trim(strings.Join(strings.Split(fmt.Sprint(test.a), " "), ","), "[]")
		t.Run("Trial Division: "+testName, func(t *testing.T) {
			t.Parallel()

			result := IsPrimeNumbers(&TrialDivision{}, test.a)

			assert.Equal(t, test.expected, result)
		})
	}
}

func Test_ParseNumbers(t *testing.T) {
	tests := []struct {
		a        []any
		expected []int
	}{
		{[]any{0, 1, -10, -1000, 100}, []int{0, 1, -10, -1000, 100}},
		{[]any{2, 3, 5, 7, 11.01}, []int{2, 3, 5, 7, 0}},
		{[]any{922337203687, 922337217019, 922337217021}, []int{922337203687, 922337217019, 922337217021}},
	}

	for _, test := range tests {
		test := test
		testName := strings.Trim(strings.Join(strings.Split(fmt.Sprint(test.a), " "), ","), "[]")
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			result, err := ParseNumbers(test.a)

			if assert.NoError(t, err) {
				assert.Equal(t, test.expected, result)
			}
		})
	}
}

func Test_ParseNumbersError(t *testing.T) {
	t.Parallel()

	t.Run("Check ErrParse error", func(t *testing.T) {
		_, err := ParseNumbers([]any{1, "2", 3})

		if !errors.Is(err, ErrParse) {
			t.Errorf("wrong error: %v", err)
		}
	})
}