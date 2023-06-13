package application

import "devops_coding_sessions/domain"

type UseCaseInteractor interface {
	HandleInputRequest(data domain.InputData) string
}
