package wx

import (
	"encoding/json"
	"fmt"
	"strconv"
)


func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

/*
 struct convert json string
*/
func Struct2JsonString(structt interface{}) (jsonString string, err error) {
	data, err := json.Marshal(structt)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
