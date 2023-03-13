package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func Test_PrimeNumbersHandler(t *testing.T) {
	tests := []struct {
		a        []int
		expected []bool
	}{
		{[]int{0, 1, -10, -2, -2000}, []bool{false, false, false, false, false}},
		{[]int{2, 3, 17, 13}, []bool{true, true, true, true}},
		{[]int{999999000001, 999999000002}, []bool{true, false}},
		{[]int{}, []bool{}},
	}

	for _, test := range tests {
		test := test
		testName := strings.Trim(strings.Join(strings.Split(fmt.Sprint(test.a), " "), ","), "[]")
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			e := echo.New()
			e.POST("/", PrimeNumbersHandler)
			reqBody, err := json.Marshal(test.a)

			assert.NoError(t, err)
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)
			var actual []bool
			err = json.Unmarshal(rec.Body.Bytes(), &actual)

			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		})

	}

}
