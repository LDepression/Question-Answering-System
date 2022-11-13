package email

type Email struct {
	ValidEmail string
	SmtpHost   string
	SmtpEmail  string //
	SmtpPass   string
}

// InitEmail 初始化邮箱的相关信息
func InitEmail(ValidEmail, SmtpHost, SmtpEmail, SmtpPass string) *Email {
	return &Email{
		ValidEmail: ValidEmail,
		SmtpHost:   SmtpHost,
		SmtpEmail:  SmtpEmail,
		SmtpPass:   SmtpPass,
	}
}
