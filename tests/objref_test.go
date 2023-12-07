package test

import (
	"testing"
	"aiowmi/tools"
)

func TestBinToStr(t *testing.T) {
	// Test Case 1: Basic Test
	binaryData1 := []byte{
		0x78, 0x56, 0x34, 0x12, // Little-endian uint32
		0x02, 0x01, // Little-endian uint16
		0x04, 0x03, // Little-endian uint16
		0x0A, 0x0B, // Big-endian uint16
		0x0C, 0x0D, // Big-endian uint16
		0x0E, 0x0F, 0x10, 0x11, // Big-endian uint32
	}
	expectedUUID1 := "12345678-0102-0304-0A0B-0C0D0E0F1011"

	// Test Case 2: All Zeros
	binaryData2 := make([]byte, 16)
	expectedUUID2 := "00000000-0000-0000-0000-000000000000"

	// Test Case 3: Maximum Values
	binaryData3 := []byte{
		0xFF, 0xFF, 0xFF, 0xFF, // Little-endian uint32 (max value)
		0xFF, 0xFF, // Little-endian uint16 (max value)
		0xFF, 0xFF, // Little-endian uint16 (max value)
		0xFF, 0xFF, // Big-endian uint16 (max value)
		0xFF, 0xFF, // Big-endian uint16 (max value)
		0xFF, 0xFF, 0xFF, 0xFF, // Big-endian uint32 (max value)
	}
	expectedUUID3 := "FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF"

	// Test Case 4: Random Values
	binaryData4 := []byte{
		0x01, 0x23, 0x45, 0x67, // Little-endian uint32
		0xAB, 0xCD, // Little-endian uint16
		0xEF, 0x98, // Little-endian uint16
		0x76, 0x54, // Big-endian uint16
		0x32, 0x10, // Big-endian uint16
		0xFE, 0xDC, 0xBA, 0x98, // Big-endian uint32
	}
	expectedUUID4 := "67452301-CDAB-98EF-7654-3210FEDCBA98"

	testCases := []struct {
		name     string
		binary   []byte
		expected string
	}{
		{"Basic Test", binaryData1, expectedUUID1},
		{"All Zeros", binaryData2, expectedUUID2},
		{"Maximum Values", binaryData3, expectedUUID3},
		{"Random Values", binaryData4, expectedUUID4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tools.Bin_to_str(tc.binary, 0)
			if result != tc.expected {
				t.Errorf("Expected: %s, Got: %s", tc.expected, result)
			}
		})
	}
}
