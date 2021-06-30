package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
)

func main() {
	// Create a webservice
	webservice := new(restful.WebService)
	// Create a route and attach it to a handler
	webservice.Route(webservice.GET("/ping").To(pingTime))
	// Add the service to application
	restful.Add(webservice)
	http.ListenAndServe(":8080", nil)
}

func pingTime(req *restful.Request, resp *restful.Response) {
	// Write to the response
	io.WriteString(resp, fmt.Sprintf("%s", time.Now()))
}
