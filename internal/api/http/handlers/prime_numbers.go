package handler

import (
	"encoding/json"
	"net/http"
	"primes/internal/model"

	"github.com/labstack/echo"
)

// PrimeNumbersHandler is HTTP request handler checks if each integer is prime or not using the IsPrimeNumbers function from the model package,
// and returns a slice of booleans indicating whether each integer was prime or not.
func PrimeNumbersHandler(ctx echo.Context) error {
	jsonData := ctx.Request().Body
	var data []any

	err := json.NewDecoder(jsonData).Decode(&data)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	result, err := model.IsPrimeNumbers(data)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, result)
}
