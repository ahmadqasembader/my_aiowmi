package ndr

// See MS-DCOM 2.2.18 [ OBJREF Flags ]

import (
	"aiowmi/tools"
	"encoding/binary"
	"errors"
)

const OBJREF_SIZE = 20
const CLSID_SIZE = 16

const (
	FLAGS_OBJREF_STANDARD = 0x00000001
	FLAGS_OBJREF_HANDLER  = 0x00000002
	FLAGS_OBJREF_CUSTOM   = 0x00000004
	FLAGS_OBJREF_EXTENDED = 0x00000008
)

type OBJREF struct {
	//General attributes
	Signature uint32
	Flags     uint32
	IID       string
	CLSID 	  string
}


func (objref *OBJREF) Get_data() [] byte{
	buffer := make([]byte, 0)
	data   := make([]byte, 8)

	binary.LittleEndian.PutUint32(data[ : 4], objref.Signature)
	binary.LittleEndian.PutUint32(data[4 : 8], objref.Flags)

	buffer = append(buffer, data...)

	// convert them to bytes in order to be able to append to the buffer
	clsid  := []byte(objref.CLSID)
	iid    := []byte(objref.IID)

	buffer = append(buffer, clsid...)
	buffer = append(buffer, iid...)
	return buffer
}


// // ReadObjRef reads an OBJREF from binary data at a given offset.
func (objrefstd *OBJREF) Read_OBJREF(data []byte, offset int) error {
	if len(data)-offset < OBJREF_SIZE {
		return errors.New("insufficient data for unpacking OBJREF")
	}

	// Unpack the data with the appropriate offset 
	objrefstd.Signature = binary.LittleEndian.Uint32(data[offset : offset+4]) // 4 bytes, singature always = 0x574f454d
	objrefstd.Flags = binary.LittleEndian.Uint32(data[offset+4 : offset+8]) // 4 bytes
	offset += 8

	objrefstd.IID = tools.Bin_to_str(data, offset)
	offset += CLSID_SIZE

	return nil
}
