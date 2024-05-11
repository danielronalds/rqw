package rqwui

import (
	"net/http"

	"github.com/charmbracelet/huh"
	"github.com/danielronalds/rqw/rqwlib"
)

// Entry point for running the UI for getting the request the User wants to send
//
// Currently implementing using charmbracelets huh, but likely to change to bubble tea
func RunUI(req rqwlib.Request, send bool) (rqwlib.Request, bool) {
	if !req.ValidUrl() {
		urlForm := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Title("URL").Value(&req.Url),
			),
		).WithTheme(huh.ThemeBase())

		urlForm.Run()
	}

	if !req.ValidMethod() {
		methodForm := huh.NewForm(huh.NewGroup(
			huh.NewSelect[string]().Title("Method").Options(
				huh.NewOption("Get", http.MethodGet),
				huh.NewOption("Post", http.MethodPut),
				huh.NewOption("Put", http.MethodPost),
				huh.NewOption("Delete", http.MethodDelete),
			).Value(&req.Method),
		)).WithTheme(huh.ThemeBase())

		methodForm.Run()
	}

	if req.RequiresBody() {
		bodyForm := huh.NewForm(huh.NewGroup(
			huh.NewText().Title("Request Body").Value(&req.Body),
		)).WithTheme(huh.ThemeBase())
		bodyForm.Run()
	}

	if !send {
		send = true
		confirmationForm := huh.NewForm(
			huh.NewGroup(
				huh.NewNote().Title("URL").Description(req.Url),
				huh.NewNote().Title("Method").Description(req.Method),
				huh.NewNote().Title("Body").Description(req.Body),
				huh.NewConfirm().Title("Send Request?").Negative("Don't Send").Value(&send).Affirmative("Send"),
			),
		).WithTheme(huh.ThemeBase())

		confirmationForm.Run()
	}

	return req, send
}
