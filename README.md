# Go Requests  [![GoDoc](https://godoc.org/github.com/sirupsen/logrus?status.svg)](https://golang.org/pkg/net/http/)

Go Request is a simplified http client helper, which is completely compatible with net/http module of GO, 
you can find the original package [here](https://golang.org/pkg/net/http/).

## Table of Contents
 - [Installation](#installation)
 - [Example](#Example)
 - [Request Methods](#Methods)
    - [GlobalOptions](#GlobalOptions)
    - [NewClient](#NewClient)
    - [Options](#Options)
    - [NewRequest](#NewRequest)
    - [Send](#Send)
 - [Response Methods](#ResponseMethods)
    - [GetBody](#GetBody)
    
#### Installation
    
    go get -t github.com/parvez0/go-request/request

#### Example

Import the requests module into your file and create a new client.

```go
package main

import "github.com/parvez0/go-requests/requests"

var client = requests.NewClient()
```

Performing a GET request

```go
options := requests.Options{
	Url: "/users",
	Method: "GET",
	Query: map[string]string{"type": "free"},
}

// creating a new request
err := client.NewRequest(options)

if err != nil{
    // handle error
}

// performing a get request on 
res, err := client.Send()

if err != nil{
    // handle error
}
// returns a status code
statusCode := res.GetStatusCode()

// returns []bytes of body
body, err = res.GetBody()
if err != nil{
	// handle error
}
```

Performing a POST request by reusing the same client

```go
headers := http.Header{}
headers.Set("Content-Type", "application/json")
options := requests.Options{
           	Url: "/user",
           	Method: "POST",
           	Headers: headers,
           	Body: body // accepts type interface, and tries to convert it to json
           }
// creating a new request
err := client.NewRequest(options)

if err != nil{
    // handle error
}

// performing a get request on 
res, err := client.Send()

if err != nil{
    // handle error
}
// returns a status code
statusCode := res.GetStatusCode()

// returns []bytes of body
body, err = res.GetBody()
if err != nil{
	// handle error
}
```
#### Request Methods
- #####GlobalOptions
  You can use global options to set request settings globally, pass the global options object to the NewClient function
  and will be used by default, you can overwrite these settings at request level 
```go
    // global client options
    gOptions := GlobalOptions {
    	Timeout time.Duration
    	BasePath string
    	Headers http.Header
    }

    client := requests.NewClient(goptions)
```
- ##### NewClient
  You can create your own request client with custom settings as follows
```go
   //define your own transport object
    transport := http.Transport{
    		MaxIdleConns:           10,
    		MaxIdleConnsPerHost:    20,
    		MaxConnsPerHost:        40,
    	}
    
    cient := requests.NewClient(transport)
```          
 OR
```go
    // create a http.Client and pass it to requests wrapper
    client := http.Client{}
    wrapper := requests.NewClient(client)
``` 

