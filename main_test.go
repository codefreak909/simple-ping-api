package main

import (
	"errors"
	"testing"
)

func TestDivide(t *testing.T) {
	testcases := []struct {
		inputX, inputY int
		expectedOutput int
		expectedErr    error
	}{
		{inputX: 4, inputY: 2, expectedOutput: 2, expectedErr: nil},
		{inputX: 1, inputY: 2, expectedOutput: 0, expectedErr: nil},
		{inputX: 1, inputY: 0, expectedOutput: 0, expectedErr: errors.New("not divisible by zero")},
	}

	for i := range testcases {
		actualOutput, err := divide(testcases[i].inputX, testcases[i].inputY)
		if actualOutput != testcases[i].expectedOutput {
			t.Errorf("Testcase %d Failed. Expected %d Got %d", i, testcases[i].expectedOutput, actualOutput)
		}

		if err != nil && err.Error() != testcases[i].expectedErr.Error() {
			t.Errorf("Testcase %d Failed. Expected err %v Got %v", i+1, testcases[i].expectedErr.Error(), err.Error())
		}
	}
}
