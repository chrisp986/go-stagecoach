package notification

import (
	"crypto/tls"
	"log"
	"os"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

var (
	htmlBody = `<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>StageCoach!</title>
	</head>
		<body>
			<p>This is a mail sent by <a href="https://github.com/chrisp986/go-stagecoach">Stagecoach!</a></p>
			<p>Made by chrisp986</p>
		</body>
	</html>`
)

//SendMail takes sender mail config and takes data that needs to be sent vie smtp
func SendMail() bool {

	smtpServer := os.Getenv("toy_smtp")
	username := os.Getenv("toy_user")
	password := os.Getenv("toy_pw")

	client := mail.NewSMTPClient()

	//SMTP Client
	client.Host = smtpServer
	client.Port = 587
	client.Username = username
	client.Password = password
	client.Encryption = mail.EncryptionTLS
	client.ConnectTimeout = 10 * time.Second
	client.SendTimeout = 10 * time.Second

	//KeepAlive is not setted because by default is false
	client.TLSConfig = &tls.Config{InsecureSkipVerify: true, ServerName: smtpServer}

	//Connect to client
	smtpClient, err := client.Connect()

	if err != nil {
		log.Println("Expected nil, got", err, "connecting to client")
		return false
	}

	err = sendEmail(htmlBody, "bbb@gmail.com", smtpClient)
	if err != nil {
		log.Println("Expected nil, got", err, "sending email")
		return false
	}

	return true
}

func sendEmail(htmlBody string, to string, smtpClient *mail.SMTPClient) error {
	//Create the email message
	email := mail.NewMSG()

	email.SetFrom("From Example <from.email@example.com>").
		AddTo(to).
		SetSubject("New Go Email")

	//Get from each mail
	email.GetFrom()
	email.SetBody(mail.TextHTML, htmlBody)

	//Send with high priority
	// email.SetPriority(mail.PriorityLow)

	//Pass the client to the email message to send it
	err := email.Send(smtpClient)

	return err
}
