package server

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"time"
)

type Response struct {
	Time string `json:"time"`
	Services []Service `json:"services"`
}
 
type Service struct {
	Name string `json:"name"`
	SubServices []SubService `json:"subs"`
}

type SubService struct {
	Name string `json:"name"`
	Status bool `json:"status"`
}

func getRandBool() bool {
    now := time.Now()
    return now.UnixNano()%2 == 0
}

func GetResponse() (Response, error) {
	jsonResponse, err := ioutil.ReadFile("server/response.json")
	if err != nil {
		fmt.Println(err)
	}

	var response Response

	err = json.Unmarshal([]byte(jsonResponse), &response)
	if err != nil {
		fmt.Println(err)
	}

	response.Services[0].SubServices[0].Status = getRandBool();
	response.Services[0].SubServices[1].Status = getRandBool();
	response.Services[0].SubServices[2].Status = getRandBool();
	response.Services[1].SubServices[0].Status = getRandBool();
	response.Services[1].SubServices[1].Status = getRandBool();
	response.Services[2].SubServices[0].Status = getRandBool();
	response.Services[2].SubServices[1].Status = getRandBool();
	
	response.Time = string(time.Now().Format("2 Jan 2006 15:04:05"));

	return response, nil
}