package api

import (
	"fmt"
	"net/http"
	"req-smr/internal/constants"
	"req-smr/internal/usecases"
)

func StartAPI() {
	fmt.Printf("STEP:START_API_PORT: %s\n", constants.ProxyPort)
	http.ListenAndServe(constants.ProxyPort, &usecases.Proxy{})
}