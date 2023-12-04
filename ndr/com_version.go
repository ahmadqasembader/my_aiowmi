package ndr

// MS-DCOM 2.2.11 [COMVERSION]

/*
	The COMVERSION structure is used to specify the major and minor version of either the client or the
	server DCOM Remote Protocol implementation.
*/

import (
	"encoding/binary"
)

var COM_VERSION_MAJOR uint = 7
var COM_VERSION_MINOR uint = 5

// returns major and midor version of COM in little endian byte format. 
func GetCOMVersion()[]byte {
	var buffer []byte

	value := binary.LittleEndian.Uint16([]byte{byte(COM_VERSION_MAJOR), byte(COM_VERSION_MAJOR>>8)})	
	buffer = append(buffer, byte(value), byte(value>>8))

	value = binary.LittleEndian.Uint16([]byte{byte(COM_VERSION_MINOR), byte(COM_VERSION_MINOR>>8)})
	buffer = append(buffer, byte(value), byte(value>>8))
	return buffer
}
