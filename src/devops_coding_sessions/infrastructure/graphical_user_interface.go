package infrastructure

import (
	"devops_coding_sessions/adapters"
)

type GraphicalUserInterface interface {
	HandleInputRequest(controller adapters.Controller)
	SendDataToController(data []byte, controller adapters.Controller) string
}
