package main

import (
	"net/http"

	"github.com/charmbracelet/huh"
)

type request struct {
    url string
    method string
    body string
}

func initialRequest() request {
    return request{
    	url:    "",
    	method: "",
    	body:   "",
    }
}

func (r request) requiresBody() bool {
    return r.method == http.MethodPost || r.method == http.MethodPut
}

func runUI() request {
    req := initialRequest()

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("URL").Value(&req.url),
		),
		huh.NewGroup(
			huh.NewSelect[string]().Title("Method").Options(
				huh.NewOption("Get", http.MethodGet),
				huh.NewOption("Post", http.MethodPut),
				huh.NewOption("Put", http.MethodPost),
				huh.NewOption("Delete", http.MethodDelete),
			).Value(&req.method),
		),
	).WithTheme(huh.ThemeBase())

	form.Run()

	if req.requiresBody() {
		bodyForm := huh.NewForm(huh.NewGroup(
			huh.NewText().Title("Request Body").Value(&req.body),
		)).WithTheme(huh.ThemeBase())
        bodyForm.Run()
	}

    return req
}
