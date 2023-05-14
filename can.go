package main

type CanSignal struct {
	Name           string
	CanID          int
	Units          string
	StartBit       byte
	Length         byte
	Offset         float64
	Scale          float64
	Maximum        float64
	Minimum        float64
	DataFormat     string
	ByteOrder      string
	DataLengthCode byte
}
