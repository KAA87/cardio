package dataCardio

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func GetDataLeeds (data []float64) {
	UmV,_ := strconv.Atoi(os.Getenv("UMV"))

	coefV := (float64(UmV) / 65536) / 2

	for k, i := range data {
		data[k] = roundDown(i * coefV , 6)
	}

	fmt.Print(data)
}


func roundDown(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Ceil(digit)
	newVal = round / pow
	return newVal
}