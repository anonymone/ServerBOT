package modules

import (
	"log"
	"net/smtp"
	"strconv"
	"strings"
)

// EmailInfo is including the information of email initializetion
type EmailInfo struct {
	LocalInfo  *LocalInfo
	Receiver   []string
	Sender     string
	ServerHost string
	ServerPort int
	APICredit  string
	TaskSend   string
	TaskCheck  string
	TaskTry    string
	Message    string
}

// Run is the interface of cron doing tasks.
func (e EmailInfo) Run() {
	var err error
	// log.Println("Enter Sender")
	if e.LocalInfo.changed {
		err = e.SendEmail(e.LocalInfo.IPs)
		// log.Println(e.LocalInfo.IPs)
	}
	if err != nil {
		log.Println("Failed to send emails.")
	}
}

// SendEmail is a simple method for sending emails.
func (e *EmailInfo) SendEmail(ips []string) (err error) {
	e.LocalInfo.changed = !e.LocalInfo.changed
	for _, ip := range e.LocalInfo.IPs {
		e.Message += ip + "\n"
	}

	reciver := e.Receiver
	auth := smtp.PlainAuth(
		"",
		e.Sender,
		e.APICredit,
		e.ServerHost,
	)
	sendList := strings.Join(e.Receiver, ",")
	log.Println("To: " + sendList + "\r\nFrom: " + e.Sender + "\r\nSubject: IPChanged Reports\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n")
	err = smtp.SendMail(
		e.ServerHost+":"+strconv.Itoa(e.ServerPort),
		auth,
		e.Sender,
		reciver,
		[]byte("To: "+sendList+"\r\nFrom: "+e.Sender+"\r\nSubject: IPChanged Reports\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n"+e.Message),
	)
	if err != nil {
		log.Println(err)
	}
	log.Println("Email is sent.")
	// log.Println("To: " + sendList + "\r\nFrom: " + e.Sender + "\r\nSubject: IPChanged Reports\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n" + e.Message)
	return err
}
