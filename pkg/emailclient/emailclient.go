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
	config.Global.SetDefault("emailclient.sender_name", "gorest")
	config.Global.SetDefault("emailclient.email", "gorest")
	config.Global.SetDefault("emailclient.password", "")
}

func New() (cli client) {
	cli.host = config.Global.GetString("emailclient.host")
	cli.port = config.Global.GetInt32("emailclient.port")
	cli.senderName = config.Global.GetString("emailclient.sender_name")
	cli.email = config.Global.GetString("emailclient.email")
	cli.password = config.Global.GetString("emailclient.password")
	cli.smtpAddr = fmt.Sprintf("%s:%d", cli.host, cli.port)

	return
}

func (cli client) generateBody(to []string, cc []string, subject, message string) string {
	body := "From: " + cli.senderName + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	return body
}

func (cli client) Send(to []string, cc []string, subject, message string) error {
	auth := smtp.PlainAuth("", cli.email, cli.password, cli.host)
	body := cli.generateBody(to, cc, subject, message)

	return smtp.SendMail(
		cli.smtpAddr, auth, cli.email, append(to, cc...), []byte(body),
	)
}

func (cli client) SendNoAuth(to []string, cc []string, subject, message string) error {
	cl, err := smtp.Dial(cli.smtpAddr)
	if cl != nil {
		defer cl.Quit()
	}

	if err != nil {
		return err
	}
	defer cl.Close()

	if err = cl.Mail(cli.senderName); err != nil {
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

	body := cli.generateBody(to, cc, subject, message)
	if _, err = w.Write([]byte(body)); err != nil {
		return err
	}

	if err = w.Close(); err != nil {
		return err
	}

	return nil
}
