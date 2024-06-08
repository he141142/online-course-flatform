package static

type WebServerFramework string

const (
	Mux WebServerFramework = "Mux"
	Gin WebServerFramework = "Gin"
	Chi WebServerFramework = "Chi"
)

type ConnProtocol string

const (
	ProtocolHttp  ConnProtocol = "http"
	ProtocolHttps ConnProtocol = "https"
	ProtocolRpc   ConnProtocol = "rpc"
)