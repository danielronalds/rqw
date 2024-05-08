package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
    request, send := runUI()

    if !send {
        fmt.Println("Canceled request")
        os.Exit(0)
    }

	res := fetchRequest(request)

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
