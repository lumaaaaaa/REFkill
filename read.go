package main

import (
	"encoding/binary"
	"os"
)

func readLineFromREF(file *os.File) string {
	line := ""
	array := make([]byte, 1)
	for array[0] != 10 {
		flag := readREFData(file, array)
		if !flag {
			line += "   "
			break
		}
		line += string(array[0])
	}
	return line[:len(line)-2]
}

func readREFData(file *os.File, array []byte) bool {
	err := binary.Read(file, binary.LittleEndian, &array)
	if err != nil {
		return false
	}

	return true
}
