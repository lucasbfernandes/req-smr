package services

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"req-smr/internal/constants"
	"req-smr/internal/models"
)

func ForwardRequest(request *models.Request) (*http.Response, error) {
	// Prepare the destination endpoint to forward the request to.
	proxyUrl := fmt.Sprintf("http://127.0.0.1:%s%s", constants.ServicePort, request.RequestURI)

	// Create an HTTP client and a proxy request based on the original request.
	httpClient := http.Client{}
	proxyReq, err := http.NewRequest(request.Method, proxyUrl, bytes.NewBuffer(request.Body))
	if err != nil {
		fmt.Printf("ERROR:CREATE_NEW_REQUEST_OBJECT %s\n", err)
		return nil, err
	}

	proxyReq.Header = request.Headers

	// Do request
	res, err := httpClient.Do(proxyReq)
	if err != nil {
		fmt.Printf("ERROR:HTTP_CLIENT_DO %s\n", err)
		return nil, err
	}

	return res, nil
}

func WriteResponse(responseWriter http.ResponseWriter, res *http.Response) {
	for name, values := range res.Header {
		responseWriter.Header()[name] = values
	}
	responseWriter.Header().Set("Server", "req-smr")
	responseWriter.WriteHeader(res.StatusCode)
	io.Copy(responseWriter, res.Body)
	res.Body.Close()
}