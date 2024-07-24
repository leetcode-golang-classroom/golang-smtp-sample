package application

import (
	"net/http"

	"github.com/leetcode-golang-classroom/golang-smtp-sample/internal/email"
	"github.com/leetcode-golang-classroom/golang-smtp-sample/internal/types"
	"github.com/leetcode-golang-classroom/golang-smtp-sample/internal/util"
)

func (app *App) SetupRoute() {
	// default route
	app.router.HandleFunc("GET /", HealthHandler)
	// setup email router
	sender := email.NewSender(app.cfg)
	dataChan := make(chan types.EmailTemplateRequestBody, 10)
	app.dataChan = dataChan
	textChan := make(chan types.EmailRequestBody, 10)
	app.textChan = textChan
	app.sender = sender
	emailHdr := email.NewHandler(sender, dataChan, textChan)
	emailRoute := email.NewRouter(emailHdr, app.router)
	emailRoute.SetupRoute()
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	util.FailOnError(util.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "status OK",
	}), "failed to response")
}
