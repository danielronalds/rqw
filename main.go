package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	req := request{}
    var send bool

	flag.StringVar(&req.url, "url", "", "The url to send to")
	flag.StringVar(&req.method, "method", "", "The method to send with")
	flag.StringVar(&req.body, "data", "", "The body of the request")
	flag.BoolVar(&send, "y", false, "Whether to prompt for confirmation before sending")

	flag.Parse()

	req, send = runUI(req, send)

	if !send {
		fmt.Println("Canceled request")
		os.Exit(0)
	}

	res := fetchRequest(req)

	fmt.Println(res.Status)

	prettyJson := getPrettyResponseBodyJson(res)

	fmt.Println(prettyJson)
}

// Parses the given responses body as JSON with nice indenting
func getPrettyResponseBodyJson(res *http.Response) string {
	body, err := io.ReadAll(res.Body)
	badlyHandleError(err)

	var jsonResponse bytes.Buffer

	json.Indent(&jsonResponse, body, "", "    ")

	return strings.ReplaceAll(jsonResponse.String(), "\\", "") // Removing escaping slashes
}

// Fetches the users specified request, returning a pointer to the http.Response object
func fetchRequest(req request) *http.Response {
	client := http.Client{Timeout: 10 * time.Second}

	reqBody := bytes.NewBuffer([]byte(req.body))

	httpReq, err := http.NewRequest(req.method, req.url, reqBody)
	badlyHandleError(err)

	res, err := client.Do(httpReq)
	badlyHandleError(err)

	return res
}

// Does what it says on the tin, badly handles an error by panicking and printing it
func badlyHandleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
