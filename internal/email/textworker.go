package email

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/leetcode-golang-classroom/golang-smtp-sample/internal/types"
	"github.com/leetcode-golang-classroom/golang-smtp-sample/internal/util"
)

type TextWorker struct {
	sender   types.Sender
	dataChan chan types.EmailRequestBody
	sync.RWMutex
}

func NewTextWorker(sender types.Sender, dataChan chan types.EmailRequestBody) *TextWorker {
	return &TextWorker{
		sender:   sender,
		dataChan: dataChan,
	}
}

func (textWorker *TextWorker) Run(ctx context.Context) error {
	textWorker.Lock()
	defer textWorker.Unlock()
	log.Println("text worker start")
	for data := range textWorker.dataChan {
		// convert to into slice
		to := strings.Split(data.ToAddr, ",")
		err := textWorker.sender.SendEmail(to, data.Subject, data.Body)
		if err != nil {
			fmt.Fprint(os.Stderr, fmt.Errorf("send email failed: %w", err))
			continue
		}
		log.Printf("send %v sucessfully \n", data.ToAddr)
	}
	<-ctx.Done()
	log.Println("text worker end")
	util.CloseTextDataChannel(textWorker.dataChan)
	return nil
}
