package main

import (
	"fmt"
	"strconv"
)

func main() {
	bedrooms, err := stringToInt("null")
	fmt.Println(bedrooms)
	fmt.Println(err)

	if err == nil {
		filters_NumberOfBedrooms := []int32{int32(bedrooms)}
		fmt.Println(filters_NumberOfBedrooms)
	}
}

func stringToInt(str string) (int, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("failed to convert with error: %w", err)
	}

	return i, nil
}
