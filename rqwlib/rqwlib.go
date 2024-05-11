package rqwlib

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

// A struct for representing a user created request
type Request struct {
	// The target URL of the request
	Url string
	// The Method the request should use. This can be trusted as the user doesn't directly input it
	Method string
	// The Body of the request. Only populated if the method requires a Body (PUT or POST)
	Body string
}

// Checks if the URL is valid
func (r Request) ValidUrl() bool {
	return len(r.Url) > 0
}

// Checks if the given method is valid.
//
// NOTE: Mutates the value of r.method to uppercase
func (r Request) ValidMethod() bool {
	r.Method = strings.ToUpper(r.Method)
	validMethods := map[string]bool{
		http.MethodGet: true, http.MethodPost: true, http.MethodPut: true, http.MethodDelete: true,
	}
	return validMethods[r.Method]
}

// Returns whether the request will require a body
func (r Request) RequiresBody() bool {
	return r.Method == http.MethodPost || r.Method == http.MethodPut
}

// Fetches the users specified request, returning a pointer to the http.Response object
func FetchRequest(req Request) (*http.Response, error) {
	client := http.Client{Timeout: 10 * time.Second}

	reqBody := bytes.NewBuffer([]byte(req.Body))

	httpReq, err := http.NewRequest(req.Method, req.Url, reqBody)
    if err != nil {
        return nil, err
    }

	res, err := client.Do(httpReq)
    if err != nil {
        return nil, err
    }

	return res, nil
}

// Parses the given responses body as JSON with nice indenting
func GetPrettyResponseBodyJson(res *http.Response) (string, error) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
        return "", err
    }

	var jsonResponse bytes.Buffer

	json.Indent(&jsonResponse, body, "", "    ")

	return strings.ReplaceAll(jsonResponse.String(), "\\", ""), nil // Removing escaping slashes
}
