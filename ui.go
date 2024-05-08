package main

import (
	"net/http"

	"github.com/charmbracelet/huh"
)

// A struct for representing a user created request
type request struct {
    // The target URL of the request
    url string
    // The method the request should use. This can be trusted as the user doesn't directly input it
    method string
    // The body of the request. Only populated if the method requires a body (PUT or POST)
    body string
}

// Returns whether the request will require a body
func (r request) requiresBody() bool {
    return r.method == http.MethodPost || r.method == http.MethodPut
}

// Entry point for running the UI for getting the request the User wants to send
//
// Currently implementing using charmbracelets huh, but likely to change to bubble tea
func runUI() request {
    req := request{}

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
