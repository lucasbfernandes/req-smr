package services

import (
	"io"
	"io/ioutil"
	"net/http"
	"req-smr/internal/models"
)

func BuildRequestObject(nativeRequest *http.Request) (*models.Request, error) {
	parsedBody, err := parseRequestBody(nativeRequest.Body)
	if err != nil {
		return nil, err
	}

	request := &models.Request{
		Method: nativeRequest.Method,
		Url: nativeRequest.URL.String(),
		Headers: nativeRequest.Header,
		Body: parsedBody,
	}
	return request, nil
}

func parseRequestBody(requestBody io.ReadCloser) ([]byte, error) {
	body, err := ioutil.ReadAll(requestBody)
	if err != nil {
		return nil, err
	}
	return body, nil
}