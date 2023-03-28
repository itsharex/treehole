package utils

import (
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"log"
	"net/smtp"
)

var (
	addr string
	user string
	pass string
	host string
	name string
)

func InitSMTP() {
	em := viper.Sub("email")
	host = em.GetString("host")
	port := em.GetString("port")
	user = em.GetString("user")
	pass = em.GetString("pass")
	addr = host + ":" + port
	name = viper.GetString("name")
}

func SendMail(to, subject, body string) error {
	em := email.NewEmail()
	em.From = name + "<" + user + ">"
	em.To = []string{to}
	em.Subject = subject
	em.Text = []byte(body)
	err := em.Send(addr, smtp.PlainAuth("", user, pass, host))
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
