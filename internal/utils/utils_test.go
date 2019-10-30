package utils

import (
	"fmt"
	"testing"
)

func isNil(v interface{}) bool {
	return v == nil
}

func TestParseHeaderKeyValue(t *testing.T) {
	type output struct {
		key          string
		value        string
		errorChecker func(error) bool
	}

	testData := []struct {
		input  string
		output output
	}{
		{
			"header:value",
			output{
				"header",
				"value",
				func(v error) bool { return isNil(v) },
			},
		},
		{
			"a:b",
			output{
				"a",
				"b",
				func(v error) bool { return isNil(v) },
			},
		},
		{
			"a",
			output{
				"",
				"",
				func(v error) bool { return isNil(v) == false },
			},
		},
	}

	for _, testDataOne := range testData {
		key, value, err := ParseHeaderKeyValue(testDataOne.input)

		getTestCaseString := func() func() string {
			var testCase string
			return func() string {
				if len(testCase) == 0 {
					testCase = fmt.Sprintf("Test case %v", testDataOne)
				}
				return testCase
			}
		}()

		if key != testDataOne.output.key {
			t.Error(fmt.Sprintf("The incorrect key. Expected: %v, got %v. %v", testDataOne.output.key, key, getTestCaseString()))
		}
		if !testDataOne.output.errorChecker(err) {
			t.Error(fmt.Sprintf("Error checking feild. %v", getTestCaseString()))
		}
		if value != testDataOne.output.value {
			t.Error(fmt.Sprintf("The incorrect value. Expected: %v, got %v. Test case %v", testDataOne.output.value, value, getTestCaseString()))
		}
	}

}
