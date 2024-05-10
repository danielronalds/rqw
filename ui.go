package main

import (
	"net/http"
	"strings"

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

// Checks if the URL is valid
func (r request) validUrl() bool {
	return len(r.url) > 0
}

// Checks if the given method is valid.
//
// NOTE: Mutates the value of r.method to uppercase
func (r request) validMethod() bool {
	r.method = strings.ToUpper(r.method)
	validMethods := map[string]bool{
		http.MethodGet: true, http.MethodPost: true, http.MethodPut: true, http.MethodDelete: true,
	}
	return validMethods[r.method]
}

// Returns whether the request will require a body
func (r request) requiresBody() bool {
	return r.method == http.MethodPost || r.method == http.MethodPut
}

// Entry point for running the UI for getting the request the User wants to send
//
// Currently implementing using charmbracelets huh, but likely to change to bubble tea
func runUI(req request, send bool) (request, bool) {
	if !req.validUrl() {
		urlForm := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Title("URL").Value(&req.url),
			),
		).WithTheme(huh.ThemeBase())

		urlForm.Run()
	}

	if !req.validMethod() {
		methodForm := huh.NewForm(huh.NewGroup(
			huh.NewSelect[string]().Title("Method").Options(
				huh.NewOption("Get", http.MethodGet),
				huh.NewOption("Post", http.MethodPut),
				huh.NewOption("Put", http.MethodPost),
				huh.NewOption("Delete", http.MethodDelete),
			).Value(&req.method),
		)).WithTheme(huh.ThemeBase())

		methodForm.Run()
	}

	if req.requiresBody() {
		bodyForm := huh.NewForm(huh.NewGroup(
			huh.NewText().Title("Request Body").Value(&req.body),
		)).WithTheme(huh.ThemeBase())
		bodyForm.Run()
	}

	if !send {
		send = true
		confirmationForm := huh.NewForm(
			huh.NewGroup(
				huh.NewNote().Title("URL").Description(req.url),
				huh.NewNote().Title("Method").Description(req.method),
				huh.NewNote().Title("Body").Description(req.body),
				huh.NewConfirm().Title("Send Request?").Negative("Don't Send").Value(&send).Affirmative("Send"),
			),
		).WithTheme(huh.ThemeBase())

		confirmationForm.Run()
	}

	return req, send
}
