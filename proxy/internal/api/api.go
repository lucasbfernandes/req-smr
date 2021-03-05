package api

import (
	"net/http"
	"req-smr/internal/constants"
	"req-smr/internal/usecases"
)

func StartAPI() {
	http.ListenAndServe(constants.ProxyPort, &usecases.Proxy{})
}