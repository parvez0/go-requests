package requests_test

import (
	"github.com/parvez0/go-requests/requests"
	"net/http"
	"testing"
	"time"
)

// global http client
var client *

// creating a http client with default settings
func TestCreateClient(t *testing.T) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	gOptions := requests.GlobalOptions{
		Timeout:  30 * time.Second,
		BasePath: "https://app.yellowmessenger.com",
		Headers:  headers,
	}
	client = requests.NewClient(gOptions)
	t.Logf("client created with global options - %+v", gOptions)
}

// initializing the request with custom options
func TestDoGetRequest(t *testing.T) {
	options := requests.Options{
		Url: "/api/login",
		Method: "GET",
	}
	err := client.NewRequest(options)
	if err != nil{
		t.Fatalf("failed to create initailize the request object - %v", err)
	}
	res, err := client.Send()
	if err != nil{
		t.Fatalf("failed to make the get request - %v", err)
	}
	statusCode := res.GetStatusCode()
	_, err = res.GetBody()
	if err != nil{
		t.Fatalf("failed to get body from response - %v", err)
	}
	t.Logf("success in making the api call with status code - %d", statusCode)
}
