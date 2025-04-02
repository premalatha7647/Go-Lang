package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFlotToFile(value float64, fileName string) {
	valueTxt := fmt.Sprint(value)
	err := os.WriteFile(fileName, []byte(valueTxt), 0644)
	fmt.Println(err)
}

func ReadFloatValue(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to find the file")
	}
	valueTxt := string(data)
	value, err := strconv.ParseFloat(valueTxt, 64)
	if err != nil {
		return 1000, errors.New("failed to parse the value")
	}
	return value, nil
}
