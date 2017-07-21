package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Age      int     `json:"age"`
}

type Sensor struct {
	SensorNumber    string
	SensorType      string
	EquipmentNumber string
	EquipmentType   string
	Time            int64
	Temperature     []float32
	Humidity        []float32
	GPSLongitude    float32
	GPSLatitude     float32
	Address         string
}

func main() {
	jsonStr := `{"username": "liuruichao", "password": "heheda", "age": 20}`
	var user User
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("username: %s, password: %s, age: %d.\n", user.Username, user.Password, user.Age)

	body, err := json.Marshal(user)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(body))

	jsonStr = `{     "SensorNumber": "cgq20170330001",     "SensorType": "HD-3K1",     "EquipmentNumber": "sb20170330000",     "EquipmentType": "what",     "Time": 1490155871000,     "Temperature": [         12.2,         12.3     ],     "Humidity": [         20.3,         20.4     ],     "GPSLongitude": 113.653056,     "GPSLatitude": 34.860076,     "Address": "what" }`
	var sensor Sensor
	err = json.Unmarshal([]byte(jsonStr), &sensor)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(sensor)
}
