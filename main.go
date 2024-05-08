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

	"github.com/charmbracelet/huh"
)

const GET = "GET"
const POST = "POST"
const PUT = "PUT"
const DELETE = "DELETE"

func main() {
	var url string
	var method string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("URL").Value(&url),
		),
		huh.NewGroup(
			huh.NewSelect[string]().Title("Method").Options(
				huh.NewOption("Get", GET),
				huh.NewOption("Post", POST),
				huh.NewOption("Put", PUT),
				huh.NewOption("Delete", DELETE),
			).Value(&method),
		),
	).WithTheme(huh.ThemeBase())

	form.Run()

    res := fetchRequest(url, method)

    fmt.Println(res.Status)

    body, err := io.ReadAll(res.Body)
    if err != nil {
        log.Fatalln(err)
    }

    var jsonResponse bytes.Buffer

    json.Indent(&jsonResponse, body, "", "  ")

    prettyJson := strings.ReplaceAll(jsonResponse.String(), "\\", "")

    fmt.Println(prettyJson)
}

func fetchRequest(url string, method string) *http.Response {
    var res *http.Response
    var err error

    client := http.Client{ Timeout: 10 * time.Second }

    switch method {
    case GET:
        res, err = client.Get(url)
    }

    if err != nil {
        panic(err)
    }

    return res
}
