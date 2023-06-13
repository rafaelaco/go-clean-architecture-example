package application

import "devops_coding_sessions/domain"

type UseCaseInteractorImpl struct {
	Repository domain.Repository
}

func (useCaseInteractor UseCaseInteractorImpl) HandleInputRequest(data domain.InputData) string {
	var person domain.Person = domain.Person(data)

	if !person.Validate() {
		panic("The name is invalid!")
	}
	person.AddLastNameRule()
	useCaseInteractor.Repository.SavePerson(person)

	return person.Name
}
