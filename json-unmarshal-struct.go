// https://tutorialedge.net/golang/go-json-tutorial/

package main

import (
	"encoding/json"
	"fmt"
)

type Sensor struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Time     string `json:"time"`
}

func main() {
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`

	// Decode 'json' to 'struct'
	var structData Sensor
	err := json.Unmarshal([]byte(jsonString), &structData)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", structData)
}
