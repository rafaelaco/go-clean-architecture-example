package infrastructure

import (
	"devops_coding_sessions/adapters"

	"github.com/gin-gonic/gin"
)

type WebApiImpl struct{}

func (api WebApiImpl) HandleInputRequest(controller adapters.Controller) {
	request := gin.Default()

	request.POST("/input", func(context *gin.Context) {
		jsonRawData, _ := context.GetRawData()
		var response string = api.SendDataToController(jsonRawData, controller)

		context.JSON(200, gin.H{
			"message": response,
		})
	})

	request.Run(":8080")
}

func (api WebApiImpl) SendDataToController(data []byte, controller adapters.Controller) string {
	return controller.HandleInputRequest(data)
}
