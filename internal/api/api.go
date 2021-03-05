package api

import (
	"fmt"
	"net/http"
	"req-smr/internal/constants"
	"req-smr/internal/usecases"
)

func StartAPI() {
	http.ListenAndServe(fmt.Sprintf(":%d", constants.ProxyPort), &usecases.Proxy{})
}