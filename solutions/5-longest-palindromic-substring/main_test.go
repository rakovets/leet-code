package main

import "testing"
func Test_longestPalindrome(t *testing.T) {
    testCases := []struct{
		input string
		expected string
	} {
		{"a", "a"},
		{"ac", "a"},
		{"bb", "bb"},
		{"cbbd", "bb"},
		{"babad", "bab"},
		{"azwdzwmwcqzgcobeeiphemqbjtxzwkhiqpbrprocbppbxrnsxnwgikiaqutwpftbiinlnpyqstkiqzbggcsdzzjbrkfmhgtnbujzszxsycmvipjtktpebaafycngqasbbhxaeawwmkjcziybxowkaibqnndcjbsoehtamhspnidjylyisiaewmypfyiqtwlmejkpzlieolfdjnxntonnzfgcqlcfpoxcwqctalwrgwhvqvtrpwemxhirpgizjffqgntsmvzldpjfijdncexbwtxnmbnoykxshkqbounzrewkpqjxocvaufnhunsmsazgibxedtopnccriwcfzeomsrrangufkjfzipkmwfbmkarnyyrgdsooosgqlkzvorrrsaveuoxjeajvbdpgxlcrtqomliphnlehgrzgwujogxteyulphhuhwyoyvcxqatfkboahfqhjgujcaapoyqtsdqfwnijlkknuralezqmcryvkankszmzpgqutojoyzsnyfwsyeqqzrlhzbc", "sooos"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := longestPalindrome(tc.input)
			if (tc.expected != actual) {
				t.Errorf(`longestPalindrome("%s") = "%s", want "%s"`, tc.input, actual, tc.expected)
			} else {
				t.Logf(`longestPalindrome("%s") = "%s", want '%s'`, tc.input, actual, tc.expected)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
    testCases := []struct{
		input string
		expected bool
	} {
		{"a", true},
		{"ac", false},
		{"bb", true},
		{"cbb", false},
		{"aba", true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := isPalindrome(tc.input)
			if (tc.expected != actual) {
				t.Errorf(`isPalindrome("%s") = "%v", want "%v"`, tc.input, actual, tc.expected)
			} else {
				t.Logf(`isPalindrome("%s") = "%v", want '%v'`, tc.input, actual, tc.expected)
			}
		})
	}
}