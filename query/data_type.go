package query

import (
	"strconv"
	"strings"
	"time"
)

// DataType represent the type of data
type DataType int

// Integer Values.
const (
	DataTypeBoolean DataType = iota + 1

	DataTypeString

	DataTypeDouble

	DataTypeInteger

	DataTypeLong

	DataTypeTime

	DataTypeUInteger

	DataTypeULong
)

// ParseBoolean will try parsing given string to coresponsing boolean type
func ParseBoolean(str string) (bool, bool) {

	switch str {
	case "1":
		fallthrough
	case "true":
		fallthrough
	case "True":
		fallthrough
	case "TRUE":
		return true, true
	case "0":
		fallthrough
	case "false":
		fallthrough
	case "False":
		fallthrough
	case "FALSE":
		return false, true

	}

	return false, false

}

// ParseDouble will try parsing given string to coresponsing double type
func ParseDouble(str string) (float64, bool) {

	s, err := strconv.ParseFloat(str, 64)

	return s, err == nil

}

// ParseInteger will try parsing given string to coresponsing integer type
func ParseInteger(str string) (int, bool) {

	s, err := strconv.Atoi(str)

	return s, err == nil

}

// ParseLong will try parsing given string to coresponsing int64 type
func ParseLong(str string) (int64, bool) {

	rs, err := strconv.ParseInt(str, 10, 64)

	return rs, err == nil

}

// ParseTime will try parsing given string to coresponsing time type
func ParseTime(str string) (*time.Time, bool) {

	rs, valid := ParseLong(str)

	if !valid || rs <= 1000000 {
		return nil, false
	}

	v := time.Unix(rs, 0)

	return &v, true
}

// ParseInteger will try parsing given string to coresponsing uint type
func ParseUInteger(str string) (uint, bool) {
	s, err := strconv.ParseUint(str, 10, 32)
	return uint(s), err == nil
}

// ParseULong will try parsing given string to coresponsing uint64 type
func ParseULong(str string) (uint64, bool) {
	rs, err := strconv.ParseUint(str, 10, 64)
	return rs, err == nil

}

func ParseValue(typ DataType, str string) (interface{}, bool) {
	switch typ {
	case DataTypeString:
		return replacer.Replace(str), true
	case DataTypeBoolean:
		return ParseBoolean(str)
	case DataTypeInteger:
		return ParseInteger(str)
	case DataTypeLong:
		return ParseLong(str)
	case DataTypeDouble:
		return ParseDouble(str)
	case DataTypeTime:
		return ParseTime(str)
	case DataTypeUInteger:
		return ParseUInteger(str)
	case DataTypeULong:
		return ParseULong(str)
	}

	return nil, false
}

var replacer = strings.NewReplacer("'", "''", "%", "\\%")
