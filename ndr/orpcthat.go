package ndr

import (
	"encoding/binary"
	"fmt"
)
type ORPCTHAT struct {
	Flags      uint
	Extensions uint
}

func (orpcthat *ORPCTHAT) From_data(data []byte, offset int) (int, error){

	if len(data)-offset < 8 {
		return offset, fmt.Errorf("insufficient data for unpacking")
	}
	offset = 0

	orpcthat.Flags = uint(binary.LittleEndian.Uint32(data[ offset : offset+4 ]))
	orpcthat.Extensions = uint(binary.LittleEndian.Uint32(data[ offset+4 : offset+8 ]))

	// TODO: support to the exenstions
	return offset + 8, nil
}