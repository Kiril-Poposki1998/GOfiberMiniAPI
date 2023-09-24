package model

type Person struct {
	name    string
	surname string
}

func SetName(person *Person, name string) *Person {
	person.name = name
	return person
}

func SetSurname(person *Person, surname string) *Person {
	person.surname = surname
	return person
}

func GetName(person *Person) string {
	return person.name
}

func GetSurname(person *Person) string {
	return person.surname
}
