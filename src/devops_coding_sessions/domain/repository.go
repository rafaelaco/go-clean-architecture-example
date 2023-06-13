package domain

type Repository interface {
	Connect() bool
	Close() bool
	SavePerson(person Person) bool
}
