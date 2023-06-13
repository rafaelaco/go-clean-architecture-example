package infrastructure

import (
	"devops_coding_sessions/domain"
	"fmt"
)

type FileRepositoryImpl struct{}

func (repository FileRepositoryImpl) Connect() bool {
	fmt.Println("Connecting...")
	return true
}

func (repository FileRepositoryImpl) Close() bool {
	fmt.Println("Closing connection...")
	return true
}

func (repository FileRepositoryImpl) SavePerson(person domain.Person) bool {
	fmt.Println(person.Name + " saved!")
	return true
}
