package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // this is to hide the actual name and use a alias in place of that
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // this is to hide the password field entirely
	Tags     []string `json:"tags,omitempty"` // this is to hide the tags field if it is empty
}

func main() {
	fmt.Println("Welcome to json in golang")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS", 299, "LearnCodeOnline.in", "abc@123", []string{"web-dev", "js"}},
		{"MERN", 199, "LearnCodeOnline.in", "abc@123", []string{"web-dev", "js"}},
		{"Angular", 199, "LearnCodeOnline.in", "abc@123", []string{"web-dev", "js"}},
		{"Vue", 199, "LearnCodeOnline.in", "abc@123", []string{"web-dev", "js"}},
		{"Django", 299, "LearnCodeOnline.in", "abc@123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`[
		{
			"coursename": "ReactJS",
			"Price": 299,
			"website": "LearnCodeOnline.in",
			"tags": [
				"web-dev",
				"js"
			]
		}
	]`)

	// 1) Validate the JSON
	checkValid := json.Valid(jsonDataFromWeb)
	if !checkValid {
		fmt.Println("JSON was not valid")
		return
	}
	fmt.Println("JSON was valid")

	// 2) Decode into a slice of `course` structs
	var lcoCourses []course
	err := json.Unmarshal(jsonDataFromWeb, &lcoCourses)
	if err != nil {
		fmt.Println("Error unmarshalling into lcoCourses:", err)
		return
	}
	fmt.Printf("Decoded into []course:\n%#v\n\n", lcoCourses)

	// 3) Decode into a generic slice of maps
	var myOnlineData []map[string]interface{}
	err = json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	if err != nil {
		fmt.Println("Error unmarshalling into myOnlineData:", err)
		return
	}
	fmt.Printf("Decoded into []map[string]interface{}:\n%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("For key %v, value is %v\n", k, v)
	}
}

