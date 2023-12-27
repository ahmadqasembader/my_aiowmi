package tools

import (
	"math/rand"
	"encoding/binary"
	"fmt"
)

// Generates a 128-bit CID
func GenCID() []byte{
	// max := uint64(math.Pow(2, 32)) - 1
	return []byte{
		byte(rand.Uint32()),
		byte(rand.Uint32()),
		byte(rand.Uint32()),
		byte(rand.Uint32()),
	}
}

// Function to generate a fixed CID for testing
func FixedCID() []byte {
	return []byte{0x01, 0x02, 0x03, 0x04}
}


func Bin_to_str(data []byte, offset int) string {
	// Unpack the first part of the UUID (Little-endian byte order)
	uuid1 := binary.LittleEndian.Uint32(data[offset:])
	uuid2 := binary.LittleEndian.Uint16(data[offset+4:])
	uuid3 := binary.LittleEndian.Uint16(data[offset+6:])

	// Unpack the second part of the UUID (Big-endian byte order)
	uuid4 := binary.BigEndian.Uint16(data[offset+8:])
	uuid5 := binary.BigEndian.Uint16(data[offset+10:])
	uuid6 := binary.BigEndian.Uint32(data[offset+12:])

	// Format the UUID and return it as a string
	return fmt.Sprintf("%08X-%04X-%04X-%04X-%04X%08X", uuid1, uuid2, uuid3, uuid4, uuid5, uuid6)
}

//Returns a uuid without the last 4 bytes of the uuid
func UUID_custom(uuid []byte) []byte{
	return uuid[: len(uuid) - 4]
}