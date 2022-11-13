package settings

import (
	"wenba/internal/global"
	email2 "wenba/internal/pkg/email"
)

type email struct {
}

func (*email) Init() {
	global.Email = email2.InitEmail(global.Settings.Email.ValidEmail,
		global.Settings.Email.SmtpHost, global.Settings.Email.SmtpEmail, global.Settings.Email.SmtpPass)
}
