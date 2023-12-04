package ndr

// MS-DCOM 2.2.13 [ORPCTHIS and ORPCTHAT]

 /*
	The ORPCTHIS structure is the first (implicit) argument sent in an ORPC request PDU and is used to
	send ORPC extension data to the server. The ORPCTHIS structure is also sent as an explicit
	argument in activation RPC requests.
 */
 
import (
	"aiowmi/tools"
	"encoding/binary"
)

type ORPCTHIS struct{
	flags int
	reserved int
	extensions int
	cid []byte
}
 
func (orpcthis *ORPCTHIS) From_data(flags int) []byte{

	orpcthis.flags = 0
	orpcthis.reserved = 0
	orpcthis.cid = tools.GenCID()
	orpcthis.extensions = 0
	
	buffer := make([]byte, 0)

	// ensuring a min of 8 byte in the buffer
	buffer = append(buffer, make([]byte, 8)...) 

	binary.LittleEndian.PutUint32(buffer, uint32(orpcthis.flags))
	binary.LittleEndian.PutUint32(buffer, uint32(orpcthis.reserved))
	
	buffer = append(buffer, GetCOMVersion()...)


	// TODO: return the extensions in byte type with the buffer.

	return buffer
}