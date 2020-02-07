package main

import (
	dataCardio "cardio/lib/data"
	"fmt"
	"os"
)

func init() {

}

func main()  {
	os.Setenv("UMV", "2")
	os.Setenv("NUMLEADS", "12")

	filePath := "./data/3.dat"

	data,err := dataCardio.GetDataFromFile(filePath)
	if err != nil {
		fmt.Println("data not get: \n", err)
		os.Exit(1)
	}

	dataCardio.GetDataLeeds(data)

	fmt.Print(data)
}