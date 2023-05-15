package main

import (
	"fmt"
	"os"
)

const (
	CsvHeader = "Name,Can-ID,Units,Start Byte,Start Bit,Length,Offset,Scale,Maximum,Minimum,Data Format,Byte Order,Data Length Code\n"
)

func writeCSV(signals []CanSignal) bool {
	if len(signals) == 0 {
		return false
	}

	f, err := os.Create(filename[:len(filename)-4] + ".csv")
	if err != nil {
		return false
	}

	_, err = f.WriteString(CsvHeader)
	if err != nil {
		return false
	}

	for _, signal := range signals {
		_, err = f.WriteString(fmt.Sprintf("%s,%d,%s,%d,%d,%d,%f,%f,%f,%f,%s,%s,%d\n",
			signal.Name,
			signal.CanID,
			signal.Units,
			signal.StartBit/8, // start byte
			signal.StartBit%8, // start bit
			signal.Length,
			signal.Offset,
			signal.Scale,
			signal.Maximum,
			signal.Minimum,
			signal.DataFormat,
			signal.ByteOrder,
			signal.DataLengthCode))
		if err != nil {
			return false
		}
	}

	if err = f.Close(); err != nil {
		return false
	}

	return true
}
