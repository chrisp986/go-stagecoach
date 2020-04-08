package notification

import (
	"errors"
	"fmt"
	"net/smtp"
)

type loginAuth struct {
	username, password string
}

type Mail struct {
	smtpServer, loginname, password, fromAddress string
}


func SendMail() {

	fmt.Println("[SEND MAIL] Writing weekly status mail...")

	msg := []byte(
		"To: recipient@example.net\r\n" +
			"Subject: Weekly Mail!\r\n" +
			"\r\n" +
			"Greetings!\r\n")

	toAddresses := []string{"d4m1en@gmx.de"}
	auth := LoginAuth(loginname, password)
	err := smtp.SendMail(smtpServer+":587", auth, fromAddress, toAddresses, msg)

	if err != nil {
		fmt.Println("Does not work", err)
	}
	defer fmt.Println("[SEND MAIL] Pigeon is on its way!")
}


func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}


func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}


func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
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
