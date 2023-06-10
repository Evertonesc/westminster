package main

import (
	"westminster/application"
)

func main() {
	appEngine := application.New()
	appEngine.StartHTTPServer()
}
