package types

type EmailRequestBody struct {
	ToAddr  string `json:"to_addr" validate:"required"`
	Subject string `json:"subject" validate:"required"`
	Body    string `json:"body" validate:"required"`
}

type EmailTemplateRequestBody struct {
	ToAddr   string            `json:"to_addr" validate:"required"`
	Subject  string            `json:"subject" validate:"required"`
	Template string            `json:"template" validate:"required"`
	Vars     map[string]string `json:"vars"`
}
