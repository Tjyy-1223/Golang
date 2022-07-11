package main

import (
	"errors"
	"fmt"
)

func add(a, b int) (lastNum int) {
	c := a + b
	return c
}

var err = errors.New("this is my new error")

// 模拟/0错误
var errDivisionByZero = errors.New("division by zero")

func div(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errDivisionByZero
	}

	return dividend / divisor, nil
}

type ParseError struct {
	Filename string
	Line     int
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line)
}

func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}
