package main

const (
	// DefaultFromEmailID - from mail id
	DefaultFromEmailID = "golangTest2021@gmail.com" //golang123!
	// DefaultSenderEmailPassword  - mail app password
	DefaultSenderEmailPassword = ""
	// DefaultMailServerPort - smtp port number
	DefaultMailServerPort = "587"
	// DefaultMailServer - smtp server
	DefaultMailServer = "smtp.gmail.com"
)

// Configuration -
type Configuration struct {
	//SenderEmailID - sender email id
	SenderEmailID string `json:"senderEmailId"`
	//SenderEmailPassword - sender email password
	SenderEmailPassword string `json:"senderEmailPassword"`
	//MailServerPort - smtp port
	MailServerPort string `json:"mailServerPort"`
	//MailServer - sender email server
	MailServer string `json:"mailServer"`
}

//RecipientInfo - recipient info from sns
type RecipientInfo struct {
	// ChannelType -  email, sms or both
	ChannelType string `json:"channelType"`
	//AppName            string `json:"appName"`
	// ToEmail - mail to address
	ToEmail string `json:"toEmail"`
	// CCEmail - cc mail id
	CCEmail string `json:"ccEmail"`
	// BCCEmail - bcc mail id
	BCCEmail string `json:"bccEmail"`
	// EmailType - html or plain
	EmailType string `json:"emailType"` // html or plain
	// EmailAttachmentURL - attachement url
	EmailAttachmentURL string `json:"emailAttachmentURL"`
	// PhoneNumber - Phone number to send sms
	PhoneNumber string `json:"phoneNumber"`
	// Region - region details (for SMS) - us-west-2
	Region string `json:"region"`
	// MessageHeader - message header
	MessageHeader string `json:"messageHeader"`
	//MessageContent- message content
	MessageContent string `json:"messageContent"`
}

// IChannel - interface for communictaion channel (sms/email)
type IChannel interface {
	Send(recipientInfo RecipientInfo) error
}
