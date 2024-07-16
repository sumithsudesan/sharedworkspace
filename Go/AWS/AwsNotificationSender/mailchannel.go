package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"strconv"

	gomail "gopkg.in/mail.v2"
)

//MailChannel -
type MailChannel struct {
	config *Configuration
}

// NewMailChannel - create an instance of MailChannel object
func NewMailChannel(mailConfig *Configuration) IChannel {
	mailSender := MailChannel{
		config: mailConfig,
	}
	return &mailSender
}

// Send - sent mail to specified id
func (s *MailChannel) Send(recipientInfo RecipientInfo) error {

	fmt.Printf("[INFO] Sending mail, recipient:  %v\n", recipientInfo)

	if len(s.config.MailSrv()) == 0 {
		fmt.Printf("[ERROR] SMTP mail server param is empty.")
		return errors.New("Failed to send mail- SMTP mail server param is empty")
	}

	if len(s.config.MailSrvPort()) == 0 {
		fmt.Printf("[ERROR] SMTP mail server port param is empty.")
		return errors.New("Failed to send mail- SMTP mail server param  port is empty")
	}

	if len(s.config.SenderEmail()) == 0 {
		fmt.Printf("[ERROR] Sender mail param is empty.")
		return errors.New("Failed to send mail- Sender mail param is empty")
	}

	if len(s.config.SenderPassword()) == 0 {
		fmt.Printf("[ERROR] Sender mail password param is empty.")
		return errors.New("Failed to send mail- Sender mail  password param is empty")
	}

	if len(recipientInfo.ToEmail) == 0 {
		fmt.Printf("[ERROR] Recipient mail param is empty.")
		return errors.New("Failed to send mail- Recipient mail param is empty")
	}

	if len(recipientInfo.CCEmail) == 0 {
		fmt.Printf("[ERROR] CC mail param is empty.")
	}

	if len(recipientInfo.BCCEmail) == 0 {
		fmt.Printf("[WARING] BCC mail param is empty.")
	}

	if len(recipientInfo.MessageContent) == 0 {
		fmt.Printf("[ERROR] Message content param is empty.")
		return errors.New("Failed to send mail- Message content param is empty")
	}

	if len(recipientInfo.EmailType) == 0 {
		fmt.Printf("[WARNING] Mail content param is empty, setting mail type as plain/text")
		recipientInfo.EmailType = "text/plain"
	}

	mailClient := gomail.NewMessage()

	// Set E-Mail sender
	mailClient.SetHeader("From", s.config.SenderEmail())

	// Set E-Mail receivers
	mailClient.SetHeader("To", recipientInfo.ToEmail)
	mailClient.SetHeader("Cc", recipientInfo.CCEmail)
	mailClient.SetHeader("Bcc", recipientInfo.BCCEmail)

	// Set E-Mail subject
	mailClient.SetHeader("Subject", recipientInfo.MessageHeader)

	// Set E-Mail body. You can set plain text or html with text/html
	mailClient.SetBody(recipientInfo.EmailType, recipientInfo.MessageContent)

	SrvPort, err := strconv.Atoi(s.config.MailSrvPort())
	// Now send E-Mail
	if err != nil {
		fmt.Printf("[ERROR] Conversion error; Failed to sent mail, err:  %v\n", err)
		panic(err)
	}

	// Settings for SMTP server
	d := gomail.NewDialer(s.config.MailSrv(), SrvPort, s.config.SenderEmail(), s.config.SenderPassword())

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(mailClient); err != nil {
		fmt.Printf("[ERROR] Failed to sent mail, err:  %v\n", err)
		panic(err)
	}

	fmt.Printf("[INFO] Mail sent successfully, recipient:  %v\n", recipientInfo)

	return nil
}
