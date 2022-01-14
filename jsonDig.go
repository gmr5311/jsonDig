package jsondig

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func FindInJSON(original, searchKey string, expectedValue interface{}) bool {
	found := false
	var converted map[string]interface{}
	json.Unmarshal([]byte(original), &converted)

	for key, value := range converted {
		if key == searchKey && !found {
			fmt.Printf("Found key \"%s\"\n", key)
			if reflect.TypeOf(value).Kind() == reflect.Float64 && reflect.TypeOf(expectedValue).Kind() == reflect.Int {
				var err error
				temp := fmt.Sprintf("%d", expectedValue)
				expectedValue, err = strconv.ParseFloat(temp, 64)
				if err != nil {
					fmt.Printf("Failed to convert float: %s\n", err)
				}
			}
			fmt.Printf("Type of value: %T\n", value)
			fmt.Printf("Type of expected value: %T\n", expectedValue)
			if value == expectedValue {
				found = true
				break
			}
		} else if reflect.TypeOf(value) == reflect.TypeOf(converted) && !found {
			tempString, err := json.Marshal(value)
			if err != nil {
				fmt.Printf("The map failed to marshal: %s\n", err)
			}
			found = FindInJSON(string(tempString), searchKey, expectedValue)
		}
	}
	return found
}
