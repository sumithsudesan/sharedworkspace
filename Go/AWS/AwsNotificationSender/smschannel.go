package main

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

//SmsChannel - To Send SMS
type SmsChannel struct {
}

// NewSmsChannel - create new SmsChannel
func NewSmsChannel() IChannel {
	smschannel := SmsChannel{}
	return &smschannel
}

// Send - send SMS
func (s *SmsChannel) Send(recipientInfo RecipientInfo) error {
	fmt.Printf("[INFO] Sending sms to client\n")

	if len(recipientInfo.PhoneNumber) == 0 {
		fmt.Printf("[ERROR] PhoneNumber param is empty\n")
		return errors.New("Failed to send SMS- PhoneNumber param is empty")
	}

	if len(recipientInfo.MessageContent) == 0 {
		fmt.Printf("[ERROR] SMS MessageContent param is empty\n")
		return errors.New("Failed to send SMS- MessageContent param is empty")
	}

	// Creating aws session
	awsSession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(recipientInfo.Region), //us-west-2
	}))

	if awsSession == nil {
		fmt.Printf("[ERROR] Failed to create aws session for region : %v\n", recipientInfo.Region)
		return errors.New("Failed to create aws session")
	}

	// cretating aws service
	svc := sns.New(awsSession)
	if svc == nil {
		fmt.Printf("[ERROR] Failed to create svc for region : %v\n", recipientInfo.Region)
		return errors.New("Failed to create svc")
	}

	params := &sns.PublishInput{
		Subject:     aws.String(recipientInfo.MessageHeader),
		Message:     aws.String(recipientInfo.MessageContent),
		PhoneNumber: aws.String(recipientInfo.PhoneNumber),
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"AWS.SNS.SMS.SMSType": &sns.MessageAttributeValue{StringValue: aws.String("Transactional"), DataType: aws.String("String")},
		},
	}

	resp, err := svc.Publish(params)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Printf("[ERROR] Failed to Publish SMS ('%s'), error: %v\n", recipientInfo.PhoneNumber, err)
		return err
	}
	fmt.Printf("[INFO] SMS reponse : %v", resp)
	return nil
}
