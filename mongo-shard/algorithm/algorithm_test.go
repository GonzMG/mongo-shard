package algorithm

import (
	"testing"
)

func TestHashAlgorithmCRC32(t *testing.T) {
	testCases := []struct {
		name     string
		inputKey string
		expected int
	}{
		{"Empty input", "", 0},
		{"Input: test1", "test1", 2326977762},
		{"Weird characters", "ññññññ````", 2934175523},
	}

	alg := &crc32Algorithm{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := alg.hash([]byte(tc.inputKey))
			if output != tc.expected {
				t.Errorf("%v | EXPECTED: %v | OUTPUT: %v", tc.name, tc.expected, output)
			}
		})
	}
}

func TestHashAlgorithmCRC16(t *testing.T) {
	testCases := []struct {
		name     string
		inputKey string
		expected int
	}{
		{"Empty input", "", 65535},
		{"Input: test1", "test1", 940},
		{"Weird characters", "ññññññ````", 38400},
	}

	alg := &crc16Algorithm{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := alg.hash([]byte(tc.inputKey))
			if output != tc.expected {
				t.Errorf("%v | EXPECTED: %v | OUTPUT: %v", tc.name, tc.expected, output)
			}
		})
	}
}
