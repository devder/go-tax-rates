package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64
	for _, str := range strings {
		floatPrice, err := strconv.ParseFloat((str), 64)

		if err != nil {
			return nil, errors.New("failed to convert string")
		}

		floats = append(floats, floatPrice)
	}

	return floats, nil
}
