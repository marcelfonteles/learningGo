package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	First string `json:"f_name"`
	Last  string`json:"l_name"`
	Age   int`json:"age"`
	Sex   string`json:"sex"`
}

func main() {
	user1 := User{
		First: "Marcel",
		Last:  "Vieira",
		Age:   24,
		Sex:   "M",
	}
	var user2 User
	user2 = User{
		First: "Larissa",
		Last:  "Fonteles",
		Age:   29,
		Sex:   "F",
	}
	user3 := User{
		First: "Silvana",
		Last:  "Albuquerque",
		Age:   56,
		Sex:   "F",
	}
	xUsers := []User{user1, user2, user3}
	marshalUsers, err := json.Marshal(xUsers)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(marshalUsers)
	stringMarshalUsers := string(marshalUsers)
	fmt.Println(stringMarshalUsers)
}
