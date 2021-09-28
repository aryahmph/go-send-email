package service

import "go-send-email/model/domain"

type MailService interface {
	Send(mail domain.SendMail)
}
