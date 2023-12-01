package ndr

// MS-DCOM 2.2.11 [COMVERSION]

var COM_VESION_MAJOR uint = 7
var COM_VESION_MINOR uint = 5

func GetCOMVersion()[]byte {
	var buffer []byte
	buffer = append(buffer, byte(COM_VESION_MAJOR), byte(COM_VESION_MAJOR>>8))
	buffer = append(buffer, byte(COM_VESION_MINOR), byte(COM_VESION_MINOR>>8))
	return buffer
}
