package email

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/leetcode-golang-classroom/golang-smtp-sample/internal/types"
	"github.com/leetcode-golang-classroom/golang-smtp-sample/internal/util"
)

type Worker struct {
	sender   types.Sender
	dataChan chan types.EmailTemplateRequestBody
	sync.RWMutex
}

func NewWorker(sender types.Sender, dataChan chan types.EmailTemplateRequestBody) *Worker {
	return &Worker{
		sender:   sender,
		dataChan: dataChan,
	}
}

func (worker *Worker) Run(ctx context.Context) error {
	worker.Lock()
	defer worker.Unlock()
	log.Println("worker start")
	for data := range worker.dataChan {
		// convert to into slice
		to := strings.Split(data.ToAddr, ",")
		tmpl, err := template.ParseFiles(fmt.Sprintf("./templates/%s.html", data.Template))
		if err != nil {
			fmt.Fprint(os.Stderr, fmt.Errorf("failed to parse template: %w", err))
			continue
		}
		// Render the template with the map data
		var rendered bytes.Buffer
		if err := tmpl.Execute(&rendered, data.Vars); err != nil {
			fmt.Fprint(os.Stderr, fmt.Errorf("failed to render template: %w", err))
			continue
		}

		err = worker.sender.SendHtmlEmail(to, data.Subject, rendered.String())
		if err != nil {
			fmt.Fprint(os.Stderr, fmt.Errorf("send email failed: %w", err))
			continue
		}
		log.Printf("send %s successfully", data.ToAddr)
	}
	<-ctx.Done()
	log.Println("worker end")
	util.CloseDataChannel(worker.dataChan)
	return nil
}
