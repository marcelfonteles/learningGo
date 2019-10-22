package main

import "fmt"

func init() {
	fmt.Println("Remebering methods")
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Sex       string
}

type Human interface {
	speak()
}

type HumanPointer interface {
	speakPointer()
}

func (p Person) speak() {
	fmt.Println("this is a Person speaking... blá blá blá...")
}

func (p *Person) speakPointer() {
	fmt.Println("this is a Person speaking... blá blá blá...")
}

func saySomething(human Human) {
	human.speak()
}

func saySomethingPointer(human HumanPointer) {
	human.speakPointer()
}

func main() {
	p1 := Person{
		FirstName: "Marcel",
		LastName:  "Vieira",
		Age:       24,
		Sex:       "Male",
	}

	saySomething(p1)
	saySomething(&p1)
	saySomethingPointer(&p1)

}
