package ndr

import (
	"encoding/binary"
	"aiowmi/tools"
)

type OBJREFStandard struct {
	objref 	  OBJREF
	OXID 	  uint16
	OID 	  uint16
	IPID      string
	Counter   uint16
	//Res_addr  DualStringArray

	std_flags uint16
}

func (objref_std *OBJREFStandard) From_data(data []byte, offset, size int) *OBJREFStandard{

	//NOTE: check the sizes with respect to the decdoing and the size of the data types

	//end := offset + size
	objref_std.objref.Read_OBJREF(data, offset)
	offset += OBJREF_SIZE

	objref_std.std_flags = binary.LittleEndian.Uint16(data[offset : offset+4])
	objref_std.Counter = binary.LittleEndian.Uint16(data[offset+4 : offset+8])
	objref_std.OXID = binary.LittleEndian.Uint16(data[offset+8 : offset+16])
	objref_std.OID = binary.LittleEndian.Uint16(data[offset+16 : offset+24])

	if objref_std.Counter == 0{
		panic(0)
	}

	offset += 24
	objref_std.IPID = tools.Bin_to_str(data, offset)
	offset += CLSID_SIZE

	// TODO: assisng the res_addr to the rest of the data

	return objref_std	
}


func (objref_std *OBJREFStandard) Get_data() []byte{
	buffer := make([]byte, 0)

	binary.LittleEndian.PutUint16(buffer, objref_std.std_flags)
	binary.LittleEndian.PutUint32(buffer, 0)

	binary.LittleEndian.PutUint16(buffer, objref_std.OXID)
	binary.LittleEndian.PutUint16(buffer, objref_std.OID)
	//binary.LittleEndian.PutUint64(data, DualStringArray)

	temp := objref_std.objref.Get_data()
	buffer = append(buffer, temp...)
	
	return buffer
}
