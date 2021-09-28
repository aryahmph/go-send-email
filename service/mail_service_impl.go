package service

import (
	"fmt"
	"go-send-email/helper"
	"go-send-email/model/domain"
	"net/smtp"
	"strings"
)

type MailServiceImpl struct {
	ConfigMail domain.ConfigMail
}

func NewMailServiceImpl(configMail domain.ConfigMail) *MailServiceImpl {
	return &MailServiceImpl{ConfigMail: configMail}
}

func (service *MailServiceImpl) Send(sendMail domain.SendMail) {
	body := "From: " + service.ConfigMail.Name + "\n" +
		"To: " + strings.Join(sendMail.To, ",") + "\n" +
		"Subject: " + sendMail.Subject + "\n\n" + sendMail.Message

	auth := smtp.PlainAuth("", service.ConfigMail.Email, service.ConfigMail.Password, service.ConfigMail.SmtpHost)
	smtpAddr := fmt.Sprintf("%s:%d", service.ConfigMail.SmtpHost, service.ConfigMail.SmtpPort)

	err := smtp.SendMail(smtpAddr, auth, service.ConfigMail.Email, sendMail.To, []byte(body))
	helper.PanicIfError(err)
}
