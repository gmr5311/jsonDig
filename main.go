package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

const silly = `{"id":"1","name":"bob","deep":{"nestedID":7,"nestedName":"joe", "deeper":{"super":7,"something":"cool string"}}}`

func main() {
	fmt.Printf("result of find: %t\n", findInJSON(silly, "super", 7))
}

func findInJSON(original, searchKey string, expectedValue interface{}) bool {
	found := false
	var converted map[string]interface{}
	json.Unmarshal([]byte(original), &converted)

	for key, value := range converted {
		if key == searchKey && !found {
			if reflect.TypeOf(value).Kind() == reflect.Float64 && reflect.TypeOf(expectedValue).Kind() == reflect.Int {
				var err error
				temp := fmt.Sprintf("%d", expectedValue)
				expectedValue, err = strconv.ParseFloat(temp, 64)
				if err != nil {
					fmt.Printf("Failed to convert float: %s\n", err)
				}
			}
			if value == expectedValue {
				found = true
				break
			}
		} else if reflect.TypeOf(value) == reflect.TypeOf(converted) && !found {
			tempString, err := json.Marshal(value)
			if err != nil {
				fmt.Printf("The map failed to marshal: %s\n", err)
			}
			found = findInJSON(string(tempString), searchKey, expectedValue)
		}
	}
	return found
}
