package ndr

import "encoding/binary"
import "aiowmi/tools"

type OBJREFSTD struct {
	std_flags    uint32
	c_public_ref uint32
	oxid         uint64
	oid          uint64
	ipid         string // IPID is a 16 byte data structure, uint64 will overflow
}

func (objref_std *OBJREFSTD) From_data(data []byte, offset int) *OBJREFSTD {
	objref_std_size := 24

	// unpack the following
	// std_falgs, c_public_ref, oxid, oid
	objref_std.std_flags 	=	binary.LittleEndian.Uint32(data[ : 4])
	objref_std.c_public_ref = 	binary.LittleEndian.Uint32(data[4 : 8])
	objref_std.oxid 		= 	binary.LittleEndian.Uint64(data[8 : 16])
	objref_std.oid 			= 	binary.LittleEndian.Uint64(data[16 : 24])

	if objref_std.c_public_ref == 0 {
		panic(0)
	}

	offset += objref_std_size

	objref_std.ipid = tools.Bin_to_str(data, offset)
	offset += CLSID_SIZE

	return objref_std
}