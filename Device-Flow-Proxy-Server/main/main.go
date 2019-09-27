package main

import (
	deviceAuthHandler "github.com/ayeshajay/Device-Flow-Proxy-Server/Handler"
	"net/http"
)

func main() {
	http.HandleFunc("/device_authorization", deviceAuthHandler.DeviceAuthHandler)
	http.HandleFunc("/token", deviceAuthHandler.TokenHandler)
	http.ListenAndServe(":8000", nil)
}

