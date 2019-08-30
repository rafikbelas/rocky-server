package server

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Response struct {
	Services []Service `json:"services"`
}
 
type Service struct {
	Name string `json:"name"`
	Status bool `json:"status"`
}

func GetReports() {
	jsonResponse, err := ioutil.ReadFile("server/response.json")
	if err != nil {
		fmt.Println(err)
	}	

	fmt.Println(string(jsonResponse))

	var service Service

	err = json.Unmarshal([]byte(jsonResponse), &service)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", service)

}