package adapters

import (
	"devops_coding_sessions/application"
	"devops_coding_sessions/domain"
	"encoding/json"
)

type ControllerImpl struct {
	Repository domain.Repository
}

func (controller ControllerImpl) convertRawDataToBusinessData(rawData []byte) domain.InputData {
	var inputDataConverted domain.InputData
	json.Unmarshal([]byte(rawData), &inputDataConverted)

	return inputDataConverted
}

func (controller ControllerImpl) HandleInputRequest(rawData []byte) string {
	var inputData domain.InputData = controller.convertRawDataToBusinessData(rawData)
	var useCaseInteractor application.UseCaseInteractor = application.UseCaseInteractorImpl{Repository: controller.Repository}

	return useCaseInteractor.HandleInputRequest(inputData)
}
