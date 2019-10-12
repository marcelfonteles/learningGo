package main

import (
	"fmt"
	"sort"
)

type Dog struct {
	Breed string
	Height float64 // centimeter
	Weight float64 // kilogram
	LifeExpectation int // year
}

type xDog []Dog

func (xd xDog) Len() int           { return len(xd) }
func (xd xDog) Less(i, j int) bool { return xd[i].Height < xd[j].Height }
func (xd xDog) Swap(i, j int)      { xd[i], xd[j] = xd[j], xd[i] }



func main() {
	lucyMaria := Dog{
		Breed:           "Schnauzer",
		Height:          52.3,
		Weight:          7.2,
		LifeExpectation: 12,
	}
	kat := Dog{
		Breed:           "Hard foot",
		Height:          45.2,
		Weight:          6.3,
		LifeExpectation: 13,
	}
	kit := Dog{
		Breed:           "Hard foot",
		Height:          43.2,
		Weight:          12.1,
		LifeExpectation: 13,
	}
	dogs := xDog{lucyMaria, kit, kat}

	sort.Sort(dogs)
	for _, value := range dogs {
		fmt.Println("Breed:", value.Breed)
		fmt.Println("Height:", value.Height)
		fmt.Println("Weight:", value.Weight)
		fmt.Println("Life expectation:", value.LifeExpectation)
		fmt.Println("---------------------------------")
	}



}
