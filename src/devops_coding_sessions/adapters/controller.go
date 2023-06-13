package adapters

import "devops_coding_sessions/domain"

type Controller interface {
	convertRawDataToBusinessData(rawData []byte) domain.InputData
	HandleInputRequest(rawData []byte) string
}
