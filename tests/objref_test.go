package test

import (
	"aiowmi/ndr"
	"aiowmi/tools"
	"bytes"
	"testing"
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

func TestObjref(t *testing.T){

	objref_1 := ndr.OBJREF {
		Signature: 1464812877, // = 0x574F454D
		Flags: 0x00000001, // OBJREF_STANDARD
		IID: "4d9f4ab8-7d1c-11cf861e-0020af6e7c57",
		CLSID: "000001a5-0000-0000-c000-000000000046", // CLSID_ActivationContextInfo
	}
	expected_1 := []byte{
		77, 69, 79, 87, // Signature
		1, 0, 0, 0, // Flag
		48, 48, 48, 48, 48, 49, 97, 53, 45, 48, 48, 48, 48, 45, 48, 48, 
		48, 48, 45, 99, 48, 48, 48, 45, 48, 48, 48, 48, 48, 48, 48, 48, 
		48, 48, 52, 54, 52, 100, 57, 102, 52, 97, 98, 56, 45, 55, 100, 
		49, 99, 45, 49, 49, 99, 102, 56, 54, 49, 101, 45, 48, 48, 50,
		48, 97, 102, 54, 101, 55, 99, 53, 55,
	}
	results_1 := objref_1.Get_data()
	
	if  !bytes.Equal(expected_1, results_1){
		t.Errorf("Result - 1 ===>> Expected: %s, Got: %s", expected_1, results_1)
	}
} 