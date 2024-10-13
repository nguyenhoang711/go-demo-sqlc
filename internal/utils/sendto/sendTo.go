package sendto

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/hoangnguyen/demo-sqlc/global"
	"go.uber.org/zap"
)

type EmailAddress struct {
	Name string `json:"name"`
	Address string `json:"address"`
}

type Mail struct {
	From EmailAddress
	To []string
	Subject string
	Body string
}

const (
	SMTPHost= "smtp.gmail.com"
	SMTPPort= "587"
	SMTPUsername= "phuocnt.hutech@gmail.com"
	SMTPPassword= "soagxllpomzsxwif"
)

func BuildMessage(mail Mail) string {
    msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
    msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
    msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
    msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
    msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

    return msg
}


func SendTextEmailOtp(to []string, from string, otp string) error {
	contentMail := Mail{
		From: EmailAddress{Address: from, Name: "test"},
		To: to,
		Subject: "OTP Verification",
		Body: fmt.Sprintf("<a>OTP verification code is %s</a>, Please enter it to verify your account!", otp),
	}

	messageMail := BuildMessage(contentMail)

	//send OTP
	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)
	err := smtp.SendMail(SMTPHost + ":587", auth, from, to, []byte(messageMail))
	if err != nil {
		global.Logger.Error("Email send failed::", zap.Error(err))
		return err
	}
	return nil
}