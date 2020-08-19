package notification

import (
	"errors"
	"fmt"
	"net/smtp"
)

type credentials struct {
	username, password string
}

type Mail struct {
	smtpServer, loginname, password, fromAddress string
}

func (m *Mail) SendMail() {

	fmt.Println("[SEND MAIL] Writing weekly status mail...")

	msg := []byte(
		"To: recipient@example.net\r\n" +
			"Subject: Weekly Mail!\r\n" +
			"\r\n" +
			"Greetings!\r\n")

	toAddresses := []string{"test@gmx.de"}
	auth := loginAuth(m.loginname, m.password)
	err := smtp.SendMail(m.smtpServer+":587", auth, m.fromAddress, toAddresses, msg)

	if err != nil {
		fmt.Println("Does not work", err)
	}
	defer fmt.Println("[SEND MAIL] Pigeon is on its way!")
}

func loginAuth(username, password string) smtp.Auth {
	return &credentials{username, password}
}

func (a *credentials) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *credentials) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unkown from Server")
		}
	}
	return nil, nil
}
