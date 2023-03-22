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
		a        []any
		expected []bool
	}{
		{[]any{0, 1, -10, -2, -2000}, []bool{false, false, false, false, false}},
		{[]any{0.0, 4.0, 3.2, 17.0, -13.1}, []bool{false, false, false, true, false}},
		{[]any{999999000001, 999999000002}, []bool{true, false}},
		{[]any{}, []bool{}},
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

func Test_PrimeNumbersHandlerBadRequest(t *testing.T) {
	t.Parallel()

	t.Run("Check BadRequest error", func(t *testing.T) {
		e := echo.New()
		e.POST("/", PrimeNumbersHandler)
		reqBody, err := json.Marshal([]any{1, "2", 3})

		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
