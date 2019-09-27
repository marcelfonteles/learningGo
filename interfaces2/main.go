package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	FirstName string
	LastName  string
	Sex       string
	Age       int
}

type Lawer struct {
	Person
	hasOAB    bool
	OABnumber int
}

type Medic struct {
	Person
	CrmNumber int
	Field string
}

type Programmer struct {
	Person
	StackOverFlow               bool
	ProgrammingLanguages        []string
	Coffees                     []string
	MeaningOfLifeUiverseAndMore int
}

type People interface {
	read() string
	speak(say string) string
	fight(action string) string
}

func (l Lawer) read() string {
	return "Lawer needs a lot of reader"
}

func (l Lawer) speak(say string) string {
	// Lawer can change what you say
	say = "You said this: ..."
	return say
}
func (l Lawer) fight(action string) string {
	return "Lawer does not. Lawer run from fight."
}

func (l Lawer) sayMyName() string {
	return l.FirstName + " " + l.LastName
}

func main() {
	var person People
	lawer1 := Lawer{
		Person:    Person{
			FirstName: "Leo",
			LastName: "Fonteles",
		},
		hasOAB:    true,
		OABnumber: 101,
	}

	person = lawer1
	fmt.Printf("%T\n", person)
	fmt.Printf("%v\n", person)
	fullName := lawer1.sayMyName()
	fmt.Println(fullName)

	var lawer2 People = Lawer{
		Person:    Person{
			FirstName: "Leo",
			LastName: "Fonteles",
		},
		hasOAB:    true,
		OABnumber: 101,
	}

	fmt.Println(reflect.TypeOf(lawer2).PkgPath(), reflect.TypeOf(lawer2).Name())
	fmt.Println(reflect.TypeOf(lawer2).String())

	var i People
	fmt.Printf("--------------------------------\n")
	fmt.Printf("%T\n", i)
	fmt.Printf("%v\n", i)
	var lawer3 Lawer
	fmt.Printf("%T\n", lawer3)
	fmt.Printf("%v\n", lawer3)
	i = lawer3
	fmt.Printf("--------------------------------\n")
	fmt.Printf("%T\n", i)
	fmt.Printf("%v\n", i)


}
