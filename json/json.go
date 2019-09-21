package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type ColorGroup struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Colors []string `json:"otherColors"`
}

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
	Sex  string `json:"sex"`
}

func main() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crinmson", "Red", "Ruby", "Maroon"},
	}

	jsonGroup := group.Marshal()
	fmt.Println(string(jsonGroup))

	var group2 ColorGroup
	group2 = group2.Unmarshal(jsonGroup)
	fmt.Println(group2)

	user1 := User{
		ID:   1,
		Name: "Marcel",
		Age:  24,
		Sex:  "Male",
	}
	user2 := User{
		ID:   2,
		Name: "Larissa",
		Age:  28,
		Sex:  "Female",
	}
	var xUser []User
	xUser = []User{
		user1,
		user2,
	}
	fmt.Println(xUser)
	b, err := json.Marshal(xUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

}

func (group ColorGroup) Marshal() []byte {
	b, err := json.Marshal(group)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func (group ColorGroup) Unmarshal(b []byte) ColorGroup {
	err := json.Unmarshal(b, &group)
	if err != nil {
		log.Fatal(err)
	}
	return group
}

