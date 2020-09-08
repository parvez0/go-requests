package requests_test

import (
	"github.com/parvez0/go-requests/requests"
	"net/http"
	"testing"
	"time"
)

// global http client
var client *requests.Client

// creating a http client with default settings
func TestCreateClient(t *testing.T) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	gOptions := requests.GlobalOptions{
		Timeout:  30 * time.Second,
		BasePath: "https://www.random.org/",
		Headers:  headers,
	}
	client = requests.NewClient(gOptions)
	t.Logf("client created with global options - %+v", gOptions)
}

// initializing the request with custom options
func TestGetRequest(t *testing.T) {
	options := requests.Options{
		Url: "/quick-pick/",
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

// testing get request with query parameters
func TestQueryGetRequest(t *testing.T) {
	options := requests.Options{
		Url: "/quick-pick/",
		Method: "GET",
		Query: map[string]string{"tickets": "2", "lottery": "5x69.1x26"},
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