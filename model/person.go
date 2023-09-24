package model

type Person struct {
	Name    string `json:"name" xml:"name" form:"name"`
	Surname string `json:"surname" xml:"surname" form:"surname"`
}

func SetName(person *Person, name string) *Person {
	person.Name = name
	return person
}

func SetSurname(person *Person, surname string) *Person {
	person.Surname = surname
	return person
}

func GetName(person *Person) string {
	return person.Name
}

func GetSurname(person *Person) string {
	return person.Surname
}
