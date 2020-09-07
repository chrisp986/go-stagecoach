package notification

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/smtp"
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

	smtpServer := os.Getenv("mailtrap_smtp")
	username := os.Getenv("mailtrap_user")
	password := os.Getenv("mailtrap_pw")

	if smtpServer == "" || username == "" || password == "" {
		log.Println("Variable in SendMail() is empty")
		return false
	}

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

// NEW TESTING

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")
		}
	}
	return nil, nil
}

func SendMail2() bool {
	smtpServer := os.Getenv("office_smtp")
	username := os.Getenv("toy_user")
	password := os.Getenv("toy_pw")
	receiver := []string{"cpeters986@gmail.com"}

	if smtpServer == "" || username == "" || password == "" {
		log.Println("Variable in SendMail() is empty")
		return false
	}

	auth := LoginAuth(username, password)

	// client, err := smtp.Dial(smtpServer)
	// client.Auth(LoginAuth(username, password))

	err := smtp.SendMail(smtpServer+":25", auth, username, receiver, []byte("test"))

	if err != nil {
		fmt.Println("SendMail2 err", err)
		return false
	}
	return true
}
