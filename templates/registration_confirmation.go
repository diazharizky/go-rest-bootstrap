package templates

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

type RegistrationConfirmationTemplate struct {
	Name  string
	Email string
}

func (tmp RegistrationConfirmationTemplate) RenderBody() (*bytes.Buffer, error) {
	t, err := template.ParseFiles("./registration_confirmation.html")
	if err != nil {
		log.Fatalf("failed to parse template: %v", err)
		return nil, err
	}

	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write(fmt.Appendf(nil, "Subject: Welcome to Our Service!\n%s\n", mimeHeaders))

	if err := t.Execute(&body, tmp); err != nil {
		log.Fatalf("failed to execute template: %v", err)
		return nil, err
	}

	return &body, nil
}
