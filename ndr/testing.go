package ndr

import "fmt"

func orpcthatTest() error{	
	orpcthatData := []byte{1, 0, 0, 0, 2, 0, 0, 0}
	expectedData := ORPCTHAT{flags: 1, extensions: 2}

	orpcthat := ORPCTHAT{}

	offset, err := orpcthat.from_data(orpcthatData, 0)

	if err != nil {
		return fmt.Errorf("Unexpected error: %v", err)
	}
	if orpcthat !=  expectedData{
		fmt.Errorf("Unexpected result. Got %+v, expected %+v", orpcthat, expectedData)
	}
	if offset != 8 {
		return fmt.Errorf("Offset is incorrect: %d -- %v", offset, err)
	}
	
	return nil
}