package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/danielronalds/rqw/rqwlib"
	"github.com/danielronalds/rqw/rqwui"
)

func main() {
	req := rqwlib.Request{}
    var send bool

	flag.StringVar(&req.Url, "url", "", "The url to send to")
	flag.StringVar(&req.Method, "method", "", "The method to send with")
	flag.StringVar(&req.Body, "data", "", "The body of the request")
	flag.BoolVar(&send, "y", false, "Whether to prompt for confirmation before sending")

	flag.Parse()

    req, send = rqwui.RunUI(req, send)

	if !send {
		fmt.Println("Canceled request")
		os.Exit(0)
	}

	res, err := req.FetchResponse()
    badlyHandleError(err)

	fmt.Println(res.Status)

	prettyJson, err := rqwlib.GetPrettyResBodyJson(res)
    badlyHandleError(err)

	fmt.Println(prettyJson)
}

// Does what it says on the tin, badly handles an error by panicking and printing it
func badlyHandleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
