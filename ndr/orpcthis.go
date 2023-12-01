package ndr

// MS-DCOM 2.2.13 [ORPCTHIS and ORPCTHAT]

import "aiowmi"

type ORPCTHIS struct{
	flags int
	reserved int
	extensions int
	cid []byte
}
 
func (orpcthis *ORPCTHIS) from_data(flags int) []byte{

	// EXT32 := "<L"
	// EXT64 := "<Q"

	orpcthis.flags = 0
	orpcthis.reserved = 0
	orpcthis.cid = main.GenCID()
	orpcthis.extensions = 0

	
	var buffer []byte
	buffer = append(buffer, byte(orpcthis.flags))
	buffer = append(buffer	, byte(orpcthis.reserved))

	// TODO: return the extensions in byte type with the buffer.

	return buffer
}