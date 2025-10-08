package emailclient

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"

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

type templateInterface interface {
	RenderBody() (*bytes.Buffer, error)
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

func (cli client) Send(to []string, cc []string, subject string, tpl templateInterface) error {
	auth := smtp.PlainAuth("", cli.email, cli.password, cli.host)
	body, err := tpl.RenderBody()
	if err != nil {
		return err
	}

	return smtp.SendMail(
		cli.smtpAddr, auth, cli.email, append(to, cc...), body.Bytes(),
	)
}

func (cli client) SendNoAuth(to []string, cc []string, subject string, tpl templateInterface) error {
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

	body, err := tpl.RenderBody()
	if err != nil {
		return err
	}

	if _, err = w.Write(body.Bytes()); err != nil {
		return err
	}

	if err = w.Close(); err != nil {
		return err
	}

	return nil
}
