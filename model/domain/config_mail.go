package domain

type ConfigMail struct {
	SmtpHost, Name, Email, Password string
	SmtpPort                        int
}