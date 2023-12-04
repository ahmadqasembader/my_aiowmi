package test

import (
	"aiowmi/ndr"
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
