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
		BasePath: "https://app.yellowmessenger.com",
		Headers:  headers,
	}
	client := requests.NewClient(gOptions)
	options := requests.Options{
		Url: "/api/data/nlp/onlyMapping",
		Method: "GET",
		Query: map[string]string{"bot": "x1572327973956"},
	}
	err := client.NewRequest(options)
	if err != nil{
		log.Fatalf("failed to create initailize the request object - %v", err)
	}
	res, err := client.Send()
	if err != nil{
		log.Fatalf("failed to make the get request - %v", err)
	}
	statusCode := res.GetStatusCode()
	_, err = res.GetBody()
	if err != nil{
		log.Fatalf("failed to get body from response - %v", err)
	}
	log.Printf("success in making the api call with status code - %d", statusCode)
}
