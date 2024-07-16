package main

import (
	"errors"
	"fmt"
	"strings"
)

//ChannelFactory - To create chanel object
type ChannelFactory struct {
}

// Send - create chanel  object and send notifications
//channelType - mail, SMS, or ALL
func (s *ChannelFactory) Send(request RecipientInfo, mailConfig *Configuration) error {
	// c
	if len(request.ChannelType) == 0 {
		fmt.Printf("[ERROR] channel type empty\n")
		return errors.New("Failed to create channel- channel type empty")
	}

	fmt.Printf("[INFO] Creating channel object , channel type :%v\n", request.ChannelType)

	// MAIL
	if strings.ToUpper(request.ChannelType) == "MAIL" || strings.ToUpper(request.ChannelType) == "ALL" {
		mail := NewMailChannel(mailConfig)
		if mail != nil {
			mail.Send(request)
		}
	}

	// SMS
	if strings.ToUpper(request.ChannelType) == "SMS" || strings.ToUpper(request.ChannelType) == "ALL" {
		sms := NewSmsChannel()
		if sms != nil {
			sms.Send(request)
		}
		return nil // only for last condition
	}

	return errors.New("Invalid Channel type")
}
