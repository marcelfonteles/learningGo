package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"

)

func main() {
	password := "somePassword"
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Bcrypt password:", bs)

	userPassword := "somePassword"
	err = bcrypt.CompareHashAndPassword(bs, []byte(userPassword))
	if err != nil {
		fmt.Println("wrong password,", err)
	} else {
		fmt.Println("correct password. You can enter now!")
	}

}
