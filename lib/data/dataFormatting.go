package dataCardio

func leadVoltage(data []float64) []float64 {
	for k, i := range data {
		data[k] = i * 2
	}

	return data
}