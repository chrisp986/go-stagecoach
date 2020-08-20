package notification

import (
	"crypto/tls"
	"log"
	"time"

	"github.com/chrisp986/go-stagecoach/pkg/config"
	mail "github.com/xhit/go-simple-mail/v2"
)

var (
	htmlBody = `<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Hello Gophers!</title>
	</head>
		<body>
			<p>This is the <b>Go gopher</b>.</p>
			<p>Image created by Renee French</p>
		</body>
	</html>`
)

func getSenderData() *config.MailConfig {
	mcfg := config.MailConfig{}
	senderData := mcfg.ReadMailConfig()

	return senderData
}

//SendMail takes sender mail config and takes data that needs to be sent vie smtp
func SendMail() bool {

	sender := getSenderData()
	client := mail.NewSMTPClient()

	//SMTP Client
	client.Host = sender.SmtpServer
	client.Port = 587
	client.Username = sender.Loginname
	client.Password = sender.Password
	client.Encryption = mail.EncryptionTLS
	client.ConnectTimeout = 10 * time.Second
	client.SendTimeout = 10 * time.Second

	//KeepAlive is not settted because by default is false
	client.TLSConfig = &tls.Config{InsecureSkipVerify: true, ServerName: sender.SmtpServer}

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

// 	mail.senderId = "theemail@example.com"
// 	mail.toIds = []string{"anotheremail@example.com"}
// 	mail.subject = "This is the email subject"
// 	mail.body = "body"

// 	messageBody := mail.BuildMessage()

// 	smtpServer := SmtpServer{host: "smtp.office365.com", port: "587"}

// 	auth := smtp.PlainAuth("", mail.senderId, `mypassword`, smtpServer.host)

// 	fmt.Println(auth)

// 	tlsconfig := &tls.Config{
// 		InsecureSkipVerify: true,
// 		ServerName:         smtpServer.host,
// 	}

// 	conn, err := tls.Dial("tcp", "smtp.office365.com:587", tlsconfig)

// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	client, err := smtp.NewClient(conn, smtpServer.host)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	if err = client.Auth(auth); err != nil {
// 		log.Panic(err)
// 	}

// 	if err = client.Mail(mail.senderId); err != nil {
// 		log.Panic(err)
// 	}
// 	for _, k := range mail.toIds {
// 		if err = client.Rcpt(k); err != nil {
// 			log.Panic(err)
// 		}
// 	}

// 	w, err := client.Data()
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	_, err = w.Write([]byte(messageBody))
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	err = w.Close()
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	client.Quit()

// 	log.Println("Mail sent successfully")
// }

// func (m *Mail) SendMail() {

// 	log.Println("[SEND MAIL] Writing weekly status mail...")

// 	msg := []byte(
// 		"To: recipient@example.net\r\n" +
// 			"Subject: Weekly Mail!\r\n" +
// 			"\r\n" +
// 			"Greetings!\r\n")

// 	toAddresses := []string{"test@gmx.de"}
// 	auth := loginAuth(m.Loginname, m.Password)
// 	err := smtp.SendMail(m.SMTPServer+":587", auth, "m.FromAddress", toAddresses, msg)

// 	if err != nil {
// 		log.Println("Does not work", err)
// 	}
// 	defer log.Println("[SEND MAIL] Pigeon is on its way!")
// }

// func loginAuth(username, password string) smtp.Auth {
// 	return &credentials{username, password}
// }

// func (a *credentials) Start(server *smtp.ServerInfo) (string, []byte, error) {
// 	return "LOGIN", []byte(a.username), nil
// }

// func (a *credentials) Next(fromServer []byte, more bool) ([]byte, error) {
// 	if more {
// 		switch string(fromServer) {
// 		case "Username:":
// 			return []byte(a.username), nil
// 		case "Password:":
// 			return []byte(a.password), nil
// 		default:
// 			return nil, errors.New("unkown from Server")
// 		}
// 	}
// 	return nil, nil
// }
