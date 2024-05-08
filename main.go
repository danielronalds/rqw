package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

)

func main() {
    request := runUI()

	res := fetchRequest(request)

	fmt.Println(res.Status)

	body, err := io.ReadAll(res.Body)
	badlyHandleError(err)

	var jsonResponse bytes.Buffer

	json.Indent(&jsonResponse, body, "", "  ")

	prettyJson := strings.ReplaceAll(jsonResponse.String(), "\\", "")

	fmt.Println(prettyJson)
}

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
