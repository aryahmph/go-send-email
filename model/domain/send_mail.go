package domain

type SendMail struct {
	To               []string
	Subject, Message string
}
