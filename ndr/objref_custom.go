package ndr

import (
	"aiowmi/tools"
	"encoding/binary"
)

type OBJREFCustom struct{
	objref		OBJREF
	cb_extension   uint16
	obj_ref_size   int
	obj_data    []byte
}


const (
	IID_IActivationPropertiesIn  = "000001A2-0000-0000-C000-000000000046"
	CLSID_ActivationPropertiesIn = "00000338-0000-0000-c000-000000000046"
)

func (obj_custom *OBJREFCustom) init() *OBJREFCustom{
	obj_custom.objref.Signature = 0x574F454D  // = 1464812877
	obj_custom.objref.Flags = FLAGS_OBJREF_CUSTOM
	obj_custom.objref.IID = IID_IActivationPropertiesIn // TODO: cut the last 4 bytes of the IID 
														// And convert all CLSID and IID to bin
	obj_custom.objref.CLSID = CLSID_ActivationPropertiesIn


	obj_custom.cb_extension = 0
	obj_custom.obj_ref_size = 0
	obj_custom.obj_data = nil

	return obj_custom
}

func (objref_custom *OBJREFCustom) From_data(data []byte, offset, size int) *OBJREFCustom{
	//TODO: size should be platform independent. The size of two short integers 
	CUSTOM_SIZE := 8
	end := offset + size

	objref_custom.objref.Read_OBJREF(data, offset)
	offset += OBJREF_SIZE

	objref_custom.objref.CLSID = tools.Bin_to_str(data, offset)
	offset += CLSID_SIZE
	
	//TODO: unpack the cb_extension and the objreft reference from the incoming data

	offset += CUSTOM_SIZE

	objref_custom.obj_data = data[offset : end]

	return objref_custom
}

func (objref_custom *OBJREFCustom) Set_object(blob []byte) {
	objref_custom.obj_data = blob
	objref_custom.obj_ref_size = len(blob)
} 

func (objref_custom *OBJREFCustom) Get_data() []byte{
	data   := make([]byte, 0)

	temp := objref_custom.objref.Get_data()
	offset := len(temp)
	data = append(data[: offset], temp...)
	binary.LittleEndian.PutUint16(data[offset :offset + 4], objref_custom.cb_extension)
	binary.LittleEndian.PutUint16(data[offset + 4 : offset + 8], uint16(objref_custom.obj_ref_size))

	data = append(data[offset + 8 : ], objref_custom.obj_data...)

	return data
}