package requestparameter

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mis-hashemi/request-parameter/query"
)

func ParseEchoQueryString(c echo.Context, paramInfoMap map[string]query.RequestParameter) (query.QueryInfo, error) {
	var queries []query.Query

	for param, paramInfo := range paramInfoMap {
		value := c.QueryParam(param)
		if value == "" {
			if !paramInfo.Optional {
				return nil, fmt.Errorf("missing required parameter: %s", param)
			}
			continue // Skip optional parameter
		}

		query, err := parseQuery(param, value, paramInfo.Definition)
		if err != nil {
			return nil, err
		}
		queries = append(queries, query)
	}

	// Assuming all queries are combined with AND
	return query.NewMultipleAndQueryInfo(queries...), nil
}

func parseQuery(param, value string, definition query.QueryDefinition) (query.Query, error) {

	// Split the value into the operator and the value part (e.g., "mroeq:18" -> "mroeq" and "18")
	parts := strings.SplitN(value, ":", 2)
	if len(parts) < 1 {
		return nil, fmt.Errorf("invalid value for parameter %s: %s", param, value)
	}

	operatorPart := strings.TrimSpace(parts[0])
	operator, paramExpectation, err := query.ParseQueryOperator(operatorPart)
	if err != nil {
		return nil, err
	}
	if !isSupportedOperator(operator, definition) {
		return nil, fmt.Errorf("operator %s is not supported for parameter %s", operator, param)
	}
	if paramExpectation == query.ParameterExpectationZero {
		return query.NewEmptyQuery(definition.GetName(), operator), nil
	}

	valuePart := strings.TrimSpace(parts[1])
	parsedValue, ok := query.ParseValue(definition.GetType(), valuePart)
	if !ok {
		return nil, fmt.Errorf("invalid value for parameter %s: %s", param, valuePart)
	}

	return query.NewQuery(definition.GetName(), operator, query.NewOperand(parsedValue)), nil
}

func isSupportedOperator(operator query.QueryOperator, definition query.QueryDefinition) bool {
	supportedOperators := definition.GetSupportedOperators()
	for _, op := range supportedOperators {
		if op == operator {
			return true
		}
	}
	return false
}
