package email

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/leetcode-golang-classroom/golang-smtp-sample/internal/types"
	"github.com/leetcode-golang-classroom/golang-smtp-sample/internal/util"
)

type Handler struct {
	sender   *Sender
	dataChan chan types.EmailTemplateRequestBody
	textChan chan types.EmailRequestBody
}

func NewHandler(sender *Sender, dataChan chan types.EmailTemplateRequestBody, textChan chan types.EmailRequestBody) *Handler {
	return &Handler{
		sender:   sender,
		dataChan: dataChan,
		textChan: textChan,
	}
}
func (handler *Handler) HTMLTemplateEmailHandler(w http.ResponseWriter, r *http.Request) {
	var request types.EmailTemplateRequestBody
	// decode data into request
	if err := util.ParseJSON(r, &request); err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// validate input
	if err := util.Validdate.Struct(request); err != nil {
		var valErrs validator.ValidationErrors
		if errors.As(err, &valErrs) {
			util.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", valErrs))
			return
		}
	}
	log.Printf("send job to worker %v\n", request.ToAddr)
	handler.dataChan <- request
	util.WriteJSON(w, http.StatusCreated, map[string]string{
		"message": fmt.Sprintf("send email %v worker successful", request.ToAddr),
	})
}

func (handler *Handler) EmailHandler(w http.ResponseWriter, r *http.Request) {
	var request types.EmailRequestBody
	// decode data into request
	if err := util.ParseJSON(r, &request); err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// validate input
	if err := util.Validdate.Struct(request); err != nil {
		var valErrs validator.ValidationErrors
		if errors.As(err, &valErrs) {
			util.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", valErrs))
			return
		}
	}
	log.Printf("send job to text Worker %v\n", request.ToAddr)
	handler.textChan <- request
	util.WriteJSON(w, http.StatusCreated, map[string]string{
		"message": fmt.Sprintf("send email to %v successful", request.ToAddr),
	})
}
