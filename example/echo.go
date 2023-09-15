package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	requestparameter "github.com/mis-hashemi/request-parameter"
	"github.com/mis-hashemi/request-parameter/query"
)

func main() {
	e := echo.New()
	// Define the query parameter information.
	paramInfoMap := map[string]query.RequestParameter{
		"age": {
			Definition: query.NewQueryDefinition("age", query.GetAllQueryOperator(), query.DataTypeInteger),
			Optional:   false,
		},
		"name": {
			Definition: query.NewQueryDefinition("name", query.GetAllQueryOperator(), query.DataTypeString),
			Optional:   true,
		},
	}

	e.GET("/query", func(c echo.Context) error {

		// Parse the query parameters.
		queryInfo, err := requestparameter.ParseEchoQueryString(c, paramInfoMap)
		if err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
		}

		// Use the queryInfo to filter your data or perform other operations.
		// Here, we're just returning the query as JSON for demonstration purposes.
		return c.JSON(http.StatusOK, queryInfo)
	})

	e.Start(":8080")
}
