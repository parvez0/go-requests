package requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
)

// custom types
type Client struct {
	Client *http.Client
	Request *http.Request
	Headers http.Header
	BasePath string
}

type Response struct {
	Res *http.Response
}

// global client options
type GlobalOptions struct {
	Timeout time.Duration
	BasePath string
	Headers http.Header
}

// options object
type Options struct {
	Url string
	Method string
	Headers http.Header
	Body interface{}
	Query map[string]string
}

// initialize logger
var log = NewLogger()

// create a request client with global configurations
func NewClient(args ...interface{}) *Client {
	var client Client
	client.Client = &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       30 * time.Second,
	}

	// fill the http client partially based on the config provided by the user
	for _, arg := range args{
		switch arg.(type) {
		case http.Client:
			log.Debugf("provided http.Client object replacing with original")
			cli := arg.(http.Client)
			client.Client = &cli
		case *http.Client:
			log.Debugf("provided pointer to http.Client object replacing with original")
			client.Client = arg.(*http.Client)
		case http.Transport:
			log.Debugf("provided http.Transport object replacing with original")
			cli := arg.(http.Transport)
			client.Client.Transport = &cli
		case *http.Transport:
			log.Debugf("provided pointer to http.Transport object replacing with original")
			client.Client.Transport = arg.(*http.Transport)
		case http.CookieJar:
			log.Debugf("provided http.CookieJar object replacing with original")
			cli := arg.(http.CookieJar)
			client.Client.Jar = cli
		case GlobalOptions:
			log.Debugf("provided an global options setting it in global config")
			opts := arg.(GlobalOptions)
			client.Client.Timeout = opts.Timeout
			client.Headers = opts.Headers
			client.BasePath = opts.BasePath
		default:
			log.Infof("type of client does not match %s", reflect.TypeOf(arg))
		}
	}

	return &client
}

// response wrapper for returning bytes of data returned by response
func (wr *Response) GetBody() ([]byte, error) {
	defer wr.Res.Body.Close()
	return ioutil.ReadAll(wr.Res.Body)
}

// response wrapper for returning headers filed from response
func (wr *Response) GetHeaders() http.Header {
	return wr.Res.Header
}

// response wrapper for returning the status code
func (wr *Response) GetStatusCode() int {
	return wr.Res.StatusCode
}

// creates a new request object, this needs to be invoke for every call
// client will hold all the global info about the request
func (client *Client) NewRequest(options Options) error {
	var err error
	uri := UriBuilder(client.BasePath, options.Url, options.Query)
	body, err := RequestBodyBuilder(options.Body)
	if err != nil{
		return err
	}
	client.Request, err = http.NewRequest(options.Method, uri, body)
	if len(options.Headers) != 0{
		client.Request.Header = options.Headers
	}
	return err
}

// converts the body to json, if body is empty returns null or empty
// returns a buffer of stringified struct or request body
func RequestBodyBuilder(body interface{}) (*bytes.Buffer, error) {
	var reader *bytes.Buffer
	switch body.(type) {
	case string:
		log.Debug("converting the request body to bytes of string")
		reader = bytes.NewBuffer([]byte(body.(string)))
	default:
		log.Debug("trying to convert the request body to json")
		mr, err := json.Marshal(body)
		if err != nil{
			log.Debugf("failed json marshall - %v", err)
			return nil, err
		}
		log.Debugf("stringifies json request body - %v", mr)
		reader = bytes.NewBuffer(mr)
	}
	return reader, nil
}

// triggers the api call and returns error or response object
func (client *Client) Send() (*Response, error) {
	var wr Response
	// calling the api
	resp, err := client.Client.Do(client.Request)
	// pass the api response to current client wrapper
	wr.Res = resp
	return &wr, err
}

// builds the uri with all the query parameters and if basepath is provided attaches that as well
// if basepath and full url both provided, full url will take precedence
func UriBuilder(basePath string, url string, qp map[string]string) string {
	qString := ""
	for k, v := range qp {
		qString += k +"="+ v
	}
	if strings.Contains(url, "?"){
		url += qString
	}
	url += "?" + qString
	if strings.HasPrefix(url, "http"){
		url += qString
	} else if strings.HasSuffix(basePath, "/"){
		url = basePath + qString
	} else {
		url = basePath + "/" + qString
	}
	return url
}
