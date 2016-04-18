package go_email
import (
	"github.com/scorredoira/email"
	"net/smtp"
"log"
)

func SendMail(header,from,password,smtpUri,port,fromName,title,body,attachmentFilePath string, tos[] string) error {
	header += "SendMail: "

	m := email.NewMessage(title, body)
	m.From = fromName
	m.To = tos
	if attachmentFilePath != "" {
		err := m.Attach(attachmentFilePath)
		log.Println(header, "1: ", err.Error())
		if err != nil {
			return err
		}
	}

	auth := smtp.PlainAuth(
		"",
		from,
		password,
		smtpUri,
	)

	err := email.Send(smtpUri+":"+port, auth, m)
	if err != nil {
		log.Println(header, "2: ", err.Error())
		return  err
	}

	return nil
}