package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	value          string
	expectedResult string
	expectedErr    error
	actualResult   string
	actualErr      error
}

func TestHelloWorldBfSuccess(t *testing.T) {
	testCase := TestCase{
		value:          "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>",
		expectedResult: "Hello World!",
	}

	a, _ := readData([]byte(testCase.value))

	assert.Equal(t, testCase.expectedResult, a)
}

func TestHelloWorldBfFail(t *testing.T) {
	testCase := TestCase{
		value:          "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+",
		expectedResult: "Hello World!",
	}

	_, err := readData([]byte(testCase.value))

	assert.Nil(t, testCase.expectedErr, err)
}
