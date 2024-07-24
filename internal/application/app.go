package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/leetcode-golang-classroom/golang-smtp-sample/internal/config"
	"github.com/leetcode-golang-classroom/golang-smtp-sample/internal/types"
	"github.com/leetcode-golang-classroom/golang-smtp-sample/internal/util"
)

type App struct {
	cfg        *config.Config
	router     *http.ServeMux
	dataWorker types.Worker
	dataChan   chan types.EmailTemplateRequestBody
	sender     types.Sender
	textChan   chan types.EmailRequestBody
	textWorker types.Worker
}

func New(cfg *config.Config) *App {
	app := &App{
		cfg:    cfg,
		router: http.NewServeMux(),
	}
	app.SetupRoute()
	app.SetupWorker()
	return app
}

func (app *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.cfg.Port),
		Handler: app.router,
	}
	log.Printf("Starting server on %d\n", app.cfg.Port)
	errCh := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			errCh <- fmt.Errorf("failed to start server: %w", err)
		}
		util.CloseChannel(errCh)
	}()
	go func() {
		err := app.dataWorker.Run(ctx)
		if err != nil {
			errCh <- fmt.Errorf("failed to start worker: %w", err)
		}
		util.CloseChannel(errCh)
	}()
	go func() {
		err := app.textWorker.Run(ctx)
		if err != nil {
			errCh <- fmt.Errorf("failed to start text worker: %w", err)
		}
		util.CloseChannel(errCh)
	}()
	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		log.Printf("app stop")
		timeout, cancel := context.WithTimeout(ctx, time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
