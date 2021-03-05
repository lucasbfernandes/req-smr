package usecases

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"req-smr/internal/constants"
	"req-smr/internal/models"
	"req-smr/internal/services"
)

type Proxy struct{}

func (proxy *Proxy) ServeHTTP(responseWriter http.ResponseWriter, httpRequest *http.Request) {

	// Build request object
	request, err := services.BuildRequestObject(httpRequest)
	if err != nil {
		fmt.Printf("ERROR:BUILD_REQUEST_OBJECT %s\n", err)
		// Return error message for client
		http.Error(responseWriter, err.Error(), http.StatusBadGateway)
		return
	}

	// Persist log (RAFT)
	err = services.SetRequest(request)
	if err != nil {
		fmt.Printf("ERROR:PERSIST_LOG %s\n", err)
		// Return error message for client
		http.Error(responseWriter, err.Error(), http.StatusBadGateway)
		return
	}

	// Forward request
	res, err := proxy.forwardRequest(request)
	if err != nil {
		fmt.Printf("ERROR:FORWARD_REQUEST %s\n", err)
		http.Error(responseWriter, err.Error(), http.StatusBadGateway)
		return
	}

	proxy.writeResponse(responseWriter, res)

	return
}

func (proxy *Proxy) forwardRequest(request *models.Request) (*http.Response, error) {
	// Prepare the destination endpoint to forward the request to.
	proxyUrl := fmt.Sprintf("http://127.0.0.1:%d%s", constants.ServicePort, request.RequestURI)

	// Create an HTTP client and a proxy request based on the original request.
	httpClient := http.Client{}
	proxyReq, err := http.NewRequest(request.Method, proxyUrl, bytes.NewReader(request.Body))
	if err != nil {
		fmt.Printf("ERROR:CREATE_NEW_REQUEST_OBJECT %s\n", err)
		return nil, err
	}

	// Do request
	res, err := httpClient.Do(proxyReq)
	if err != nil {
		fmt.Printf("ERROR:HTTP_CLIENT_DO %s\n", err)
		return nil, err
	}

	return res, nil
}

func (proxy *Proxy) writeResponse(responseWriter http.ResponseWriter, res *http.Response) {
	for name, values := range res.Header {
		responseWriter.Header()[name] = values
	}
	responseWriter.Header().Set("Server", "req-smr")
	responseWriter.WriteHeader(res.StatusCode)
	io.Copy(responseWriter, res.Body)
	res.Body.Close()
}