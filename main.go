package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func main()  {

	getData()
	//getData1()


}

type Header struct {
	Xz int16
}

func getData() {
	file, err := os.Open("3.dat")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	header := Header{}

	for{
		data := readNextBytes(file, 2)

		buffer := bytes.NewBuffer(data)

		err = binary.Read(buffer, binary.LittleEndian, &header)
		if err != nil {
			log.Fatal("binary.Read failed", err)
		}

		fmt.Print(header.Xz)
		fmt.Print(" ")
	}
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

//
//func getData1() {
//	dat, err := ioutil.ReadFile("2.dat")
//
//	if err != nil{
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	fmt.Print(string(dat))
//}