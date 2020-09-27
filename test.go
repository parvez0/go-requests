package main

import (
	"github.com/parvez0/go-requests/requests"
	"log"
	"net/http"
	"time"
)

func main() {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	gOptions := requests.GlobalOptions{
		Timeout:  30 * time.Second,
		BasePath: "http://localhost:5000",
		Headers:  headers,
	}
	client := requests.NewClient(gOptions)
	options := requests.Options{
		Url: "/test",
		Method: "POST",
	}
	req, err := client.NewRequest(options)
	if err != nil{
		log.Fatalf("failed to create initailize the request object - %v", err)
	}
	res, err := req.Send()
	if err != nil{
		log.Fatalf("failed to make the get request - %v", err)
	}
	statusCode := res.GetStatusCode()
	body, err := res.GetBody()
	if err != nil{
		log.Fatalf("failed to get body from response - %v", err)
	}
	log.Printf("success in making the api call with status code - %d - body - %+v", statusCode, string(body))
}
