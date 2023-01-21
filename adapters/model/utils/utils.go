package utils

import (
	"fmt"
	"strconv"
)

func ConvertIntToString(num int)(string,error)  {
	numStr := strconv.Itoa(num)
	if numStr == "" {
		return "", fmt.Errorf("Failed to convert integer to string")
	}
	return numStr, nil
}