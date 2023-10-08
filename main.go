package main

import (
	"fmt"
	"os"
)

var (
	fileName string
)

func printHelp() {
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  REFkill [input file .REF]")

	os.Exit(0)
}

func main() {
	fmt.Println("(*) REFkill - v1.2 - Release Build")

	args := os.Args[1:]

	if len(args) == 0 {
		printHelp()
	}

	fileName = args[0]

	// Open encrypted .REF file
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read and verify header
	text := readLineFromREF(file)
	if text == "Racelogic Can Data File V1a" {
		success := parseREF(file)
		if success {
			fmt.Println("(!) Successfully decrypted", "'"+fileName+"', wrote data to", "'"+fileName[:len(fileName)-4]+".csv'")
		} else {
			fmt.Println("(x) Failed to decrypt", "'"+fileName+"'! Ensure the file data is correct.")
			os.Exit(1)
		}
	} else {
		fmt.Println("(x) File", "'"+fileName+"'", "does not appear to be a valid Racelogic Can Data File (V1a).")
	}
}
