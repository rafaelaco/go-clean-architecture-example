package main

import (
	"devops_coding_sessions/adapters"
	"devops_coding_sessions/infrastructure"
)

func GetController() adapters.Controller {
	return adapters.ControllerImpl{Repository: infrastructure.FileRepositoryImpl{}}
}

func GetGraphicalUserInterface() infrastructure.GraphicalUserInterface {
	return infrastructure.WebApiImpl{}
}
