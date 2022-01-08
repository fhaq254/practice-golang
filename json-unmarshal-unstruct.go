// https://tutorialedge.net/golang/go-json-tutorial/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`

	// Decode 'json' to 'map'
	var unstructData map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &unstructData)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", unstructData)
}
