package main

import "testing"

func Test_Reverse_ReturnCorrectAnswer(t *testing.T) {
	testCases := []struct {
		input int
		expected int
	} {
		{123, 321},
		{-123, -321},
		{120, 21},
		{1534236469, 0},
	}

	for _, tc := range testCases {
		t.Run("test", func(t *testing.T) {
			actual := reverse(tc.input)

			if (tc.expected != actual) {
				t.Errorf(`Reverse(%d) = %d, want "%d"`, tc.input, actual, tc.expected)
			} else {
				t.Logf(`Reverse(%d) = %d, want "%d"`, tc.input, actual, tc.expected)
			}
		})
	}
}
