package tools

import (
	"fmt"
	"net/smtp"
)

func SendEmail(email string, msg string) {
	from := ""
	password := ""
	to := []string{email}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	addr := smtpHost + ":" + smtpPort
	message := []byte(msg)
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(addr, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Mail Sent")
}

func GenerateCode() string {
	// charset := []any{'A', 'B', 'C', 'D'}
	// rand.Int(random.Numeric, big.NewInt(67))
	return ""
}
