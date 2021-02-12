package services

import (
	"encoding/json"
	"req-smr/internal/models"
)

func RequestToByteArray(request *models.Request) ([]byte, error) {
	serializedRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	return serializedRequest, nil
}

func ByteArrayToRequest(serializedRequest []byte) (*models.Request, error) {
	var request *models.Request
	err := json.Unmarshal(serializedRequest, request)
	if err != nil {
		return nil, err
	}
	return request, nil
}