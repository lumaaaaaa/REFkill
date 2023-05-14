package main

import (
	"fmt"
	"os"
)

const (
	filename = "BMW-3 Series (F30 F31 F34) 2012-2020.REF"
)

func main() {
	fmt.Println("(*) REFkill - v1.1 - Debug Build")

	// Open encrypted .REF file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read and verify header
	text := readLineFromREF(file)
	if text == "Racelogic Can Data File V1a" {
		success := parseREF(file)
		if success {
			fmt.Println("(!) Successfully decrypted", "'"+filename+"', wrote data to", "'"+filename[:len(filename)-4]+".csv'")
		} else {
			fmt.Println("(x) Failed to decrypt", "'"+filename+"'! Ensure the file data is correct.")
			os.Exit(1)
		}
	} else {
		fmt.Println("(x) File", "'"+filename+"'", "does not appear to be a valid Racelogic Can Data File (V1a).")
	}
}
