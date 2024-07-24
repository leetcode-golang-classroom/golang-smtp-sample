package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/leetcode-golang-classroom/golang-smtp-sample/internal/types"
)

var Validdate = validator.New()

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s %v", msg, err)
	}
}

func CloseChannel(ch chan error) {
	if _, ok := <-ch; ok {
		close(ch)
	}
}

func CloseDataChannel(ch chan types.EmailTemplateRequestBody) {
	if _, ok := <-ch; ok {
		close(ch)
	}
}

func CloseTextDataChannel(ch chan types.EmailRequestBody) {
	if _, ok := <-ch; ok {
		close(ch)
	}
}
func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	errResp := WriteJSON(w, status, map[string]string{"error": err.Error()})
	if errResp != nil {
		log.Fatal(errResp)
	}
}
