package utils

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"log"
	"net/smtp"
)

var (
	addr    string
	user    string
	pass    string
	host    string
	name    string
	replyTo string
)

func InitSMTP() {
	em := viper.Sub("email")
	host = em.GetString("host")
	port := em.GetString("port")
	user = em.GetString("user")
	pass = em.GetString("pass")
	addr = host + ":" + port
	name = viper.GetString("name")
	replyTo = em.GetString("replyTo")
}

func SendMail(to, subject, body string) error {
	em := email.NewEmail()
	em.From = name + "<" + user + ">"
	em.To = []string{to}
	em.Subject = subject
	em.Text = []byte(body)
	em.ReplyTo = []string{replyTo}
	err := em.SendWithTLS(addr, smtp.PlainAuth("", user, pass, host), &tls.Config{
		ServerName: host,
	})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
