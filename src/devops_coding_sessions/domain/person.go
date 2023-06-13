package domain

type Person struct {
	Name string
}

func (person Person) Validate() bool {
	if person.Name == "" || person.Name == "Batman" {
		return false
	}

	return true
}

func (person *Person) AddLastNameRule() {
	if person.Name == "Martin" {
		person.Name += " Fowler"
	}
}
