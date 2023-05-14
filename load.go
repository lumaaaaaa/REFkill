package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/4kills/go-libdeflate/v2"
)

func parseREF(file *os.File) bool {
	// Read the unit serial number and ignore
	readLineFromREF(file)
	dataLength, success := getCodedDataLength(file)
	if !success {
		return false
	}

	array := make([]byte, dataLength)
	readREFData(file, array)
	dc, err := libdeflate.NewDecompressor()
	if err != nil {
		return false
	}

	totalDataLength, success := getCodedDataLength(file)
	if !success {
		return false
	}

	num3 := 12
	num4 := 0

	var signals []CanSignal
	for int64(num4) < int64(uint64(totalDataLength)) {
		dataLength, success = getCodedDataLength(file)
		if !success {
			return false
		}

		array = make([]byte, dataLength)
		success = readREFData(file, array)
		if !success {
			// If we couldn't read anymore data, we have successfully read the file
			return writeCSV(signals)
		}

		//decompressedArray := make([]byte, len(array))
		_, decompressedArray, err := dc.DecompressZlib(array, nil)
		if err != nil {
			panic(err)
		}
		decompressedString := string(decompressedArray)

		data := strings.Split(decompressedString, ",")
		if len(data) >= num3 {
			id, _ := strconv.Atoi(data[1])
			startBit, _ := strconv.Atoi(data[3])
			length, _ := strconv.Atoi(data[4])
			offset, _ := strconv.ParseFloat(data[5], 64)
			scale, _ := strconv.ParseFloat(data[6], 64)
			maximum, _ := strconv.ParseFloat(data[7], 64)
			minimum, _ := strconv.ParseFloat(data[8], 64)
			dataLengthCode, _ := strconv.Atoi(data[11])

			signal := CanSignal{
				Name:           data[0],
				CanID:          id,
				Units:          data[2],
				StartBit:       byte(startBit),
				Length:         byte(length),
				Offset:         offset,
				Scale:          scale,
				Maximum:        maximum,
				Minimum:        minimum,
				DataFormat:     data[9],
				ByteOrder:      data[10],
				DataLengthCode: byte(dataLengthCode),
			}

			signals = append(signals, signal)

			fmt.Printf("%s | %d | %s | %d | %d | %f | %f | %f | %f | %s | %s | %d\n",
				signal.Name,
				signal.CanID,
				signal.Units,
				signal.StartBit,
				signal.Length,
				signal.Offset,
				signal.Scale,
				signal.Maximum,
				signal.Minimum,
				signal.DataFormat,
				signal.ByteOrder,
				signal.DataLengthCode)
		}

		num4++
	}
	return writeCSV(signals)
}

func getCodedDataLength(file *os.File) (uint, bool) {
	array := make([]byte, 2)
	var codedDataLength uint

	flag := readREFData(file, array)
	if flag {
		codedDataLength = uint((int(array[0]) << 8) | int(array[1]))
		return codedDataLength, true
	} else {
		return 0, false
	}
}
