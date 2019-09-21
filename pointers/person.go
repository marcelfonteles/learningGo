package pointers


type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func ChangeMe(pointerPerson *Person) {
	pointerPerson.FirstName = "Marcel"
	pointerPerson.LastName  = "Vieira"
	pointerPerson.Age       = 24
}