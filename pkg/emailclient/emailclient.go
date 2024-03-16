package emailclient

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"github.com/diazharizky/go-rest-bootstrap/config"
)

type client struct {
	host       string
	port       int32
	senderName string
	email      string
	password   string
	smtpAddr   string
}

func init() {
	config.Global.SetDefault("emailclient.host", "localhost")
	config.Global.SetDefault("emailclient.port", 1025)
	config.Global.SetDefault("emailclient.sender_name", "rest")
	config.Global.SetDefault("emailclient.email", "rest")
	config.Global.SetDefault("emailclient.password", "")
}

func New() (c client) {
	c.host = config.Global.GetString("emailclient.host")
	c.port = config.Global.GetInt32("emailclient.port")
	c.senderName = config.Global.GetString("emailclient.sender_name")
	c.email = config.Global.GetString("emailclient.email")
	c.password = config.Global.GetString("emailclient.password")
	c.smtpAddr = fmt.Sprintf("%s:%d", c.host, c.port)

	return
}

func (c client) generateBody(to []string, cc []string, subject, message string) string {
	body := "From: " + c.senderName + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	return body
}

func (c client) Send(to []string, cc []string, subject, message string) error {
	auth := smtp.PlainAuth("", c.email, c.password, c.host)
	body := c.generateBody(to, cc, subject, message)

	return smtp.SendMail(
		c.smtpAddr, auth, c.email, append(to, cc...), []byte(body),
	)
}

func (c client) SendNoAuth(to []string, cc []string, subject, message string) error {
	cl, err := smtp.Dial(c.smtpAddr)
	if cl != nil {
		defer cl.Quit()
	}
	if err != nil {
		return err
	}
	defer cl.Close()

	if err = cl.Mail(c.senderName); err != nil {
		return err
	}

	for _, recipient := range to {
		if err = cl.Rcpt(recipient); err != nil {
			log.Printf("invalid recipient: %s", recipient)
		}
	}

	w, err := cl.Data()
	if err != nil {
		return err
	}

	body := c.generateBody(to, cc, subject, message)
	if _, err = w.Write([]byte(body)); err != nil {
		return err
	}

	if err = w.Close(); err != nil {
		return err
	}

	return nil
}
