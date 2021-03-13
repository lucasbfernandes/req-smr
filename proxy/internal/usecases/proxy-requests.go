package usecases

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"req-smr/internal/services"
)

type Proxy struct{}

func (proxy *Proxy) ServeHTTP(responseWriter http.ResponseWriter, httpRequest *http.Request) {

	requestId := uuid.New().String()
	services.RequestChanMap[requestId] = make(chan bool)

	// TODO remove this when possible
	lastEntry, err := services.GetRequests()
	if err != nil {
		fmt.Printf("ERROR:GET_LAST_ENTRY %s\n", err)
		// Return error message for client
		http.Error(responseWriter, err.Error(), http.StatusBadGateway)
		return
	}
	fmt.Printf("STEP:LAST_ENTRY: %s\n", lastEntry)

	// Build request object
	fmt.Printf("STEP:INCOMING_REQUEST %s\n", httpRequest)
	fmt.Println("STEP:BUILD_REQUEST_OBJECT")
	request, err := services.BuildRequestObject(httpRequest, "123")
	if err != nil {
		fmt.Printf("ERROR:BUILD_REQUEST_OBJECT %s\n", err)
		// Return error message for client
		http.Error(responseWriter, err.Error(), http.StatusBadGateway)
		return
	}

	fmt.Println("STEP:PERSIST_LOG")
	// Persist log (RAFT)
	err = services.SetRequest(request)
	if err != nil {
		fmt.Printf("ERROR:PERSIST_LOG %s\n", err)
		// Return error message for client
		http.Error(responseWriter, err.Error(), http.StatusBadGateway)
		return
	}

	// TODO Maybe add infinite loop with maximum execution time?
	fmt.Printf("STEP:WAITING_REQUEST_CHANNEL\n")
	<-services.RequestChanMap[requestId]
	fmt.Printf("STEP:FINISHED_WAITING_REQUEST_CHANNEL\n")

	// Forward request
	fmt.Println("STEP:PROXY_HTTP_FORWARD_REQUEST")
	res, err := services.ForwardRequest(request)
	if err != nil {
		fmt.Printf("ERROR:FORWARD_REQUEST %s\n", err)
		http.Error(responseWriter, err.Error(), http.StatusBadGateway)
		return
	}

	fmt.Println("STEP:PROXY_HTTP_WRITE_RESPONSE")
	// Return response to client
	services.WriteResponse(responseWriter, res)

	return
}