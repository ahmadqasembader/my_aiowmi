package ndr

import (
	"encoding/binary"
	"fmt"
)
type ORPCTHAT struct {
	flags      uint
	extensions uint
}

func (orpcthat *ORPCTHAT) from_data(data []byte, offset int) (int, error){
	offset = 0

	if len(data)-offset < 8 {
		return offset, fmt.Errorf("insufficient data for unpacking")
	}

	orpcthat.flags = uint(binary.LittleEndian.Uint32(data[ offset : offset+4 ]))
	orpcthat.extensions = uint(binary.LittleEndian.Uint32(data[ offset+4 : offset+8 ]))

	// TODO: support to the exenstions
	return offset + 8, nil
}