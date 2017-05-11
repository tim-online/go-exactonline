package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"strings"

	"github.com/tim-online/go-exactonline/utils"
)

const (
	mediaType = "application/json"
	charset   = "utf-8"
)

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func New(http *http.Client) *Client {
	return &Client{
		http: http,
	}
}

type Client struct {
	// HTTP client used to communicate with the API.
	http *http.Client

	// Url pointing to base Exact Online API
	baseURL *url.URL

	// Debugging flag
	debug bool

	// User agent for client
	userAgent string

	// Optional function called after every successful request made to the DO APIs
	onRequestCompleted RequestCompletionCallback
}

func (c *Client) SetBaseURL(baseURL *url.URL) {
	// set base url for use in http client
	c.baseURL = baseURL
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Client) NewRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	u := c.GetEndpoint(path)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// optionally pass along context
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	req.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", mediaType, charset))
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", c.userAgent)
	return req, nil
}

func (c *Client) GetEndpoint(path string) *url.URL {
	basePath := strings.TrimSuffix(c.baseURL.Path, "/")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	u := *c.baseURL
	u.Path = basePath + path
	return &u
}

// Do sends an API request and returns the API response. The API response is XML decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, responseBody interface{}) (*http.Response, error) {
	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, err
	}

	// interface implements io.Writer: write Body to it
	// if w, ok := response.Envelope.(io.Writer); ok {
	// 	_, err := io.Copy(w, httpResp.Body)
	// 	return httpResp, err
	// }

	// $top=1
	// {
	// 	"d" : {
	// 		"results" : [
	// 			{}
	// 		]
	// 	}
	// }

	// $top=2
	// {
	// 	"d" : [
	// 		{}
	// 	]
	// }

	// no $top
	// {
	// 	"d" : {
	// 		"results": [
	// 			{}
	// 		]
	// 	}
	// }

	type Envelope struct {
		D utils.JsonTester `json:"d"`
	}

	envelope := &Envelope{}
	err = json.NewDecoder(httpResp.Body).Decode(envelope)
	if err != nil {
		return httpResp, err
	}

	// get bytes
	b := []byte(envelope.D.RawMessage)

	// check if interface has ".Results" field
	r := reflect.ValueOf(responseBody)
	configField := reflect.Indirect(r).FieldByName("Results")
	hasResults := configField.IsValid()

	// check if json is an object
	isArray := envelope.D.IsArray()

	// conversie doen
	if hasResults && isArray {
		b = append([]byte(`{"results":`), b...)
		b = append(b, []byte("}")...)

		err = json.Unmarshal(b, responseBody)
		return httpResp, err
	}

	err = json.Unmarshal(b, responseBody)
	return httpResp, err
}