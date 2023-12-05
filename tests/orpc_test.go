package test

import (
	"aiowmi/ndr"
	"aiowmi/tools"
	"bytes"
	"fmt"
	"testing"
)

func TestOrpcthat(t *testing.T){	
	orpcthat := ndr.ORPCTHAT{}


	orpcthatData := []byte{1, 0, 0, 0, 2, 0, 0, 0}
	expectedData := ndr.ORPCTHAT{Flags: 1, Extensions: 2}


	offset, err := orpcthat.From_data(orpcthatData, 0)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		panic(err)
	}
	if orpcthat !=  expectedData{
		t.Errorf("Unexpected result. Got %+v, expected %+v", orpcthat, expectedData)
	}
	if offset != 8 {
		t.Errorf("Offset is incorrect: %d -- %v", offset, err)
		panic(err)
	}
}

// The FixedCID function is used instead of GenCID for testing purposes
// The test will fail if you don't fix the CID in the orpcthis.go as well
func TestOrpcthis(t *testing.T) {
	// Test case 1: Flags = 42
	flags1 := 43

	orpcthis1 := ndr.ORPCTHIS{}
	result1 := orpcthis1.From_data(flags1)

	expectedData1 := append([]byte{0, 0, 0, 0, 0, 0, 0, 0}, ndr.GetCOMVersion()...)
	expectedData1 = append(expectedData1, tools.FixedCID()...)
	
	if !bytes.Equal(result1, expectedData1) {
		t.Errorf("Unexpected result for test case 1. Got %+v, expected %+v", result1, expectedData1)
	}else{
		fmt.Printf("Expected Output: %d - Acutal Output: %d \n", expectedData1, result1)
	}

	// Test case 2: Flags = 0
	flags2 := 0

	orpcthis2 := ndr.ORPCTHIS{}
	result2 := orpcthis2.From_data(flags2)

	
	expectedData2 := append([]byte{0, 0, 0, 0, 0, 0, 0, 0}, ndr.GetCOMVersion()...)
	expectedData2 = append(expectedData2, tools.FixedCID()...)

	if !bytes.Equal(result2, expectedData2) {
		t.Errorf("Unexpected result for test case 2. Got %+v, expected %+v", result2, expectedData2)
	}else{
		fmt.Printf("Expected Output: %d - Acutal Output: %d \n", expectedData2, result2)
	}

}
