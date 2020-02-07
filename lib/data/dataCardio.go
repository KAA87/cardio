package dataCardio

import (
	"bytes"
	"encoding/binary"
	"os"
)

// Byte ...
type Byte struct {
	Val int8
}

// GetData ...
func GetDataFromFile(filePath string) ([]float64, error) {
	dataValues := make([]float64,0)

	file, err := os.Open(filePath)
	if err != nil{
		return nil,err
	}
	defer file.Close()

	byte := Byte{}

	for{
		data,err := readNextBytes(file, 2)
		if err != nil {
			break
		}

		buffer := bytes.NewBuffer(data)

		err = binary.Read(buffer, binary.LittleEndian, &byte)
		if err != nil{
			return nil, err
		}

		dataValues = append(dataValues, float64(byte.Val))
	}

	return dataValues, nil
}

// readNextBytes ...
func readNextBytes(file *os.File, number int) ([]byte, error)  {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}