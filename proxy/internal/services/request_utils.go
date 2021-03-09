package services

import (
	"io"
	"io/ioutil"
	"net/http"
	"req-smr/internal/models"
)

var RequestChanMap = make(map[string]chan bool)

func BuildRequestObject(nativeRequest *http.Request, requestId string) (*models.Request, error) {
	parsedBody, err := parseRequestBody(nativeRequest.Body)
	if err != nil {
		return nil, err
	}

	request := &models.Request{
		Id: requestId,
		RequestURI: nativeRequest.RequestURI,
		Host: nativeRequest.Host,
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