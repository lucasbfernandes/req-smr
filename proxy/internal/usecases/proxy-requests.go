package usecases

import (
	"fmt"
	"net/http"
	"req-smr/internal/services"
)

type Proxy struct{}

func (proxy *Proxy) ServeHTTP(responseWriter http.ResponseWriter, httpRequest *http.Request) {

	//requestId := uuid.New().String()
	//services.RequestChanMap[requestId] = make(chan bool)

	// Build request object
	request, err := services.BuildRequestObject(httpRequest, "123")
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

	// Maybe add infinite loop with maximum execution time?
	// <-services.RequestChanMap[requestId]

	// Forward request
	res, err := services.ForwardRequest(request)
	if err != nil {
		fmt.Printf("ERROR:FORWARD_REQUEST %s\n", err)
		http.Error(responseWriter, err.Error(), http.StatusBadGateway)
		return
	}

	// Return response to client
	services.WriteResponse(responseWriter, res)

	return
}