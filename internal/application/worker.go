package application

import "github.com/leetcode-golang-classroom/golang-smtp-sample/internal/email"

func (app *App) SetupWorker() {
	app.dataWorker = email.NewWorker(app.sender, app.dataChan)
	app.textWorker = email.NewTextWorker(app.sender, app.textChan)
}
