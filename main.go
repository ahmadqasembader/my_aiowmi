package main

import (
		"aiowmi/ndr"
		"encoding/binary"
		"fmt"
		)
func main(){
	ndr.Hello()

		// Define a binary data format
		//format := binary.LittleEndian
		var value1 uint64 = 123456789
		var value2 uint64 = 987654321
	
		// Pack data into a byte slice
		data := make([]byte, 2*8) // Two uint64 values, each 8 bytes
		binary.PutUvarint(data, value1)
		binary.PutUvarint(data[8:], value2)
	

		// Uvarint returns a uint64 and the number of bytes read

		// Unpack data from the byte slice
		unpackedValue1, buf := binary.Uvarint(data)
		unpackedValue2, buff := binary.Uvarint(data[8:])
	
		fmt.Printf("Value 1: %d - %d \n", unpackedValue1, buf)
		fmt.Printf("Value 2: %d - %d\n", unpackedValue2, buff)
}