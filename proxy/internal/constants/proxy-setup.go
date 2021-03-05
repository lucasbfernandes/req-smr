package constants

import "os"

var ProxyPort = os.Getenv("PROXY_PORT")

var ServicePort = os.Getenv("SERVICE_PORT")
