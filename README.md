
# Request Parameter Library

The Request Parameter Library is a Go package designed to handle query parameters in a structured and flexible manner. It provides a set of data types and operators for querying data efficiently.

## Table of Contents

- [Request Parameter Library](#request-parameter-library)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Data Types](#data-types)
  - [Query Operators](#query-operators)
  - [Usage](#usage)
  - [Contributing](#contributing)

## Introduction

The "request-parameter" library is a Go package that simplifies the handling of query parameters in various applications, especially those involving data querying and filtering. It defines a set of data types and query operators to create structured queries for your data.

## Data Types

The library supports the following data types:

- Boolean
- String
- Double
- Integer
- Long
- Time
- Unsigned Integer
- Unsigned Long

## Query Operators

The library provides the following query operators:

- Equal (eq)
- Not Equal (neq)
- More Than (mr)
- Equal or More Than (mroeq)
- Less Than (ls)
- Equal or Less Than (lsoeq)
- Contain (cn)
- Not Contain (ncn)
- In (in)
- Not In (nin)
- Empty (empt)
- Not Empty (nempt)

For a complete list of query operators and their descriptions, see the [Query Operators](#query-operators) section.

## Usage

To use the "request-parameter" library in your Go project, you can follow these steps:

1. Install the library in your project:

   ```shell
   go get github.com/mis-hashemi/request-parameter

## Examples

Here are some examples of how to use the "request-parameter" library:

```
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

```

```
localhost:8080/query?name=cn:fatemeh,ali,sara&age=mroeq:18
```

For more advanced usage and integration details, please refer to the library's documentation or code examples in the "examples" directory.

## Contributing

Contributions to the "request-parameter" library are welcome. If you have ideas for improvements, bug fixes, or new features, please feel free to open an issue or submit a pull request. Make sure to follow the project's code of conduct.