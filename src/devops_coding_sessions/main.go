package main

import (
	"devops_coding_sessions/adapters"
	"devops_coding_sessions/infrastructure"
)

var controller adapters.Controller
var gui infrastructure.GraphicalUserInterface

func init() {
	controller = GetController()
	gui = GetGraphicalUserInterface()
}

func main() {
	gui.HandleInputRequest(controller)
}
